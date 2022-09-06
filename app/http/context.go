package http

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
	return receiver.Context
}

func (receiver Ctx) Id() uint64 {
	parts := strings.Split(receiver.Token(), ".")
	if len(parts) == 3 {
		if l := len(parts[1]) % 4; l > 0 {
			parts[1] += strings.Repeat("=", 4-l)
		}
		b, err := base64.URLEncoding.DecodeString(parts[1])
		if err == nil {
			var res = map[string]uint64{"id": 0}
			_ = json.Unmarshal(b, &res)
			return res["id"]
		}
	}
	return 0
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
}
