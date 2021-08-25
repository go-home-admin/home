package doc

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/swagger"
	"net/http"
)

// Swagger  文档
func (receiver *Controller) Swagger(req *swagger.SwaggerNull, ctx *gin.Context) error {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, html)

	return nil
}

// GinHandleSwagger gin原始路由处理
// http.Get(/swagger)
func (receiver *Controller) GinHandleSwagger(ctx *gin.Context) {
	req := &swagger.SwaggerNull{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	err = receiver.Swagger(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}
}
