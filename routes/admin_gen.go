package routes

import (
	home_gin "github.com/gin-gonic/gin"
	public "github.com/go-home-admin/home/app/http/admin/public"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// AdminRoutes @Bean
type AdminRoutes struct {
	public *public.Controller `inject:""`
}

// GetAdminPublicRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetAdminPublicRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/"): c.public.GinHandleIndex,
	}
}
