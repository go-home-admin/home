package providers

import (
	"fmt"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"github.com/go-home-admin/home/database"
	"gorm.io/gorm"
)

// DatabaseProvider @Bean("database")
type DatabaseProvider struct {
	*services.Config `inject:"config, database"`
}

func (m *DatabaseProvider) Init() {
	d := m.GetString("default")

	database.DB = func() *gorm.DB {
		return NewDatabaseProvider().GetBean(d).(*gorm.DB)
	}
}

func (m *DatabaseProvider) GetBean(alias string) interface{} {
	config := m.Config.GetConfig("connections." + alias)
	if config == nil {
		if alias == "default" {
			d := m.GetString("default")
			return NewDatabaseProvider().GetBean(d).(*gorm.DB)
		}
		panic(fmt.Sprintf("你需要在config/database.yaml添加您的%v数据库配置", alias))
	}

	driver := config.GetString("driver")
	if !app.HasBean(driver) {
		switch driver {
		case "mysql":
			return NewMysqlProvider().GetBean(alias)
		case "redis":
			return NewRedisProvider().GetBean(alias)
		default:
			panic("您的数据库驱动, 必须在使用前注册")
		}
	}
	return app.GetBean(driver).(app.Bean).GetBean(alias)
}
