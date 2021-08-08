// 代码由home-admin生成, 不需要编辑它

package provoders

var ConfigSingle *Config
var IniSingle *Ini

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
