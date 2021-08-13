package open2

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/admin"
)

// Index
func (receiver *Controller) Index(req *admin.IndexRequest, ctx *gin.Context) (*admin.IndexResponse, error) {
	// TODO 这里写业务
	return &admin.IndexResponse{}, nil
}

// GinHandleIndex gin原始路由处理
// http.Get(/open/admin)
func (receiver *Controller) GinHandleIndex(ctx *gin.Context) {
	req := &admin.IndexRequest{}
	err := ctx.ShouldBind(req)

	if err != nil {
		providers.ErrorRequest(ctx, err)
		return
	}

	resp, err := receiver.Index(req, ctx)
	if err != nil {
		providers.ErrorResponse(ctx, err)
		return
	}

	providers.SuccessResponse(ctx, resp)
}
