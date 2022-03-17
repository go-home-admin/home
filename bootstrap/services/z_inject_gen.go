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
		_ContainerSingle = &Container{}
		app.AfterProvider(_ContainerSingle, "")
	}
	return _ContainerSingle
}
func NewHttpServer() *HttpServer {
	if _HttpServerSingle == nil {
		_HttpServerSingle = &HttpServer{}
		app.AfterProvider(_HttpServerSingle, "")
	}
	return _HttpServerSingle
}
