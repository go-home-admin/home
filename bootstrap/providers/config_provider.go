package providers

import (
	"embed"
	"flag"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/utils"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"path"
	"path/filepath"
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
	var DirEntry []fs.DirEntry
	var err error
	if defaultConfigDir == nil {
		// 单元测试中, 可能未初始化框架, 从本目录开始往上查找go.mod文件确定跟目录
		pwd, _ := os.Getwd()
		pwdArr := strings.Split(pwd, "/")
		parDir := ""
		for i := 0; i < len(pwdArr)-1; i++ {
			checkDir := pwd + parDir
			if _, err := os.Stat(checkDir + "/go.mod"); err == nil {
				dirs, _ := filepath.Abs(checkDir + "/" + defaultDir)
				DirEntry, err = os.ReadDir(dirs)
				if err != nil {
					panic(err)
				}
				// checkDir as root
				_ = godotenv.Load(checkDir + "/.env")
				defaultDir = checkDir + "/" + defaultDir
				break
			}
			parDir += "/.."
		}
		if _FrameworkProviderSingle == nil {
			NewFrameworkProvider()
		}
	} else {
		_ = godotenv.Load()
		DirEntry, err = defaultConfigDir.ReadDir(defaultDir)
		if err != nil {
			panic(err)
		}
	}

	for _, entry := range DirEntry {
		if path.Ext(entry.Name()) == ".yaml" {
			fileContext, err := os.ReadFile(defaultDir + "/" + entry.Name())
			if err != nil {
				panic(err)
			}
			fileContext = utils.SetEnv(fileContext)
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
	services.Init = true
}

func (c *ConfigProvider) GetBean(alias string) interface{} {
	index := strings.Index(alias, ".")
	if index == -1 {
		return c.data[alias]
	}

	fileConfig, ok := c.data[alias[:index]]
	if !ok {
		return nil
	}
	key := alias[index+1:]
	return fileConfig.GetConfig(key)
}
