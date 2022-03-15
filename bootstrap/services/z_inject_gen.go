// gen for home toolset
package services

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _ContainerSingle *Container
var _HttpServerSingle *HttpServer

func GetAllProvider() []interface{} {
	return []interface{}{
		NewContainer(),
		NewHttpServer(),
	}
}

func NewContainer() *Container {
	if _ContainerSingle == nil {
		Container := &Container{}
		app.AfterProvider(Container, "")
		_ContainerSingle = Container
	}
	return _ContainerSingle
}
func NewHttpServer() *HttpServer {
	if _HttpServerSingle == nil {
		HttpServer := &HttpServer{}
		app.AfterProvider(HttpServer, "")
		_HttpServerSingle = HttpServer
	}
	return _HttpServerSingle
}
