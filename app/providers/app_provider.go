package providers

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
)

// App 系统引导结构体
// @Bean
type App struct {
	container *services.Container `inject:""`
	resp      *Response           `inject:""`
	log       *Log                `inject:""`
}

func (a *App) Run(servers []constraint.KernelServer) {
	a.container.Run(servers)
}
