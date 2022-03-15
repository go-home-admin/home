// gen for home toolset
package providers

import (
	services "github.com/go-home-admin/home/bootstrap/services"
	logs "github.com/go-home-admin/home/bootstrap/services/logs"
)

var _IniSingle *Ini
var _LogSingle *Log
var _MysqlSingle *Mysql
var _RedisSingle *Redis
var _ResponseSingle *Response
var _AppSingle *App
var _ConfigSingle *Config

func GetAllProvider() []interface{} {
	return []interface{}{
		NewIni(),
		NewLog(),
		NewMysql(),
		NewRedis(),
		NewResponse(),
		NewApp(),
		NewConfig(),
	}
}

func NewRedis() *Redis {
	if _RedisSingle == nil {
		Redis := &Redis{}
		Redis.conf = NewConfig()
		app.AfterProvider(Redis, "")
		_RedisSingle = Redis
	}
	return _RedisSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		Response := &Response{}
		app.AfterProvider(Response, "")
		_ResponseSingle = Response
	}
	return _ResponseSingle
}
func NewApp() *App {
	if _AppSingle == nil {
		App := &App{}
		App.container = services.NewContainer()
		App.resp = NewResponse()
		App.log = NewLog()
		app.AfterProvider(App, "")
		_AppSingle = App
	}
	return _AppSingle
}
func NewConfig() *Config {
	if _ConfigSingle == nil {
		Config := &Config{}
		Config.iniConfig = NewIni()
		app.AfterProvider(Config, "")
		_ConfigSingle = Config
	}
	return _ConfigSingle
}
func NewIni() *Ini {
	if _IniSingle == nil {
		Ini := &Ini{}
		app.AfterProvider(Ini, "")
		_IniSingle = Ini
	}
	return _IniSingle
}
func NewLog() *Log {
	if _LogSingle == nil {
		Log := &Log{}
		Log.ginLog = logs.NewGinLogrus()
		Log.conf = NewConfig()
		app.AfterProvider(Log, "")
		_LogSingle = Log
	}
	return _LogSingle
}
func NewMysql() *Mysql {
	if _MysqlSingle == nil {
		Mysql := &Mysql{}
		Mysql.conf = NewConfig()
		Mysql.log = NewLog()
		app.AfterProvider(Mysql, "mysql")
		_MysqlSingle = Mysql
	}
	return _MysqlSingle
}
