package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"gorm.io/gorm"
)

// Mysql @Bean("mysql")
type Mysql struct {
	config services.Config `inject:"config, mysql"`
	dbs    map[string]*gorm.DB
}

func (m *Mysql) Init() {
	m.dbs = make(map[string]*gorm.DB)
}

func (m *Mysql) GetBean(alias string) *gorm.DB {
	return m.dbs[alias]
}
