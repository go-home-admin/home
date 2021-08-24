package http

import (
	"github.com/gin-gonic/gin"
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
	k.setHttp()
	// 这里需要注册你的业务前缀, 中间件
	k.routes.Load(
		k.httpServer.GetEngine(),
		[]route_help.GroupConfig{
			{Name: "api"},
			{
				Name:   "admin-public",
				Prefix: "/admin",
			},
			{
				Name:        "admin",
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

func (k *Kernel) setHttp() {
	k.httpServer.SetPort(k.config.GetServiceConfig("http").GetInt("port", 80))
	k.httpServer.SetDebug(k.config.IsDebug())

	// 默认允许跨域
	engine := gin.New()
	engine.Use(Cors())
	k.httpServer.SetEngine(engine)
}

// GetServer 提供统一命名规范的独立服务
func GetServer() constraint.KernelServer {
	return InitializeNewKernelProvider()
}
