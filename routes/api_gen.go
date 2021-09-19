package routes

import (
	home_gin "github.com/gin-gonic/gin"
	api_demo "github.com/go-home-admin/home/app/http/api/api_demo"
	public "github.com/go-home-admin/home/app/http/api/public"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// ApiRoutes @Bean
type ApiRoutes struct {
	public   *public.Controller   `inject:""`
	api_demo *api_demo.Controller `inject:""`
}

// GetApiRoutes Get{option (http.Route)}Routes
func (c *ApiRoutes) GetApiRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/"): c.public.GinHandleHome,

		home_api.Get("/api/demo"): c.api_demo.GinHandleHome,
	}
}
