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

func NewWorker() *Worker {
	if _WorkerSingle == nil {
		Worker := &Worker{}
		Worker.redis = providers.NewRedis()
		app.AfterProvider(Worker, "")
		_WorkerSingle = Worker
	}
	return _WorkerSingle
}
func NewKernel() *Kernel {
	if _KernelSingle == nil {
		Kernel := &Kernel{}
		Kernel.b = broker.NewRedisBroker()
		Kernel.redis = providers.NewRedis()
		Kernel.worker = NewWorker()
		app.AfterProvider(Kernel, "")
		_KernelSingle = Kernel
	}
	return _KernelSingle
}
