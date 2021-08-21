package services

import (
	"github.com/gin-gonic/gin"
	logs "github.com/go-home-admin/home/bootstrap/services/logs"
	"strconv"
)

// HttpServer 封装一些被(工具生成的)代码调用的函数
// @Bean
type HttpServer struct {
	isDebug bool
	port    string
	host    string
	engine  *gin.Engine
}

func (receiver *HttpServer) Init() {
	receiver.isDebug = false
	receiver.port = "80"
	receiver.engine = gin.New()
}

func (receiver *HttpServer) SetDebug(isDebug bool) {
	receiver.isDebug = isDebug
}

func (receiver *HttpServer) GetEngine() *gin.Engine {
	return receiver.engine
}

func (receiver *HttpServer) SetPort(port int) {
	receiver.port = strconv.Itoa(port)
}

func (receiver *HttpServer) RunListener() {
	err := receiver.GetEngine().Run(":" + receiver.port)
	if err != nil {
		logs.Error(err)
	}
}
