package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/http/route_help"
)

// @Bean
type Routes struct {
	AdminRoutes *AdminRoutes `inject:""`
}

// 映射所有组=>地址
func (r *Routes) GenRoutesConfig() map[string]route_help.GroupMap {
	return map[string]route_help.GroupMap{
		"admin-public": route_help.MergerRouteMap(
			r.AdminRoutes.GetAdminPublicRoutes(),
		),
	}
}

func (r *Routes) Load(engine *gin.Engine, config []route_help.GroupConfig, help *route_help.RouteHelp) {
	help.Load(engine, config, r.GenRoutesConfig())
}
