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

		var temp interface{} = AppSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return AppSingle
}

func InitializeNewConfigProvider() *Config {
	if ConfigSingle == nil {
		ConfigSingle = NewConfigProvider(
			InitializeNewIniProvider(),
		)

		var temp interface{} = ConfigSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
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

		var temp interface{} = IniSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return IniSingle
}

func NewLogProvider(ginLog *logs.GinLogrus) *Log {
	Log := &Log{}
	Log.ginLog = ginLog
	return Log
}

func InitializeNewLogProvider() *Log {
	if LogSingle == nil {
		LogSingle = NewLogProvider(
			logs.InitializeNewGinLogrusProvider(),
		)

		var temp interface{} = LogSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return LogSingle
}

func NewResponseProvider() *Response {
	Response := &Response{}
	return Response
}

func InitializeNewResponseProvider() *Response {
	if ResponseSingle == nil {
		ResponseSingle = NewResponseProvider()

		var temp interface{} = ResponseSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return ResponseSingle
}
