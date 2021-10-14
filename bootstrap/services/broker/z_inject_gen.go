// 代码由home-admin生成, 不需要编辑它

package broker

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var RedisBrokerSingle *RedisBroker

func NewRedisBrokerProvider() *RedisBroker {
	RedisBroker := &RedisBroker{}
	return RedisBroker
}

func InitializeNewRedisBrokerProvider() *RedisBroker {
	if RedisBrokerSingle == nil {
		RedisBrokerSingle = NewRedisBrokerProvider()

		home_constraint.AfterProvider(RedisBrokerSingle)
	}

	return RedisBrokerSingle
}
