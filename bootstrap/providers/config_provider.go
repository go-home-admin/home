package providers

import (
	"embed"
	"flag"
	"github.com/go-home-admin/home/bootstrap/services"
	"gopkg.in/yaml.v2"
	"path"
	"strings"
)

// 默认配置加载目录
var defaultConfigDir *embed.FS
var defaultDir = "config"

func SetConfigDir(fs *embed.FS) {
	defaultConfigDir = fs
}

// ConfigProvider
// @Bean("config")
type ConfigProvider struct {
	data map[string]*services.Config

	path string
	port string
}

func (c *ConfigProvider) Init() {
	c.data = make(map[string]*services.Config)

	c.initFlag()
	c.initFile()
}

func (c *ConfigProvider) initFlag() {
	flag.StringVar(&c.path, "env", "./.env", "加载配置文件")
	flag.StringVar(&c.port, "port", "8080", "http端口")
}

func (c *ConfigProvider) initFile() {
	DirEntry, err := defaultConfigDir.ReadDir(defaultDir)
	if err != nil {
		panic(err)
	}
	for _, entry := range DirEntry {
		if path.Ext(entry.Name()) == ".yaml" {
			fileContext, _ := defaultConfigDir.ReadFile(defaultDir + "/" + entry.Name())
			m := make(map[interface{}]interface{})
			err = yaml.Unmarshal(fileContext, &m)
			if err != nil {
				panic(err)
			}
			c.data[strings.TrimSuffix(entry.Name(), ".yaml")] = services.NewConfig(m)
		}
	}
}

func (c *ConfigProvider) Boot() {
	if !flag.Parsed() {
		flag.Parse()
	}

	// 单元测试中, 可能未初始化框架
	if _FrameworkProviderSingle == nil {
		NewFrameworkProvider()
	}
}

func (c *ConfigProvider) GetBean(alias string) interface{} {
	return c.data[alias]
}
