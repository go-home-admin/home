package providers

import (
	"database/sql"
	"fmt"
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Mysql @Bean("mysql")
type Mysql struct {
	conf *Config `inject:""`
	db   *gorm.DB
	log  *Log `inject:""` // 引入不用, 为了方便初始化log

	dbs map[string]*gorm.DB
}

func (m *Mysql) Init() {
	m.dbs = make(map[string]*gorm.DB)

	config := m.conf.GetServiceConfig("mysql")
	hosts := config.GetString("hosts")
	port := config.GetString("port")
	username := config.GetString("username")
	password := config.GetString("password")
	dbname := config.GetString("dbname")

	gConf := &gorm.Config{}
	// 调试时, 记录sql
	if m.conf.IsDebug() {
		gConf.Logger = &logs.MysqlLog{}
	}
	var err error
	dsn := mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, hosts, port, dbname))
	m.db, err = gorm.Open(dsn, gConf)
	if err != nil {
		logrus.Error("mysql 链接错误")
		panic(err)
	}
	// https://github.com/go-sql-driver/mysql/issues/1120
	d := m.db.ConnPool.(*sql.DB)
	d.SetConnMaxIdleTime(60 * time.Second)
}

func (m *Mysql) DB() *gorm.DB {
	return m.db
}

func (m *Mysql) GetBean(alias string) *gorm.DB {
	return m.dbs[alias]
}
