package services

import "github.com/gin-gonic/gin"

// 封装一些被(工具生成的)代码调用的函数
// @Bean
type HttpServer struct {
	isDebug bool
	engine  *gin.Engine
}

func NewHttpServerProvider() *HttpServer {
	HttpServer := &HttpServer{
		isDebug: false,
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

func (receiver *HttpServer) RunListener(port int, host ...string) {

}
