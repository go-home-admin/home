// gen for home toolset
package broker

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _RedisBrokerSingle *RedisBroker

func GetAllProvider() []interface{} {
	return []interface{}{
		NewRedisBroker(),
	}
}

func NewRedisBroker() *RedisBroker {
	if _RedisBrokerSingle == nil {
		_RedisBrokerSingle = &RedisBroker{}
		providers.AfterProvider(_RedisBrokerSingle, "")
	}
	return _RedisBrokerSingle
}
