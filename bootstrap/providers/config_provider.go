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
}

func (c *ConfigProvider) Init() {
	c.data = make(map[string]*services.Config)

	c.initFlag()
	c.initFile()
}

func (c *ConfigProvider) initFlag() {
	flag.StringVar(&c.path, "env", "./.env", "加载配置文件")
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
	} else {
		_ = godotenv.Load(c.path)
		DirEntry, err = defaultConfigDir.ReadDir(defaultDir)
		if err != nil {
			panic(err)
		}
		for _, entry := range DirEntry {
			if path.Ext(entry.Name()) == ".yaml" {
				fileContext, err := defaultConfigDir.ReadFile(defaultDir + "/" + entry.Name())
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
}

func (c *ConfigProvider) Boot() {
	if !flag.Parsed() {
		flag.Parse()
	}
	services.Init = true
}

// GetBean 约定大于一切, 自己接收的代码和配置结构要人工约束成一致
func (c *ConfigProvider) GetBean(alias string) interface{} {
	index := strings.Index(alias, ".")
	if index == -1 {
		file, ok := c.data[alias]
		if !ok {
			file = services.NewConfig(make(map[interface{}]interface{}))
		}
		return file
	}

	fileConfig, ok := c.data[alias[:index]]
	if !ok {
		return nil
	}
	arr := strings.Split(alias[index+1:], ".")
	m := fileConfig.M
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := arr[i]
		if v, ok := m[s]; ok {
			if ll == i {
				val, ook := v.(map[interface{}]interface{})
				if ook {
					return services.NewConfig(val)
				} else {
					switch v.(type) {
					case int:
						got := v.(int)
						return &got
					case uint:
						got := v.(uint)
						return &got
					case bool:
						got := v.(bool)
						return &got
					case string:
						got := v.(string)
						return &got
					}
					return v
				}
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

	return nil
}
