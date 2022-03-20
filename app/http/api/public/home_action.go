package public

import (
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
	api "github.com/go-home-admin/home/generate/proto/api"
)

// Home
func (receiver *Controller) Home(req *api.InfoRequest, ctx *gin.Context) (*api.InfoResponse, error) {
	// TODO 这里写业务
	return &api.InfoResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.InfoRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Home(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
