package providers

import (
	"flag"
	"github.com/go-home-admin/home/bootstrap/services"
)

// @Bean("config")
type ConfigProvider struct {
	data map[string]*services.Config `inject:""`

	path string
	port string
}

func (c *ConfigProvider) initFlag() {
	flag.StringVar(&c.path, "env", "./.env", "加载配置文件")
	flag.StringVar(&c.port, "port", "8080", "http端口")

	if !flag.Parsed() {
		flag.Parse()
	}
}

func (c *ConfigProvider) Init() {
	c.data = make(map[string]*services.Config)

	c.initFlag()
}

func (c *ConfigProvider) GetBean(alias string) *services.Config {
	return c.data[alias]
}
