package constraint

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/http/api"
)

type ServiceProvider interface {
	register()
	boot()
}

// Construct 具备结构体Bean
type Construct interface {
	// Init 初始化
	Init()
}

// AfterRegistration 注册完成后, 统一处理
type AfterRegistration interface {
	Boot()
}

// RunAfter 在Run执行后执行, 可以做服务准备好状态维护等
type RunAfter interface {
	RunAfter()
}

type Exit interface {
	Exit()
}

type Route interface {
	GetGroup() string
	GetRoutes() map[*api.Config]func(c *gin.Context)
}
