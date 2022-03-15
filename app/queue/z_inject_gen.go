// gen for home toolset
package queue

import (
	providers "github.com/go-home-admin/home/app/providers"
	app "github.com/go-home-admin/home/bootstrap/services/app"
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
		Kernel.redis = providers.NewRedis()
		Kernel.worker = NewWorker()
		Kernel.b = broker.NewRedisBroker()
		app.AfterProvider(Kernel, "")
		_KernelSingle = Kernel
	}
	return _KernelSingle
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
