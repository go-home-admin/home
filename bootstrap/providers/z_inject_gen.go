// gen for home toolset
package providers

import (
	services "github.com/go-home-admin/home/bootstrap/services"
)

var _ConfigProviderSingle *ConfigProvider
var _DatabaseProviderSingle *DatabaseProvider
var _FrameworkProviderSingle *FrameworkProvider
var _LogProviderSingle *LogProvider
var _MysqlProviderSingle *MysqlProvider
var _RedisProviderSingle *RedisProvider
var _RouteProviderSingle *RouteProvider

func GetAllProvider() []interface{} {
	return []interface{}{
		NewConfigProvider(),
		NewDatabaseProvider(),
		NewFrameworkProvider(),
		NewLogProvider(),
		NewMysqlProvider(),
		NewRedisProvider(),
		NewRouteProvider(),
	}
}

func NewConfigProvider() *ConfigProvider {
	if _ConfigProviderSingle == nil {
		_ConfigProviderSingle = &ConfigProvider{}
		AfterProvider(_ConfigProviderSingle, "config")
	}
	return _ConfigProviderSingle
}
func NewDatabaseProvider() *DatabaseProvider {
	if _DatabaseProviderSingle == nil {
		_DatabaseProviderSingle = &DatabaseProvider{}
		_DatabaseProviderSingle.Config = GetBean("config").(Bean).GetBean("database").(*services.Config)
		AfterProvider(_DatabaseProviderSingle, "database")
	}
	return _DatabaseProviderSingle
}
func NewFrameworkProvider() *FrameworkProvider {
	if _FrameworkProviderSingle == nil {
		_FrameworkProviderSingle = &FrameworkProvider{}
		_FrameworkProviderSingle.config = NewConfigProvider()
		_FrameworkProviderSingle.database = NewDatabaseProvider()
		_FrameworkProviderSingle.log = NewLogProvider()
		AfterProvider(_FrameworkProviderSingle, "")
	}
	return _FrameworkProviderSingle
}
func NewLogProvider() *LogProvider {
	if _LogProviderSingle == nil {
		_LogProviderSingle = &LogProvider{}
		_LogProviderSingle.Config = GetBean("config").(Bean).GetBean("app").(*services.Config)
		AfterProvider(_LogProviderSingle, "")
	}
	return _LogProviderSingle
}
func NewMysqlProvider() *MysqlProvider {
	if _MysqlProviderSingle == nil {
		_MysqlProviderSingle = &MysqlProvider{}
		_MysqlProviderSingle.config = GetBean("config").(Bean).GetBean("database").(*services.Config)
		AfterProvider(_MysqlProviderSingle, "mysql")
	}
	return _MysqlProviderSingle
}
func NewRedisProvider() *RedisProvider {
	if _RedisProviderSingle == nil {
		_RedisProviderSingle = &RedisProvider{}
		_RedisProviderSingle.Config = GetBean("config").(Bean).GetBean("database").(*services.Config)
		AfterProvider(_RedisProviderSingle, "redis")
	}
	return _RedisProviderSingle
}
func NewRouteProvider() *RouteProvider {
	if _RouteProviderSingle == nil {
		_RouteProviderSingle = &RouteProvider{}
		AfterProvider(_RouteProviderSingle, "")
	}
	return _RouteProviderSingle
}
