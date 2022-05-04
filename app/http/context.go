package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var NewContext = func(ctx *gin.Context) Context {
	return &Ctx{
		Context: ctx,
	}
}

type Ctx struct {
	*gin.Context
}

func (receiver Ctx) Success(data interface{}) {
	receiver.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
		"code": 0,
		"msg":  "",
	})
}

func (receiver Ctx) Fail(code int, msg string) {
	receiver.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
	})
}

func (receiver Ctx) Gin() *gin.Context {
	return receiver.Context
}

type Context interface {
	Success(data interface{})
	Fail(code int, msg string)
	Gin() *gin.Context
}
