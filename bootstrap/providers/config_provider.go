package providers

import (
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
					// panic(err) // 如果不是 web 项目, 允许没有.env 文件
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
					case int64:
						got := v.(int64)
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

var ROOT string = ""

func (c *ConfigProvider) GetRoot() string {
	if ROOT != "" {
		return ROOT
	}

	pwd, _ := os.Getwd()
	parDir := ""

	if defaultConfigDir == nil {
		// 单元测试中, 可能未初始化框架, 从本目录开始往上查找go.mod文件确定跟目录
		for i := 0; i <= 100; i++ {
			checkDir := pwd + parDir
			_, err1 := os.Stat(checkDir + "/go.mod")
			_, err2 := os.Stat(checkDir + "/.env")
			if err1 == nil || err2 == nil {
				parDir = checkDir
				break
			}
			parDir += "/.."
		}
	} else {
		parDir = pwd
	}
	ROOT = parDir
	return parDir
}

// InjectValue 将配置写入 dest。dest 必须为字段地址（toolset 生成 &field），与 GetBean 返回值配合用反射赋值。
func (c *ConfigProvider) InjectValue(alias string, dest interface{}) {
	c.writeInjectDest(alias, dest, c.GetBean(alias))
}

func (c *ConfigProvider) writeInjectDest(configKey string, dest, src interface{}) {
	if dest == nil {
		c.panicInjectAssign(configKey, "dest 为 nil")
	}
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.IsNil() {
		c.panicInjectAssign(configKey, "dest 必须为非 nil 的字段地址（生成代码请使用 &field）")
	}
	fieldVal := destVal.Elem()
	if !fieldVal.CanSet() {
		c.panicInjectAssign(configKey, "dest 指向的字段不可赋值（生成代码请使用 &field）")
	}

	if src != nil && c.tryInjectUnmarshaler(configKey, dest, fieldVal, src) {
		return
	}

	if src == nil {
		if fieldVal.Kind() == reflect.Ptr {
			fieldVal.Set(reflect.Zero(fieldVal.Type()))
			return
		}
		c.panicInjectAssign(configKey, "GetBean 返回 nil 且目标字段非指针类型")
	}

	srcVal := reflect.ValueOf(src)

	// GetBean 返回 *T，写入值类型字段 T（如 bool、int）
	if fieldVal.Kind() != reflect.Ptr && srcVal.Kind() == reflect.Ptr {
		if srcVal.IsNil() {
			fieldVal.Set(reflect.Zero(fieldVal.Type()))
			return
		}
		if configSetAssignable(fieldVal, srcVal.Elem()) {
			return
		}
		c.panicInjectAssign(configKey, "GetBean 指针与值类型字段不匹配")
	}

	// 目标为 *T：只写入元素值（复制），不把 GetBean 返回的指针地址赋给字段
	if fieldVal.Kind() == reflect.Ptr {
		if srcVal.Kind() == reflect.Ptr {
			if srcVal.IsNil() {
				fieldVal.Set(reflect.Zero(fieldVal.Type()))
				return
			}
			srcVal = srcVal.Elem()
		}
		elemTyp := fieldVal.Type().Elem()
		if fieldVal.IsNil() {
			newPtr := reflect.New(elemTyp)
			if configSetAssignable(newPtr.Elem(), srcVal) {
				fieldVal.Set(newPtr)
				return
			}
		} else if configSetAssignable(fieldVal.Elem(), srcVal) {
			return
		}
		c.panicInjectAssign(configKey, "GetBean 与指针字段元素类型不匹配")
	}

	if configSetAssignable(fieldVal, srcVal) {
		return
	}
	c.panicInjectAssign(configKey, "GetBean 值与字段类型不匹配")
}

// tryInjectUnmarshaler 当注入目标实现 json.Unmarshaler 且配置值为 JSON 文本（string / *string / []byte）时直接 UnmarshalJSON。
func (c *ConfigProvider) tryInjectUnmarshaler(configKey string, dest interface{}, fieldVal reflect.Value, src interface{}) bool {
	u := configJSONUnmarshalerTarget(dest, fieldVal)
	if u == nil {
		return false
	}
	data, ok := configJSONPayload(src)
	if !ok {
		return false
	}
	if err := u.UnmarshalJSON(data); err != nil {
		c.panicInjectAssign(configKey, "UnmarshalJSON: "+err.Error())
	}
	return true
}

func configJSONUnmarshalerTarget(dest interface{}, fieldVal reflect.Value) json.Unmarshaler {
	if u, ok := dest.(json.Unmarshaler); ok {
		return u
	}
	if fieldVal.Kind() == reflect.Ptr {
		if fieldVal.IsNil() {
			fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
		}
		if u, ok := fieldVal.Interface().(json.Unmarshaler); ok {
			return u
		}
	}
	if fieldVal.CanAddr() {
		if u, ok := fieldVal.Addr().Interface().(json.Unmarshaler); ok {
			return u
		}
	}
	return nil
}

func configJSONPayload(src interface{}) ([]byte, bool) {
	if src == nil {
		return nil, false
	}
	switch v := src.(type) {
	case string:
		return []byte(v), true
	case *string:
		if v == nil {
			return nil, false
		}
		return []byte(*v), true
	case []byte:
		return v, true
	case *[]byte:
		if v == nil {
			return nil, false
		}
		return *v, true
	}
	sv := reflect.ValueOf(src)
	if sv.Kind() == reflect.Ptr {
		if sv.IsNil() {
			return nil, false
		}
		return configJSONPayload(sv.Elem().Interface())
	}
	if sv.Kind() == reflect.String {
		return []byte(sv.String()), true
	}
	return nil, false
}

func configSetAssignable(field, src reflect.Value) bool {
	if src.Kind() == reflect.Ptr {
		if src.IsNil() {
			return false
		}
		src = src.Elem()
	}
	if !src.IsValid() {
		return false
	}
	if configTryCoerce(field, src) {
		return true
	}
	if src.Type().AssignableTo(field.Type()) {
		field.Set(src)
		return true
	}
	// 勿对 string 使用 reflect.Convert（如 int→string 会得到乱码）
	if field.Kind() != reflect.String && src.Type().ConvertibleTo(field.Type()) {
		field.Set(src.Convert(field.Type()))
		return true
	}
	return false
}

// configTryCoerce 处理 yaml/env 常见类型与注入字段不一致（如 app.port 为 int 8080，字段为 string / *string）。
func configTryCoerce(field, src reflect.Value) bool {
	if !src.IsValid() {
		return false
	}
	if field.Kind() == reflect.String {
		return configScalarToString(field, src)
	}
	if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.String {
		s := configFormatScalarString(src)
		if s == "" && src.Kind() != reflect.String && src.Kind() != reflect.Bool {
			return false
		}
		str := s
		field.Set(reflect.ValueOf(&str))
		return true
	}
	return false
}

func configScalarToString(field, src reflect.Value) bool {
	s := configFormatScalarString(src)
	if s == "" && src.Kind() != reflect.String && src.Kind() != reflect.Bool {
		return false
	}
	field.SetString(s)
	return true
}

func configFormatScalarString(src reflect.Value) string {
	switch src.Kind() {
	case reflect.String:
		return src.String()
	case reflect.Bool:
		return strconv.FormatBool(src.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(src.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(src.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(src.Float(), 'f', -1, 64)
	default:
		return ""
	}
}

func (c *ConfigProvider) panicInjectAssign(configKey, detail string) {
	keyPart := "空键"
	if configKey != "" {
		keyPart = "键 " + strconv.Quote(configKey)
	}
	panic("注入 " + keyPart + " 失败：无法写入字段（" + detail + "）")
}
