package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const UserKey = "user"
const UserIdKey = "user_id"

// NewContext 最好在中间件已经赋值以下两个参数
// ginCtx.Set("user", nil)
// ginCtx.Set("user_id", nil)
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

func (receiver Ctx) Fail(err error) {
	receiver.JSON(http.StatusOK, map[string]interface{}{
		"code": 1,
		"msg":  err.Error(),
	})
}

func (receiver Ctx) Gin() *gin.Context {
	return receiver.Context
}

func (receiver Ctx) User() interface{} {
	u, ok := receiver.Context.Get(UserKey)
	if !ok {
		return nil
	}
	return u
}

func (receiver Ctx) Id() uint64 {
	u, ok := receiver.Context.Get(UserIdKey)
	if !ok {
		return 0
	}
	return u.(uint64)
}

func (receiver Ctx) Token() string {
	tokenString := receiver.Context.GetHeader("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	return tokenString
}

type Context interface {
	Success(data interface{})
	Fail(err error)
	Gin() *gin.Context
	Token() string
	Id() uint64
	User() interface{}
}
