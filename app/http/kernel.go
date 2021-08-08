package http

import (
	"fmt"
	"github.com/go-home-admin/home/app/provoders"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
)

// Kernel @Bean
type Kernel struct {
	httpServer *services.HttpServer `inject:""`
	config     *provoders.Config    `inject:""`
}

func (k *Kernel) Init() {
	serviceConfig := k.config.GetServiceConfig("http")

	fmt.Println(serviceConfig)
}

func (k *Kernel) Run() {
	k.httpServer.RunListener()
}

func (k *Kernel) Exit() {

}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return InitializeNewKernelProvider()
}
