package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
)

// Kernel @Bean
type Kernel struct {
	httpServer *services.HttpServer `inject:""`
	config     *providers.Config    `inject:""`
}

func (k *Kernel) Init() {
	k.setHttp()

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
	return NewKernel()
}
