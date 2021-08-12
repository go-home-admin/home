package providers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var ErrorRequest func(ctx *gin.Context, err error)

var ErrorResponse func(ctx *gin.Context, err error)

var SuccessResponse func(ctx *gin.Context, data interface{})

// @Bean
type Response struct{}

func NewResponseProvider() *Response {
	Response := &Response{}
	Response.init()
	return Response
}

func (r Response) init() {
	ErrorRequest = func(ctx *gin.Context, err error) {

	}

	ErrorResponse = func(ctx *gin.Context, err error) {

	}

	SuccessResponse = func(ctx *gin.Context, data interface{}) {
		Json(ctx, http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    data,
		})
	}
}

// Json 唯一输出的出口
func Json(ctx *gin.Context, code int, h gin.H) {
	h["time"] = time.Now().Unix()
	str, _ := json.Marshal(h)
	ctx.Header("Content-Type", "application/json;charset=utf-8")
	ctx.String(code, string(str))
}
