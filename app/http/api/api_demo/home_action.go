package api_demo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/api"
)

// Home  首页
func (receiver *Controller) Home(req *api.HomeRequest, ctx *gin.Context) (*api.HomeResponse, error) {
	// TODO 这里写业务
	return &api.HomeResponse{}, nil
}

// GinHandleHome gin原始路由处理
// http.Get(/)
func (receiver *Controller) GinHandleHome(ctx *gin.Context) {
	req := &api.HomeRequest{}
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
