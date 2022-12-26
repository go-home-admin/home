package servers

import (
	"encoding/json"
	app2 "github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"github.com/go-home-admin/home/bootstrap/services/database"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"reflect"
	"sync/atomic"
	"time"
)

type DelayQueue interface {
	Push(DelayTask) string
	Del(string) bool
	Run()
}

var createTable = `
CREATE TABLE ` + "`" + `delay_queue` + "`" + ` (
  ` + "`" + `id` + "`" + ` varchar(64) NOT NULL,
  ` + "`" + `fail` + "`" + ` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '失败次数',
  ` + "`" + `route` + "`" + ` varchar(254) NOT NULL COMMENT '路由',
  ` + "`" + `job` + "`" + ` json NOT NULL COMMENT '任务信息',
  ` + "`" + `run_at` + "`" + ` timestamp NOT NULL DEFAULT '2022-08-25 00:00:00' COMMENT '执行时间点',
  ` + "`" + `created_at` + "`" + ` timestamp NOT NULL DEFAULT '2022-08-25 00:00:00' COMMENT '创建时间',
  ` + "`" + `in_cache` + "`" + ` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否加入缓存',
  PRIMARY KEY (` + "`" + `id` + "`" + `),
  KEY ` + "`" + `idx_delay_queue_run_at` + "`" + ` (` + "`" + `run_at` + "`" + `)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='延时队列存储';
`

// OrmDelayQueue @Bean
type OrmDelayQueue struct {
	Id        string        `gorm:"column:id;type:varchar(64);primaryKey"`                                                                             //
	Fail      uint64        `gorm:"column:fail;type:bigint(20) unsigned;default:0;comment:'失败次数'"`                                                     // 失败次数
	Route     string        `gorm:"column:route;type:varchar(254);comment:'路由'"`                                                                       // 路由
	Job       database.JSON `gorm:"column:job;type:json;comment:'任务信息'"`                                                                               // 任务信息
	RunAt     database.Time `gorm:"column:run_at;type:timestamp;index:idx_delay_queue_run_at,class:BTREE;default:2022-08-25 00:00:00;comment:'执行时间点'"` // 执行时间点
	CreatedAt database.Time `gorm:"column:created_at;type:timestamp;default:2022-08-25 00:00:00;comment:'创建时间'"`                                       // 创建时间
	InCache   uint32        `gorm:"column:in_cache;type:tinyint(4) unsigned;default:0;comment:'是否加入缓存'"`                                               // 是否加入缓存
}

func (receiver *OrmDelayQueue) TableName() string {
	return "delay_queue"
}

// DelayQueueForMysql @Bean("delay_queue")
type DelayQueueForMysql struct {
	mysql *gorm.DB `inject:"database, @config(queue.delay.connect)"`
	queue *Queue   `inject:""`

	RunAfterFuncLimit int64
}

func (d *DelayQueueForMysql) Init() {
	d.RunAfterFuncLimit = 500

	if app2.Config("queue.delay.auth_migrate", true) {
		logrus.Info("你可以修改queue.delay.auth_migrate = false; 关闭自动迁移delay_queue表")
		err := d.mysql.Exec(createTable)
		if err != nil {
			panic(err)
			return
		}
	}
}

func (d *DelayQueueForMysql) Run() {
	// 只有election选中的节点才启动
	if app.HasBean("election") {
		app.GetBean("election").(app.AppendRun).AppendRun(func() {
			d.Loop()
		})
	}
}

// Loop TODO 待优化，如果启动了广播，可以内存维护多个节点的最近任务，可以去掉定时查询
func (d *DelayQueueForMysql) Loop() {
	interval := app2.Config("queue.delay.interval", 60)

	// 从把in_cache全部加入缓存
	list := make([]*OrmDelayQueue, 0)
	d.mysql.Model(&OrmDelayQueue{}).Where("fail = 0 and in_cache = 1").
		Limit(int(d.RunAfterFuncLimit)).Order("id desc").
		Find(&list)
	for _, queue := range list {
		d.RunAfterFunc(queue)
	}
	for {
		list := make([]*OrmDelayQueue, 0)
		dbRet := d.mysql.Model(&OrmDelayQueue{}).Where("run_at <= ? and fail = 0 and in_cache = 0", time.Now().Add(time.Duration(60))).
			Limit(100).
			Find(&list)

		if dbRet.Error != nil {
			logrus.Error(dbRet.Error)
		} else if len(list) != 0 {
			delIds := make([]string, 0)
			for _, delayMsg := range list {
				// 还没有到时间就读出来, 设置到系统定时执行
				if delayMsg.RunAt.Time.Before(time.Now()) {
					if d.RunAfterFuncLimit > 0 {
						d.RunAfterFunc(delayMsg)
						d.mysql.Model(&OrmDelayQueue{}).Where("id = ?", delayMsg.Id).Update("in_cache", "1")
					}
				} else {
					if d.RunDelayJob(delayMsg) {
						delIds = append(delIds, delayMsg.Id)
					}
				}
			}
			if len(delIds) != 0 {
				d.mysql.Where("id in ?", delIds).Delete(&OrmDelayQueue{})
				// 有正常处理内容, 不需要睡眠
				continue
			}
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func (d *DelayQueueForMysql) RunDelayJob(delayMsg *OrmDelayQueue) bool {
	handle, ok := d.queue.dispatch.Load(delayMsg.Route)
	if !ok {
		logrus.Errorf("无法处理的route: %v", delayMsg.Route)
		d.mysql.Model(&OrmDelayQueue{}).Where("id = ?", delayMsg.Id).Update("fail", "1")
		return false
	}

	event := delayMsg.Job
	job := handle.(reflect.Value)
	v := job.Interface()
	newJob, ok := v.(constraint.Job)
	if ok {
		by, _ := event.MarshalJSON()
		err := json.Unmarshal(by, newJob)
		if err == nil {
			newJob.Handler()
			return true
		} else {
			logrus.Errorf("run delay job, json.Unmarshal data err = %v", err)
			d.mysql.Model(&OrmDelayQueue{}).Where("id = ?", delayMsg.Id).Update("fail", "1")
			return false
		}
	}
	logrus.Error("run delay job, v.(constraint.Job) not ok", v)
	return false
}

func (d *DelayQueueForMysql) Push(task DelayTask) string {
	uid := uuid.NewV4().String()
	delayMsg := &OrmDelayQueue{
		Id:        uid,
		Fail:      0,
		Route:     jobToRoute(task.message),
		Job:       database.NewJSON(task.message),
		RunAt:     database.Now().Add(task.interval),
		CreatedAt: database.Now(),
	}

	// 60 秒内直接设置定时齐
	if delayMsg.RunAt.Time.Before(time.Now().Add(60*time.Second)) && d.RunAfterFuncLimit > 0 {
		delayMsg.InCache = 1
	}

	ret := d.mysql.Model(&OrmDelayQueue{}).Create(delayMsg)
	if ret.Error != nil {
		if delayMsg.InCache == 1 {
			d.RunAfterFunc(delayMsg)
		}
	}
	return uid
}

func (d *DelayQueueForMysql) Del(id string) bool {
	ret := d.mysql.Where("id = ?", id).Delete(&OrmDelayQueue{})
	return ret.Error == nil
}

func (d *DelayQueueForMysql) RunAfterFunc(delayMsg *OrmDelayQueue) {
	atomic.AddInt64(&d.RunAfterFuncLimit, -1)
	dTime := delayMsg.RunAt.Unix() - time.Now().Unix()

	if dTime > 0 {
		time.AfterFunc(time.Duration(dTime)*time.Second, func() {
			atomic.AddInt64(&d.RunAfterFuncLimit, 1)
			if d.RunDelayJob(delayMsg) {
				d.mysql.Where("id = ?", delayMsg.Id).Delete(&OrmDelayQueue{})
			}
		})
	} else {
		atomic.AddInt64(&d.RunAfterFuncLimit, 1)
		if d.RunDelayJob(delayMsg) {
			d.mysql.Where("id = ?", delayMsg.Id).Delete(&OrmDelayQueue{})
		}
	}
}
