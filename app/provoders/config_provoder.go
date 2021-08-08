package provoders

// Config 外部其他服务的配置依赖提供
// @Bean
type Config struct {
	iniConfig *Ini `inject:""`
}

func NewConfigProvider(ini *Ini) *Config {
	Config := &Config{}
	Config.Init()
	return Config
}

func (g *Config) Init() {

}

func (g *Config) GetString() string {
	return ""
}
