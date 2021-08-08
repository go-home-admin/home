package services

import (
	"github.com/gin-gonic/gin"
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

func NewHttpServerProvider() *HttpServer {
	HttpServer := &HttpServer{
		isDebug: false,
		port:    "80",
		host:    "",
		engine:  gin.New(),
	}
	return HttpServer
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
	receiver.GetEngine().Run(":" + receiver.port)
}
