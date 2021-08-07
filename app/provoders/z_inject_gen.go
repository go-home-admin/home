// 代码由home-admin生成, 不需要编辑它

package provoders

var ConfigSingle *Config

func NewConfigProvider() *Config {
	Config := &Config{}
	return Config
}

func InitializeNewConfigProvider() *Config {
	if ConfigSingle == nil {
		ConfigSingle = NewConfigProvider()
	}

	return ConfigSingle
}
