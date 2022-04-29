package providers

import (
	"database/sql"
	"fmt"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// MysqlProvider @Bean("mysql")
type MysqlProvider struct {
	config *services.Config `inject:"Route, database"`
	dbs    map[string]*gorm.DB
}

func (m *MysqlProvider) Init() {
	m.dbs = make(map[string]*gorm.DB)

	connections := m.config.GetKey("connections")

	for name, dataT := range connections {
		data, ok := dataT.(map[interface{}]interface{})

		if !ok {
			continue
		}
		driver := data["driver"].(string)
		if driver != "mysql" {
			continue
		}
		config := services.NewConfig(data)
		hosts := config.GetString("host", "127.0.0.1")
		port := config.GetInt("port", 3306)
		username := config.GetString("username")
		password := config.GetString("password")
		dbname := config.GetString("database")

		gConf := &gorm.Config{}
		// 调试时, 记录sql
		//if m.conf.IsDebug() {
		//	gConf.Logger = &logs.MysqlLog{}
		//}
		dsn := mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", username, password, hosts, port, dbname))
		db, err := gorm.Open(dsn, gConf)
		if err != nil {
			logrus.Error("mysql 链接错误")
			panic(err)
		}
		// https://github.com/go-sql-driver/mysql/issues/1120
		d := db.ConnPool.(*sql.DB)
		d.SetConnMaxIdleTime(60 * time.Second)
		m.dbs[name.(string)] = db
	}
}

func (m *MysqlProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
