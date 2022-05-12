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
	connections := m.Config.GetKey("connections")

	for name, dataT := range connections {
		data, ok := dataT.(map[interface{}]interface{})

		if !ok {
			continue
		}
		if name == alias {
			driver := data["driver"].(string)
			switch driver {
			case "mysql":
				return NewMysqlProvider().GetBean(name.(string))
			default:
				return app.GetBean(driver).(app.Bean).GetBean(name.(string))
			}
		}
	}
	return nil
}
