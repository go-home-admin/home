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

		var temp interface{} = ContainerSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return ContainerSingle
}

func InitializeNewHttpServerProvider() *HttpServer {
	if HttpServerSingle == nil {
		HttpServerSingle = NewHttpServerProvider()

		var temp interface{} = HttpServerSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return HttpServerSingle
}
