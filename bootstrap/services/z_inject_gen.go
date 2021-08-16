// 代码由home-admin生成, 不需要编辑它

package services

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var AppServerSingle *AppServer
var HttpServerSingle *HttpServer

func NewAppServerProvider() *AppServer {
	AppServer := &AppServer{}
	return AppServer
}

func InitializeNewAppServerProvider() *AppServer {
	if AppServerSingle == nil {
		AppServerSingle = NewAppServerProvider()

		var temp interface{} = AppServerSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return AppServerSingle
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
