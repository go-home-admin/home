package public

import (
	gin "github.com/gin-gonic/gin"
	providers "github.com/go-home-admin/home/app/providers"
	admin "github.com/go-home-admin/home/generate/proto/admin"
)

// Logout
func (receiver *Controller) Logout(req *admin.InfoRequest, ctx *gin.Context) (*admin.InfoResponse, error) {
	// TODO 这里写业务
	return &admin.InfoResponse{}, nil
}

// GinHandleLogout gin原始路由处理
// http.Post(/logout)
func (receiver *Controller) GinHandleLogout(ctx *gin.Context) {
	req := &admin.InfoRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Logout(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
