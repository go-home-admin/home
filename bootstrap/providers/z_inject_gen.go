// gen for home toolset
package providers

import (
	services "github.com/go-home-admin/home/bootstrap/services"
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _ConfigProviderSingle *ConfigProvider
var _DatabaseProviderSingle *DatabaseProvider
var _FrameworkProviderSingle *FrameworkProvider
var _MysqlProviderSingle *MysqlProvider
var _RouteProviderSingle *RouteProvider

func GetAllProvider() []interface{} {
	return []interface{}{
		NewConfigProvider(),
		NewDatabaseProvider(),
		NewFrameworkProvider(),
		NewMysqlProvider(),
		NewRouteProvider(),
	}
}

func NewConfigProvider() *ConfigProvider {
	if _ConfigProviderSingle == nil {
		_ConfigProviderSingle = &ConfigProvider{}
		app.AfterProvider(_ConfigProviderSingle, "")
	}
	return _ConfigProviderSingle
}
func NewDatabaseProvider() *DatabaseProvider {
	if _DatabaseProviderSingle == nil {
		_DatabaseProviderSingle = &DatabaseProvider{}
		app.AfterProvider(_DatabaseProviderSingle, "")
	}
	return _DatabaseProviderSingle
}
func NewFrameworkProvider() *FrameworkProvider {
	if _FrameworkProviderSingle == nil {
		_FrameworkProviderSingle = &FrameworkProvider{}
		_FrameworkProviderSingle.config = NewConfigProvider()
		_FrameworkProviderSingle.database = NewDatabaseProvider()
		app.AfterProvider(_FrameworkProviderSingle, "")
	}
	return _FrameworkProviderSingle
}
func NewMysqlProvider() *MysqlProvider {
	if _MysqlProviderSingle == nil {
		_MysqlProviderSingle = &MysqlProvider{}
		_MysqlProviderSingle.config = app.GetBean("Route").(app.Bean).GetBean("database").(*services.Config)
		app.AfterProvider(_MysqlProviderSingle, "")
	}
	return _MysqlProviderSingle
}
func NewRouteProvider() *RouteProvider {
	if _RouteProviderSingle == nil {
		_RouteProviderSingle = &RouteProvider{}
		app.AfterProvider(_RouteProviderSingle, "")
	}
	return _RouteProviderSingle
}
