// 代码由home-admin生成, 不需要编辑它

package providers

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
)

var AppSingle *App
var ConfigSingle *Config
var IniSingle *Ini
var ResponseSingle *Response

func NewAppProvider(container *services.Container, resp *Response) *App {
	App := &App{}
	App.container = container
	App.resp = resp
	return App
}

func InitializeNewAppProvider() *App {
	if AppSingle == nil {
		AppSingle = NewAppProvider(
			services.InitializeNewContainerProvider(),

			InitializeNewResponseProvider(),
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
