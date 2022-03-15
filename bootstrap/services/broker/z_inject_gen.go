// gen for home toolset
package broker

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _RedisBrokerSingle *RedisBroker

func GetAllProvider() []interface{} {
	return []interface{}{
		NewRedisBroker(),
	}
}

func NewRedisBroker() *RedisBroker {
	if _RedisBrokerSingle == nil {
		RedisBroker := &RedisBroker{}
		app.AfterProvider(RedisBroker, "")
		_RedisBrokerSingle = RedisBroker
	}
	return _RedisBrokerSingle
}
