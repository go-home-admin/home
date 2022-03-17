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
		_RedisBrokerSingle = &RedisBroker{}
		app.AfterProvider(_RedisBrokerSingle, "")
	}
	return _RedisBrokerSingle
}
