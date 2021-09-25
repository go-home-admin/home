// 代码由home-admin生成, 不需要编辑它

package queue

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var KernelSingle *Kernel
var StreamSingle *Stream

func NewKernelProvider(def *Stream) *Kernel {
	Kernel := &Kernel{}
	Kernel.def = def
	return Kernel
}

func InitializeNewKernelProvider() *Kernel {
	if KernelSingle == nil {
		KernelSingle = NewKernelProvider(
			InitializeNewStreamProvider(),
		)

		home_constraint.AfterProvider(KernelSingle)
	}

	return KernelSingle
}

func NewStreamProvider() *Stream {
	Stream := &Stream{}
	return Stream
}

func InitializeNewStreamProvider() *Stream {
	if StreamSingle == nil {
		StreamSingle = NewStreamProvider()

		home_constraint.AfterProvider(StreamSingle)
	}

	return StreamSingle
}
