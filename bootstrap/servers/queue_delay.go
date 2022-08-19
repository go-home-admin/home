package servers

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
	"github.com/go-home-admin/home/bootstrap/services/database"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DelayQueue interface {
	Push(DelayTask) string
	Del(string) bool
	Run()
}

// OrmDelayQueue @Bean
type OrmDelayQueue struct {
	Id        string        `gorm:"primaryKey;column:id;type:varchar;not null" json:"id"`
	Job       database.JSON `gorm:"column:job;type:json;not null" json:"job"`
	RunAt     database.Time `gorm:"column:run_at;type:timestamp;not null" json:"run_at"`
	CreatedAt database.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
}

func (receiver *OrmDelayQueue) TableName() string {
	return "delay_queue"
}

// DelayQueueForMysql @Bean("delay_queue")
type DelayQueueForMysql struct {
	mysql *gorm.DB `inject:"database, @config(queue, delay.connect)"`
}

func (d *DelayQueueForMysql) Init() {
	d.mysql.AutoMigrate(&OrmDelayQueue{})
}

func (d *DelayQueueForMysql) Run() {
	// 只有election选中的节点才启动
	if app.HasBean("election") {
		app.GetBean("election").(app.AppendRun).AppendRun(func() {
			d.Loop()
		})
	}
}

func (d *DelayQueueForMysql) Loop() {
	for {

	}
}

func (d *DelayQueueForMysql) Push(task DelayTask) string {
	uid := uuid.NewV4().String()

	d.mysql.Model(&OrmDelayQueue{}).Create(&OrmDelayQueue{
		Id:        uid,
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
