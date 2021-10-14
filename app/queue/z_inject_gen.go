// 代码由home-admin生成, 不需要编辑它

package queue

import (
	"github.com/go-home-admin/home/app/providers"
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services/broker"
)

var KernelSingle *Kernel
var StreamSingle *Stream

func NewKernelProvider(b *broker.RedisBroker, redis *providers.Redis) *Kernel {
	Kernel := &Kernel{}
	Kernel.b = b
	Kernel.redis = redis
	return Kernel
}

func InitializeNewKernelProvider() *Kernel {
	if KernelSingle == nil {
		KernelSingle = NewKernelProvider(
			broker.InitializeNewRedisBrokerProvider(),
			providers.InitializeNewRedisProvider(),
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
