// gen for home toolset
package servers

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _CrontabSingle *Crontab
var _HttpSingle *Http
var _QueueSingle *Queue
var _WebsocketSingle *Websocket

func GetAllProvider() []interface{} {
	return []interface{}{
		NewCrontab(),
		NewHttp(),
		NewQueue(),
		NewWebsocket(),
	}
}

func NewCrontab() *Crontab {
	if _CrontabSingle == nil {
		_CrontabSingle = &Crontab{}
		app.AfterProvider(_CrontabSingle, "")
	}
	return _CrontabSingle
}
func NewHttp() *Http {
	if _HttpSingle == nil {
		_HttpSingle = &Http{}
		_HttpSingle.RouteProvider = providers.NewRouteProvider()
		_HttpSingle.HttpServer = services.NewHttpServer()
		_HttpSingle.Config = app.GetBean("config").(app.Bean).GetBean("app.servers.http").(*services.Config)
		_HttpSingle.TPort = *app.GetBean("config").(app.Bean).GetBean("app.servers.http.port").(*int)
		app.AfterProvider(_HttpSingle, "http")
	}
	return _HttpSingle
}
func NewQueue() *Queue {
	if _QueueSingle == nil {
		_QueueSingle = &Queue{}
		_QueueSingle.queue = app.GetBean("config").(app.Bean).GetBean("queue").(*services.Config)
		app.AfterProvider(_QueueSingle, "")
	}
	return _QueueSingle
}
func NewWebsocket() *Websocket {
	if _WebsocketSingle == nil {
		_WebsocketSingle = &Websocket{}
		app.AfterProvider(_WebsocketSingle, "")
	}
	return _WebsocketSingle
}
