// gen for home toolset
package queue

import (
	providers "github.com/go-home-admin/home/app/providers"
	broker "github.com/go-home-admin/home/bootstrap/services/broker"
)

var _KernelSingle *Kernel
var _WorkerSingle *Worker

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
		NewWorker(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		Kernel := &Kernel{}
		Kernel.b = broker.NewRedisBroker()
		Kernel.redis = providers.NewRedis()
		Kernel.worker = .New*Worker()
		_KernelSingle = Kernel
	}
	return _KernelSingle
}
func NewWorker() *Worker {
	if _WorkerSingle == nil {
		Worker := &Worker{}
		Worker.redis = providers.NewRedis()
		_WorkerSingle = Worker
	}
	return _WorkerSingle
}