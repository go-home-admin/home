package message

import (
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/services/broker"
)

// 通知其他服务的事件，而无需担心响应
// @Bean
type Kernel struct {
	b     *broker.RedisBroker `inject:""`
	redis *providers.Redis    `inject:""`
}

func (k *Kernel) Init() {
	// 注入信息通道代理商
	k.b.SetConfig(k.redis)
}

func (k *Kernel) Push(event interface{}) {
	k.b.Push(event)
}

// 事件使用这个推送
func Push(event interface{}) {
	InitializeNewKernelProvider().Push(event)
}
