package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
)

// DatabaseProvider @Bean("database")
type DatabaseProvider struct {
	config services.Config `inject:"config, database"`
}

func (m *DatabaseProvider) Init() {
	connections := m.config.GetKey("connections")

	for _, dataT := range connections {
		data, ok := dataT.(map[string]interface{})

		if !ok {
			continue
		}
		driver := data["driver"].(string)
		switch driver {
		case "postgresql":
		case "mysql":
			NewMysqlProvider()
		}
	}
}

func (m *DatabaseProvider) GetBean() interface{} {
	return nil
}
