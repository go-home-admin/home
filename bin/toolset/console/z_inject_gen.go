// gen for home toolset
package console

var _KernelSingle *Kernel

func GetAllProvider() []interface{} {
	return []interface{}{
		NewKernel(),
	}
}

func NewKernel() *Kernel {
	if _KernelSingle == nil {
		Kernel := &Kernel{}
		_KernelSingle = Kernel
	}
	return _KernelSingle
}
