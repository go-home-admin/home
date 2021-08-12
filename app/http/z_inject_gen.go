// 代码由home-admin生成, 不需要编辑它

package http

import (
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/services"
)

var KernelSingle *Kernel

func NewKernelProvider(httpServer *services.HttpServer, config *providers.Config) *Kernel {
	Kernel := &Kernel{}
	Kernel.httpServer = httpServer
	Kernel.config = config
	return Kernel
}

func InitializeNewKernelProvider() *Kernel {
	if KernelSingle == nil {
		KernelSingle = NewKernelProvider(
			services.InitializeNewHttpServerProvider(),

			providers.InitializeNewConfigProvider(),
		)
	}

	return KernelSingle
}
