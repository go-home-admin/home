// gen for home toolset
package services

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
		_ContainerSingle = Container
	}
	return _ContainerSingle
}
func NewHttpServer() *HttpServer {
	if _HttpServerSingle == nil {
		HttpServer := &HttpServer{}
		_HttpServerSingle = HttpServer
	}
	return _HttpServerSingle
}
