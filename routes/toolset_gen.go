package routes

import (
	home_gin "github.com/gin-gonic/gin"
	public "github.com/go-home-admin/home/app/http/toolset/public"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// ToolsetRoutes @Bean
type ToolsetRoutes struct {
	public *public.Controller `inject:""`
}

// GetToolsetRoutes Get{option (http.Route)}Routes
func (c *ToolsetRoutes) GetToolsetRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Post("/login"): c.public.GinHandleLogin,
	}
}
