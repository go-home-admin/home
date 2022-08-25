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
	"time"
)

type DelayQueue interface {
	Push(DelayTask) string
	Del(string) bool
	Run()
}

// OrmDelayQueue @Bean
type OrmDelayQueue struct {
	Id        string        `gorm:"primaryKey;autoIncrement;column:id;type:varchar(64);not null" json:"id"`
	Fail      int           `gorm:"column:fail;type:int" json:"fail"`
	Route     string        `gorm:"column:route;type:varchar(254)" json:"route"`
	Job       database.JSON `gorm:"column:job;type:json" json:"job"`
	RunAt     database.Time `gorm:"column:run_at;type:timestamp;index;not null" json:"run_at"`
	CreatedAt database.Time `gorm:"column:created_at;type:timestamp;not null;default:'2022-08-25 00:00:00'" json:"created_at"`
}

func (receiver *OrmDelayQueue) TableName() string {
	return "delay_queue"
}

// DelayQueueForMysql @Bean("delay_queue")
type DelayQueueForMysql struct {
	mysql *gorm.DB `inject:"database, @config(queue.delay.connect)"`
	queue *Queue   `inject:""`
}

func (d *DelayQueueForMysql) Init() {
	if app2.Config("queue.delay.auth_migrate", true) {
		logrus.Info("你可以修改queue.delay.auth_migrate = false; 关闭自动迁移delay_queue表")
		err := d.mysql.AutoMigrate(&OrmDelayQueue{})
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
	for {
		list := make([]*OrmDelayQueue, 0)
		dbRet := d.mysql.Model(&OrmDelayQueue{}).Where("run_at <= ? and fail = 0", time.Now()).Limit(100).Order("Id desc").Find(&list)

		if dbRet.Error != nil {
			logrus.Error(dbRet.Error)
		} else if len(list) != 0 {
			delIds := make([]string, 0)
			for _, delayMsg := range list {
				handle, ok := d.queue.dispatch.Load(delayMsg.Route)
				if !ok {
					logrus.Errorf("无法处理的route: %v", delayMsg.Route)
					d.mysql.Model(&OrmDelayQueue{}).Where("id <= ?", delayMsg.Id).Update("fail", "1")
					continue
				}

				event := delayMsg.Job
				job := handle.(reflect.Value)
				v := job.Interface()
				newJob, ok := v.(constraint.Job)
				if ok {
					err := json.Unmarshal([]byte(event), newJob)
					if err == nil {
						newJob.Handler()
						delIds = append(delIds, delayMsg.Id)
					} else {
						logrus.Errorf("run delay job, json.Unmarshal data err = %v", err)
						d.mysql.Model(&OrmDelayQueue{}).Where("id <= ?", delayMsg.Id).Update("fail", "1")
					}
				}
			}
			d.mysql.Where("id in ?", delIds).Delete(&OrmDelayQueue{})
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func (d *DelayQueueForMysql) Push(task DelayTask) string {
	uid := uuid.NewV4().String()

	d.mysql.Model(&OrmDelayQueue{}).Create(&OrmDelayQueue{
		Id:        uid,
		Fail:      0,
		Route:     jobToRoute(task.message),
		Job:       database.NewJSON(task.message),
		RunAt:     database.Now().Add(task.interval),
		CreatedAt: database.Now(),
	})

	return uid
}

func (d *DelayQueueForMysql) Del(id string) bool {
	ret := d.mysql.Where("id = ?", id).Delete(&OrmDelayQueue{})
	return ret.Error == nil
}
