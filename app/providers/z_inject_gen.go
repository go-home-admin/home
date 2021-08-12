// 代码由home-admin生成, 不需要编辑它

package providers

var ConfigSingle *Config
var IniSingle *Ini
var ResponseSingle *Response

func InitializeNewConfigProvider() *Config {
	if ConfigSingle == nil {
		ConfigSingle = NewConfigProvider(
			InitializeNewIniProvider(),
		)
	}

	return ConfigSingle
}

func InitializeNewIniProvider() *Ini {
	if IniSingle == nil {
		IniSingle = NewIniProvider()
	}

	return IniSingle
}

func InitializeNewResponseProvider() *Response {
	if ResponseSingle == nil {
		ResponseSingle = NewResponseProvider()
	}

	return ResponseSingle
}
