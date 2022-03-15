// gen for home toolset
package message

import (
	providers "github.com/go-home-admin/home/app/providers"
	broker "github.com/go-home-admin/home/bootstrap/services/broker"
)

var _KernelSingle *Kernel

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		Kernel := &Kernel{}
		Kernel.b = broker.NewRedisBroker()
		Kernel.redis = providers.NewRedis()
		app.AfterProvider(Kernel, "")
		_KernelSingle = Kernel
	}
	return _KernelSingle
}
