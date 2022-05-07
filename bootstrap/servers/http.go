package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/sirupsen/logrus"
)

// Http 提供者
// @Bean("http")
type Http struct {
	*providers.RouteProvider `inject:""`
	*services.HttpServer     `inject:""`
	*services.Config         `inject:"config, app.servers.http"`

	*gin.Engine
	Middleware      []gin.HandlerFunc
	MiddlewareGroup map[string][]gin.HandlerFunc

	Port string
	init bool
}

func (http *Http) Init() {
	http.Port = http.GetString("port", "80")
	http.Engine = gin.New()
}

func (http *Http) Boot() {
	if http.init {
		return
	}
	http.init = true

	// 初始化所有配置
	group := make(map[string]*gin.RouterGroup)
	for gn, _ := range http.Route {
		gc, ok := http.RouteGroupConfig[gn]
		if ok {
			group[gn] = http.Engine.Group(gc.GetPrefix())
		} else {
			http.RouteGroupConfig[gn] = &providers.GroupConfig{}
			group[gn] = http.Engine.Group("")
		}
	}

	// 全局中间件设置
	for _, handlerFunc := range http.Middleware {
		for gn, _ := range http.Route {
			gr := group[gn]
			gr.Use(handlerFunc)
		}
	}

	// 分组中间件设置
	for gn, gm := range http.Route {
		gc, _ := http.RouteGroupConfig[gn]
		gr := group[gn]

		for _, middlewareGroupName := range gc.GetMiddleware() {
			gorpMics, ok := http.MiddlewareGroup[middlewareGroupName]
			if ok {
				for _, handlerFunc := range gorpMics {
					gr.Use(handlerFunc)
				}
			}
		}

		for f, fun := range gm {
			config := *f
			switch config["method"] {
			case "get":
				gr.GET(config["url"], fun)
			case "post":
				gr.POST(config["url"], fun)
			case "put":
				gr.PUT(config["url"], fun)
			case "delete":
				gr.DELETE(config["url"], fun)
			case "options":
				gr.OPTIONS(config["url"], fun)
			case "any":
				gr.Any(config["url"], fun)
			}
		}
	}
}

func (http *Http) Run() {
	err := http.Engine.Run(":" + http.Port)
	if err != nil {
		logrus.WithFields(logrus.Fields{"port": http.Port}).Error("http发生错误")
	}
}

func (http *Http) Exit() {}
