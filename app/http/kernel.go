package http

import (
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/http/route_help"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/routes"
)

// Kernel @Bean
type Kernel struct {
	routes     *routes.Routes       `inject:""`
	httpServer *services.HttpServer `inject:""`
	config     *providers.Config    `inject:""`
}

func (k *Kernel) Init() {
	serviceConfig := k.config.GetServiceConfig("http")

	k.httpServer.SetPort(serviceConfig.GetInt("port"))
	k.httpServer.SetDebug(serviceConfig.GetBool("debug") == true)

	// 这里需要注册你的业务前缀, 中间件
	k.routes.Load(
		k.httpServer.GetEngine(),
		[]route_help.GroupConfig{
			{
				Name:   "admin-public",
				Prefix: "/admin",
			},
			{
				Name:        "admin-login",
				Prefix:      "/admin",
				Middlewares: nil,
			},
		},
		&route_help.RouteHelp{},
	)
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
