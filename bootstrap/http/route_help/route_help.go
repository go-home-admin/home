package route_help

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/http/api"
	"github.com/go-home-admin/home/bootstrap/services/logs"
)

type RouteHelp struct {
}

type GroupConfig struct {
	// option (http.RouteGroup) = "open";
	Name string
	// 前缀
	Prefix string
	// 这个分组使用中间件
	Middlewares []func(ctx *gin.Context)
}

type GroupMap map[*api.Config]func(c *gin.Context)

func MergerRouteMap(ms ...GroupMap) GroupMap {
	got := make(GroupMap)
	for _, m := range ms {
		for config, f := range m {
			got[config] = f
		}
	}
	return got
}

func (r *RouteHelp) Load(engine *gin.Engine, config []GroupConfig, allGroupRoute map[string]GroupMap) {
	isUserConfig := make(map[string]bool)
	for gn, gm := range allGroupRoute {
		if _, ok := isUserConfig[gn]; !ok {
			isUserConfig[gn] = false
		}

		for _, groupConfig := range config {
			if groupConfig.Name == gn {
				gr := engine.Group(groupConfig.Prefix)
				if groupConfig.Middlewares != nil {
					for _, middleware := range groupConfig.Middlewares {
						gr.Use(middleware)
					}
				}
				r.loadRoutes(gr, gm)
				isUserConfig[gn] = true
			}
		}
	}

	// 检查未使用的分组
	for s, b := range isUserConfig {
		if b == false {
			logs.Debug("proto的路由【", s, "】未在kernel配置前缀、中间件等信息, 它的路由信息未能正常加载")
		}
	}
}

// 分组路由加载, 前缀, 中间件, 路由函数列表
func (r *RouteHelp) MiddlewareRoutes(
	engine *gin.Engine,
	path string,
	middleware func(ctx *gin.Context),
	routes map[*api.Config]func(c *gin.Context),
) {
	g := engine.Group(path)
	g.Use(middleware)
	r.loadRoutes(g, routes)
}

func (r *RouteHelp) loadRoutes(engine *gin.RouterGroup, routes map[*api.Config]func(c *gin.Context)) {
	for f, fun := range routes {
		config := *f
		switch config["method"] {
		case "get":
			engine.GET(config["url"], fun)
		case "post":
			engine.POST(config["url"], fun)
		case "put":
			engine.PUT(config["url"], fun)
		case "delete":
			engine.DELETE(config["url"], fun)
		}
	}
}
