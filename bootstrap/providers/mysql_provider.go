package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"gorm.io/gorm"
)

// MysqlProvider @Bean("mysql")
type MysqlProvider struct {
	config *services.Config `inject:"config, database"`
	dbs    map[string]*gorm.DB
}

func (m *MysqlProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
