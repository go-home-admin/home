package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/api"
	"time"
)

var start string

func init() {
	start = time.Now().String()
}

// Home  首页
func (receiver *Controller) Home(req *api.HomeRequest, ctx *gin.Context) (*api.HomeResponse, error) {

	return &api.HomeResponse{
		StartTime: start,
		Tip:       "",
	}, nil
}

// GinHandleHome gin原始路由处理
// http.Post(/login)
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
