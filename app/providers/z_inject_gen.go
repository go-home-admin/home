// 代码由home-admin生成, 不需要编辑它

package providers

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/logs"
)

var AppSingle *App
var ConfigSingle *Config
var IniSingle *Ini
var LogSingle *Log
var MysqlSingle *Mysql
var RedisSingle *Redis
var ResponseSingle *Response

func NewAppProvider(container *services.Container, resp *Response, log *Log) *App {
	App := &App{}
	App.container = container
	App.resp = resp
	App.log = log
	return App
}

func InitializeNewAppProvider() *App {
	if AppSingle == nil {
		AppSingle = NewAppProvider(
			services.InitializeNewContainerProvider(),

			InitializeNewResponseProvider(),

			InitializeNewLogProvider(),
		)

		home_constraint.AfterProvider(AppSingle)
	}

	return AppSingle
}

func NewConfigProvider(iniConfig *Ini) *Config {
	Config := &Config{}
	Config.iniConfig = iniConfig
	return Config
}

func InitializeNewConfigProvider() *Config {
	if ConfigSingle == nil {
		ConfigSingle = NewConfigProvider(
			InitializeNewIniProvider(),
		)

		home_constraint.AfterProvider(ConfigSingle)
	}

	return ConfigSingle
}

func NewIniProvider() *Ini {
	Ini := &Ini{}
	return Ini
}

func InitializeNewIniProvider() *Ini {
	if IniSingle == nil {
		IniSingle = NewIniProvider()

		home_constraint.AfterProvider(IniSingle)
	}

	return IniSingle
}

func NewLogProvider(ginLog *logs.GinLogrus, conf *Config) *Log {
	Log := &Log{}
	Log.ginLog = ginLog
	Log.conf = conf
	return Log
}

func InitializeNewLogProvider() *Log {
	if LogSingle == nil {
		LogSingle = NewLogProvider(
			logs.InitializeNewGinLogrusProvider(),

			InitializeNewConfigProvider(),
		)

		home_constraint.AfterProvider(LogSingle)
	}

	return LogSingle
}

func NewMysqlProvider(conf *Config) *Mysql {
	Mysql := &Mysql{}
	Mysql.conf = conf
	return Mysql
}

func InitializeNewMysqlProvider() *Mysql {
	if MysqlSingle == nil {
		MysqlSingle = NewMysqlProvider(
			InitializeNewConfigProvider(),
		)

		home_constraint.AfterProvider(MysqlSingle)
	}

	return MysqlSingle
}

func NewRedisProvider(conf *Config) *Redis {
	Redis := &Redis{}
	Redis.conf = conf
	return Redis
}

func InitializeNewRedisProvider() *Redis {
	if RedisSingle == nil {
		RedisSingle = NewRedisProvider(
			InitializeNewConfigProvider(),
		)

		home_constraint.AfterProvider(RedisSingle)
	}

	return RedisSingle
}

func NewResponseProvider() *Response {
	Response := &Response{}
	return Response
}

func InitializeNewResponseProvider() *Response {
	if ResponseSingle == nil {
		ResponseSingle = NewResponseProvider()

		home_constraint.AfterProvider(ResponseSingle)
	}

	return ResponseSingle
}
