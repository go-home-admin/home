// 代码由home-admin生成, 不需要编辑它

package providers

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var ConfigSingle *Config
var IniSingle *Ini
var ResponseSingle *Response

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
