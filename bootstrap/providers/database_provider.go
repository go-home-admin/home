package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
)

// DatabaseProvider @Bean("database")
type DatabaseProvider struct {
	config services.Config `inject:"config, database"`

	m map[string]interface{}
}

func (m *DatabaseProvider) Init() {
	m.m = make(map[string]interface{})

	connections := m.config.GetKey("connections")

	for _, dataT := range connections {
		data, ok := dataT.(map[interface{}]interface{})

		if !ok {
			continue
		}
		driver := data["driver"].(string)
		switch driver {
		case "postgresql":
		case "mysql":
			m.m[driver] = NewMysqlProvider()
		}
	}
}

func (m *DatabaseProvider) GetBean(alias string) interface{} {
	return m.m[alias]
}
