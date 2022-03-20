package providers

import (
	"fmt"
	"github.com/go-home-admin/home/bootstrap/services"
)

// DatabaseProvider @Bean
type DatabaseProvider struct {
	config services.Config `inject:"config, database"`
}

func (m *DatabaseProvider) Init() {
	connections := m.config.GetKey("connections")

	for i, i2 := range connections {

		fmt.Println(i, i2)
	}
}
