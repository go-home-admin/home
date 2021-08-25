package routes

import (
	home_gin "github.com/gin-gonic/gin"
	doc "github.com/go-home-admin/home/app/http/swagger/doc"
	home_api "github.com/go-home-admin/home/bootstrap/http/api"
)

// SwaggerRoutes @Bean
type SwaggerRoutes struct {
	doc *doc.Controller `inject:""`
}

// GetSwaggerRoutes Get{option (http.Route)}Routes
func (c *SwaggerRoutes) GetSwaggerRoutes() map[*home_api.Config]func(c *home_gin.Context) {
	return map[*home_api.Config]func(c *home_gin.Context){
		home_api.Get("/swagger"):      c.doc.GinHandleSwagger,
		home_api.Get("/swagger.json"): c.doc.GinHandleSwaggerJson,
	}
}
