// 代码由home-admin生成, 不需要编辑它

package http

import (
	"github.com/go-home-admin/home/app/providers"
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/routes"
)

var KernelSingle *Kernel

func NewKernelProvider(routes *routes.Routes, httpServer *services.HttpServer, config *providers.Config) *Kernel {
	Kernel := &Kernel{}
	Kernel.routes = routes
	Kernel.httpServer = httpServer
	Kernel.config = config
	return Kernel
}

func InitializeNewKernelProvider() *Kernel {
	if KernelSingle == nil {
		KernelSingle = NewKernelProvider(
			routes.InitializeNewRoutesProvider(),

			services.InitializeNewHttpServerProvider(),

			providers.InitializeNewConfigProvider(),
		)

		home_constraint.AfterProvider(KernelSingle)
	}

	return KernelSingle
}
