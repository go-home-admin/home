// gen for home toolset
package servers

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
)

var _CrontabSingle *Crontab
var _ElectionSingle *Election
var _HttpSingle *Http
var _QueueSingle *Queue
var _WebsocketSingle *Websocket

func GetAllProvider() []interface{} {
	return []interface{}{
		NewCrontab(),
		NewElection(),
		NewHttp(),
		NewQueue(),
		NewWebsocket(),
	}
}

func NewCrontab() *Crontab {
	if _CrontabSingle == nil {
		_CrontabSingle = &Crontab{}
		providers.AfterProvider(_CrontabSingle, "")
	}
	return _CrontabSingle
}
func NewElection() *Election {
	if _ElectionSingle == nil {
		_ElectionSingle = &Election{}
		_ElectionSingle.Config = providers.GetBean("config").(providers.Bean).GetBean("election").(*services.Config)
		providers.AfterProvider(_ElectionSingle, "election")
	}
	return _ElectionSingle
}
func NewHttp() *Http {
	if _HttpSingle == nil {
		_HttpSingle = &Http{}
		_HttpSingle.RouteProvider = providers.NewRouteProvider()
		_HttpSingle.HttpServer = services.NewHttpServer()
		_HttpSingle.Config = providers.GetBean("config").(providers.Bean).GetBean("app.servers.http").(*services.Config)
		providers.AfterProvider(_HttpSingle, "http")
	}
	return _HttpSingle
}
func NewQueue() *Queue {
	if _QueueSingle == nil {
		_QueueSingle = &Queue{}
		_QueueSingle.fileConfig = providers.GetBean("config").(providers.Bean).GetBean("queue").(*services.Config)
		providers.AfterProvider(_QueueSingle, "queue")
	}
	return _QueueSingle
}
func NewWebsocket() *Websocket {
	if _WebsocketSingle == nil {
		_WebsocketSingle = &Websocket{}
		providers.AfterProvider(_WebsocketSingle, "")
	}
	return _WebsocketSingle
}
