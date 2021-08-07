// 代码由home-admin生成, 不需要编辑它

package services

var AppServerSingle *AppServer
var HttpServerSingle *HttpServer

func NewAppServerProvider() *AppServer {
	AppServer := &AppServer{}
	return AppServer
}

func InitializeNewAppServerProvider() *AppServer {
	if AppServerSingle == nil {
		AppServerSingle = NewAppServerProvider()
	}

	return AppServerSingle
}

func InitializeNewHttpServerProvider() *HttpServer {
	if HttpServerSingle == nil {
		HttpServerSingle = NewHttpServerProvider()
	}

	return HttpServerSingle
}
