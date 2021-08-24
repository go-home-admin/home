package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/toolset"
)

// Login  登陆
func (receiver *Controller) Login(req *toolset.LoginRequest, ctx *gin.Context) (*toolset.LoginResponse, error) {
	// TODO 这里写业务
	return &toolset.LoginResponse{}, nil
}

// GinHandleLogin gin原始路由处理
// http.Get(/login)
func (receiver *Controller) GinHandleLogin(ctx *gin.Context) {
	req := &toolset.LoginRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Login(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
