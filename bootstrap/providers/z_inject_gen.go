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

func GetAllProvider() []interface{} {
	return []interface{}{
		NewConfigProvider(),
		NewDatabaseProvider(),
		NewFrameworkProvider(),
		NewMysqlProvider(),
	}
}

func NewConfigProvider() *ConfigProvider {
	if _ConfigProviderSingle == nil {
		_ConfigProviderSingle = &ConfigProvider{}
		app.AfterProvider(_ConfigProviderSingle, "config")
	}
	return _ConfigProviderSingle
}
func NewDatabaseProvider() *DatabaseProvider {
	if _DatabaseProviderSingle == nil {
		_DatabaseProviderSingle = &DatabaseProvider{}
		_DatabaseProviderSingle.config = *app.GetBean("config").(app.Bean).GetBean("database").(*services.Config)
		app.AfterProvider(_DatabaseProviderSingle, "database")
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
		_MysqlProviderSingle.config = app.GetBean("config").(app.Bean).GetBean("database").(*services.Config)
		app.AfterProvider(_MysqlProviderSingle, "mysql")
	}
	return _MysqlProviderSingle
}
