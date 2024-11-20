package postgresql

import (
	"fmt"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

// PostgresqlProvider @Bean("postgresql")
type PostgresqlProvider struct {
	config *services.Config `inject:"config, database"`
	dbs    map[string]*gorm.DB
}

func (p *PostgresqlProvider) Init() {
	p.dbs = make(map[string]*gorm.DB)

	connections := p.config.GetKey("connections")

	for name, dataT := range connections {
		data, ok := dataT.(map[interface{}]interface{})

		if !ok {
			continue
		}
		driver := data["driver"].(string)
		if driver != "postgresql" {
			continue
		}
		config := services.NewConfig(data)
		host := config.GetString("host", "127.0.0.1")
		port := config.GetInt("port", 5432)
		username := config.GetString("username")
		password := config.GetString("password")
		dbname := config.GetString("database")
		timezone := config.GetString("timezone", "Asia/Shanghai")

		gConf := &gorm.Config{
			Logger: logs.NewSqlLog(logrus.StandardLogger(), logger.Config{
				SlowThreshold: 0,
				LogLevel:      logger.Warn,
				Colorful:      true,
			}),
		}
		if app.IsDebug() {
			gConf.Logger.LogMode(logger.LogLevel(logrus.DebugLevel))
		}

		dsnStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require TimeZone=%v", host, username, password, dbname, port, timezone)
		dsn := postgres.Open(dsnStr)
		db, err := gorm.Open(dsn, gConf)
		if err != nil {
			dsn = postgres.Open(strings.Replace(dsnStr, "require", "disable", 1))
			db, err = gorm.Open(dsn, gConf)
			if err != nil {
				logrus.Error("postgresql 链接错误", err)
				panic(err)
			}
		}
		p.dbs[name.(string)] = db
	}
}

func (p *PostgresqlProvider) GetBean(alias string) interface{} {
	return p.dbs[alias]
}
