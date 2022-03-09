// gen for home toolset
package providers

import (
	services "github.com/go-home-admin/home/bootstrap/services"
	logs "github.com/go-home-admin/home/bootstrap/services/logs"
)

var _ResponseSingle *Response
var _AppSingle *App
var _ConfigSingle *Config
var _IniSingle *Ini
var _LogSingle *Log
var _MysqlSingle *Mysql
var _RedisSingle *Redis

func GetAllProvider() []interface{} {
	return []interface{}{
		NewResponse(),
		NewApp(),
		NewConfig(),
		NewIni(),
		NewLog(),
		NewMysql(),
		NewRedis(),
	}
}

func NewLog() *Log {
	if _LogSingle == nil {
		Log := &Log{}
		Log.ginLog = logs.NewGinLogrus()
		Log.conf = .New*Config()
		_LogSingle = Log
	}
	return _LogSingle
}
func NewMysql() *Mysql {
	if _MysqlSingle == nil {
		Mysql := &Mysql{}
		Mysql.conf = .New*Config()
		Mysql.log = .New*Log()
		_MysqlSingle = Mysql
	}
	return _MysqlSingle
}
func NewRedis() *Redis {
	if _RedisSingle == nil {
		Redis := &Redis{}
		Redis.conf = .New*Config()
		_RedisSingle = Redis
	}
	return _RedisSingle
}
func NewResponse() *Response {
	if _ResponseSingle == nil {
		Response := &Response{}
		_ResponseSingle = Response
	}
	return _ResponseSingle
}
func NewApp() *App {
	if _AppSingle == nil {
		App := &App{}
		App.container = services.NewContainer()
		App.resp = .New*Response()
		App.log = .New*Log()
		_AppSingle = App
	}
	return _AppSingle
}
func NewConfig() *Config {
	if _ConfigSingle == nil {
		Config := &Config{}
		Config.iniConfig = .New*Ini()
		_ConfigSingle = Config
	}
	return _ConfigSingle
}
func NewIni() *Ini {
	if _IniSingle == nil {
		Ini := &Ini{}
		_IniSingle = Ini
	}
	return _IniSingle
}