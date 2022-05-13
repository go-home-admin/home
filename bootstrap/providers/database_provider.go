package providers

import (
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

	driver := config.GetString("driver")
	switch driver {
	case "mysql":
		return NewMysqlProvider().GetBean(alias)
	case "redis":
		return NewRedisProvider().GetBean(alias)
	default:
		return app.GetBean(driver).(app.Bean).GetBean(alias)
	}
}
