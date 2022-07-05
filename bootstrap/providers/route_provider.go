package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/http/api"
)

// RouteProvider 路由提供者
// @Bean
type RouteProvider struct {
	Route            map[string]map[*api.Config]func(c *gin.Context)
	RouteGroupConfig map[string]*GroupConfig
}

func (a *RouteProvider) Init() {
	a.Route = make(map[string]map[*api.Config]func(c *gin.Context))
	a.RouteGroupConfig = make(map[string]*GroupConfig)
}

func (a *RouteProvider) LoadRoute(routes []interface{}) {
	for _, inf := range routes {
		route, ok := inf.(constraint.Route)
		if ok {
			g := route.GetGroup()
			gg, gok := a.Route[g]
			if gok {
				for config, f := range route.GetRoutes() {
					gg[config] = f
				}
				a.Route[g] = gg
			} else {
				a.Route[route.GetGroup()] = route.GetRoutes()
			}
		}
	}
}

func (a *RouteProvider) Group(group string) *GroupConfig {
	g, ok := a.RouteGroupConfig[group]

	if !ok {
		g = &GroupConfig{}
		a.RouteGroupConfig[group] = g
	}
	return g
}

type GroupConfig struct {
	prefix     string
	middleware []string
}

func (g *GroupConfig) GetPrefix() string {
	return g.prefix
}
func (g *GroupConfig) GetMiddleware() []string {
	return g.middleware
}

func (g *GroupConfig) Prefix(prefix string) *GroupConfig {
	g.prefix = prefix
	return g
}

func (g *GroupConfig) Middleware(m ...string) *GroupConfig {
	if g.middleware == nil {
		g.middleware = make([]string, 0)
	}

	for _, s := range m {
		g.middleware = append(g.middleware, s)
	}
	return g
}
