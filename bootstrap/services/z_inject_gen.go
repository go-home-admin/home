// 代码由home-admin生成, 不需要编辑它

package services

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var ContainerSingle *Container
var HttpServerSingle *HttpServer

func NewContainerProvider() *Container {
	Container := &Container{}
	return Container
}

func InitializeNewContainerProvider() *Container {
	if ContainerSingle == nil {
		ContainerSingle = NewContainerProvider()

		home_constraint.AfterProvider(ContainerSingle)
	}

	return ContainerSingle
}

func NewHttpServerProvider() *HttpServer {
	HttpServer := &HttpServer{}
	return HttpServer
}

func InitializeNewHttpServerProvider() *HttpServer {
	if HttpServerSingle == nil {
		HttpServerSingle = NewHttpServerProvider()

		home_constraint.AfterProvider(HttpServerSingle)
	}

	return HttpServerSingle
}
