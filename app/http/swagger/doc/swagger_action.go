package doc

import (
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
	swagger "github.com/go-home-admin/home/generate/proto/swagger"
)

// Swagger
func (receiver *Controller) Swagger(req *swagger.InfoRequest, ctx *gin.Context) (*swagger.InfoResponse, error) {
	// TODO 这里写业务
	return &swagger.InfoResponse{}, nil
}

// GinHandleSwagger gin原始路由处理
// http.Get(/swagger)
func (receiver *Controller) GinHandleSwagger(ctx *gin.Context) {
	req := &swagger.InfoRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Swagger(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
