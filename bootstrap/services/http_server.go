package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

// HttpServer 封装一些被(工具生成的)代码调用的函数
type HttpServer struct {
	isDebug bool
	port    string
	host    string
	engine  *gin.Engine
}

func (receiver *HttpServer) Init() {
	receiver.port = "80"
}

func (receiver *HttpServer) SetDebug(isDebug bool) {
	receiver.isDebug = isDebug
	if receiver.isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func (receiver *HttpServer) SetEngine(engine *gin.Engine) {
	receiver.engine = engine
}

func (receiver *HttpServer) GetEngine() *gin.Engine {
	return receiver.engine
}

func (receiver *HttpServer) SetPort(port int) {
	if port == 0 {
		port = 80
		logrus.Warning("端口错误, 转为80")
	}
	receiver.port = strconv.Itoa(port)
}

func (receiver *HttpServer) RunListener() {
	err := receiver.GetEngine().Run(":" + receiver.port)
	if err != nil {
		logrus.WithFields(logrus.Fields{"port": receiver.port}).Error("http发送错误")
	}
}
