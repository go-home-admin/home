package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"gorm.io/gorm"
)

// @Bean("mysql")
type MysqlProvider struct {
	config services.Config `inject:"config, mysql"`
	dbs    map[string]*gorm.DB
}

func (m *MysqlProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
