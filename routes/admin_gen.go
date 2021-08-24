package routes

import (
	home_gin "github.com/gin-gonic/gin"
	admin_user "github.com/go-home-admin/home/app/http/admin/admin_user"
	public "github.com/go-home-admin/home/app/http/admin/public"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// AdminRoutes @Bean
type AdminRoutes struct {
	admin_user *admin_user.Controller `inject:""`
	public     *public.Controller     `inject:""`
}

// GetAdminRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetAdminRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/info"): c.admin_user.GinHandleInfo,
	}
}

// GetAdminPublicRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetAdminPublicRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Post("/login"):  c.public.GinHandleLogin,
		home_api.Post("/logout"): c.public.GinHandleLogout,
	}
}
