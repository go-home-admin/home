package routes

import (
	home_gin "github.com/gin-gonic/gin"
	open2 "github.com/go-home-admin/home/app/http/admin/open2"
	public "github.com/go-home-admin/home/app/http/admin/public"
	test "github.com/go-home-admin/home/app/http/admin/test"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// AdminRoutes @Bean
type AdminRoutes struct {
	public *public.Controller `inject:""`
	open2  *open2.Controller  `inject:""`
	test   *test.Controller   `inject:""`
}

// GetOpenRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetOpenRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/admin"):  c.public.GinHandleIndex,
		home_api.Post("/admin"): c.public.GinHandlePostIndex,

		home_api.Get("/open/admin"):  c.open2.GinHandleIndex,
		home_api.Post("/open/admin"): c.open2.GinHandlePostIndex,
	}
}

// GetTestRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetTestRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/test/admin"):  c.test.GinHandleIndex,
		home_api.Post("/test/admin"): c.test.GinHandlePostIndex,
	}
}
