// gen for home toolset
package servers

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
	services "github.com/go-home-admin/home/bootstrap/services"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _HttpSingle *Http

func GetAllProvider() []interface{} {
	return []interface{}{
		NewHttp(),
	}
}

func NewHttp() *Http {
	if _HttpSingle == nil {
		_HttpSingle = &Http{}
		_HttpSingle.RouteProvider = providers.NewRouteProvider()
		_HttpSingle.HttpServer = services.NewHttpServer()
		_HttpSingle.Config = app.GetBean("config").(app.Bean).GetBean("app.servers.http").(*services.Config)
		app.AfterProvider(_HttpSingle, "http")
	}
	return _HttpSingle
}
