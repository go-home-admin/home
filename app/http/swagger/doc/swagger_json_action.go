package doc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinHandleSwaggerJson gin原始路由处理
// http.Get(/swagger.json)
func (receiver *Controller) GinHandleSwaggerJson(ctx *gin.Context) {
	ctx.String(http.StatusOK, description)
}
