// gen for home toolset
package broker

import ()

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
