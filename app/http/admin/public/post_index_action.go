package public

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/generate/proto/admin"
)

// PostIndex
func (receiver *Controller) PostIndex(req *admin.IndexRequest, ctx *gin.Context) (*admin.IndexResponse, error) {
	// TODO 这里写业务
	return &admin.IndexResponse{}, nil
}

// GinHandlePostIndex gin原始路由处理
// http.Post(/admin)
func (receiver *Controller) GinHandlePostIndex(ctx *gin.Context) {
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
