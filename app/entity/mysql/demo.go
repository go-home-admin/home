package mysql

import (
	database "github.com/go-home-admin/home/bootstrap/services/database"
	"gorm.io/gorm"
)

type ActionLogs struct {
	Id       int64         `gorm:"primaryKey;column:id"`        //
	Date     database.Time `gorm:"column:date"`                 // 时间
	Guard    string        `gorm:"column:guard"`                // guard
	Ip       string        `gorm:"column:ip"`                   // ip
	Uri      string        `gorm:"primaryKey;column:uri"`       // 请求地址
	Params   string        `gorm:"column:params"`               // 参数
	UserId   int64         `gorm:"column:user_id"`              // uid
	User     string        `gorm:"column:user"`                 // uid
	HttpType string        `gorm:"primaryKey;column:http_type"` // http_type
}

// OrmActionLogs @Bean("action_logs")
type OrmActionLogs struct {
	db *gorm.DB `inject:"database, mysql"`
}

func (orm *OrmActionLogs) Insert(row ActionLogs) *gorm.DB {
	return orm.db.Create(row)
}

func (orm *OrmActionLogs) Inserts(rows []ActionLogs) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmActionLogs) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&ActionLogs{}, conds...)
}
