package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http/admin/public"
	"github.com/go-home-admin/home/bootstrap/http/api"
)

// @Bean
type AdminRoutes struct {
	public *public.Controller `inject:""`
}

// GetAuthRoutes Get{option (http.Route)}Routes
func (c *AdminRoutes) GetAuthRoutes() map[*api.Config]func(c *gin.Context) {
	return map[*api.Config]func(c *gin.Context){
		api.Post("/user/register"): c.public.GinHandleIndex,
	}
}
