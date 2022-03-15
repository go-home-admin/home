// gen for home toolset
package console

import ()

var _KernelSingle *Kernel

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		Kernel := &Kernel{}
		app.AfterProvider(Kernel, "")
		_KernelSingle = Kernel
	}
	return _KernelSingle
}
