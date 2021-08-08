// 代码由home-admin生成, 不需要编辑它

package http

import (
	provoders "github.com/go-home-admin/home/app/provoders"
	"github.com/go-home-admin/home/bootstrap/services"
)

var KernelSingle *Kernel

func NewKernelProvider(httpServer *services.HttpServer, config *provoders.Config) *Kernel {
	Kernel := &Kernel{}
	Kernel.httpServer = httpServer
	Kernel.config = config
	return Kernel
}

func InitializeNewKernelProvider() *Kernel {
	if KernelSingle == nil {
		KernelSingle = NewKernelProvider(
			services.InitializeNewHttpServerProvider(),
			provoders.InitializeNewConfigProvider(),
		)
	}

	return KernelSingle
}
