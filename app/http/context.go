package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/database"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const UserIdKey = "user_id"

// UserModel 不能赋值指针
var UserModel interface{}

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

	UserInfo interface{}
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
	if receiver.UserInfo == nil {
		receiver.InitUser()
	}
	return receiver.UserInfo
}

func (receiver Ctx) Id() uint64 {
	u, ok := receiver.Context.Get(UserIdKey)
	if !ok {
		logrus.Fatal("id 不存在, todo Context.Set(UserIdKey, Uid)")
		return 0
	}
	return u.(uint64)
}

func (receiver Ctx) IdStr() string {
	u, ok := receiver.Context.Get(UserIdKey)
	if !ok {
		return ""
	}
	return u.(string)
}

func (receiver Ctx) Token() string {
	tokenString := receiver.Context.GetHeader("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	}
	return tokenString
}

func (receiver Ctx) InitUser() {
	if receiver.UserInfo == nil {
		uid, ok := receiver.Context.Get(UserIdKey)
		if ok {
			user := UserModel
			database.DB().Model(UserModel).First(&user, uid)

			receiver.UserInfo = user
		}
	}
}

type Context interface {
	Success(data interface{})
	Fail(err error)
	Gin() *gin.Context
	Token() string
	Id() uint64
	IdStr() string
	User() interface{}
}
