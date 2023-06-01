package providers

import (
	"embed"
	"flag"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var envPath string

func init() {
	flag.StringVar(&envPath, "env", "./.env", "加载配置文件")
}

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
	if !flag.Parsed() {
		flag.Parse()
	}

	c.data = make(map[string]*services.Config)
	c.path = envPath
	c.initFile()
}

func (c *ConfigProvider) initFile() {
	var DirEntry []fs.DirEntry
	var err error
	if defaultConfigDir == nil {
		// 单元测试中, 可能未初始化框架, 从本目录开始往上查找go.mod文件确定跟目录
		pwd, _ := os.Getwd()
		parDir := ""
		for i := 0; i <= 100; i++ {
			checkDir := pwd + parDir
			_, err1 := os.Stat(checkDir + "/go.mod")
			_, err2 := os.Stat(checkDir + "/.env")
			if err1 == nil || err2 == nil {
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

		if _FrameworkProviderSingle == nil {
			NewFrameworkProvider()
		}
	} else {
		if _, err := os.Stat(c.path); err != nil {
			log.Error(".env file, " + err.Error())
		} else {
			err = godotenv.Load(c.path)
			if err != nil {
				panic("godotenv.Load error: " + err.Error())
			}
		}

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
	services.Init = true
}

// GetBean 约定大于一切, 自己接收的代码和配置结构要人工约束成一致
// , 后面的字符作为默认值, 默认值只能是字符串返回
func (c *ConfigProvider) GetBean(alias string) interface{} {
	index := strings.Index(alias, ",")
	var aliasDef string
	if index != -1 {
		aliasKey := alias[:index]
		aliasDef = strings.Trim(alias[index+1:], " ")
		alias = aliasKey
	}

	index = strings.Index(alias, ".")
	if index == -1 {
		file, ok := c.data[alias]
		if !ok {
			file = services.NewConfig(make(map[interface{}]interface{}))
		}
		return file
	}

	fileConfig, ok := c.data[alias[:index]]
	if !ok {
		if index == -1 {
			var aliasDef interface{}
			// 不存在分隔符=没有默认值
			return &aliasDef
		}
		return &aliasDef
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
				return &aliasDef
			}
		} else {
			return &aliasDef
		}
	}

	return &aliasDef
}
