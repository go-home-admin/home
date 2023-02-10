package app

import (
	"bytes"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"runtime"
	"strconv"
)

var isDebug = 0

// IsDebug 是否开启调试模式
func IsDebug() bool {
	switch isDebug {
	case 1:
		return false
	case 2:
		return true
	default:
		config := app.GetBean("config").(app.Bean).GetBean("app").(*services.Config)
		d := config.GetBool("debug", false)
		if d {
			isDebug = 2
		} else {
			isDebug = 1
		}
		return d
	}
}

// Name 应用名称
func Name() string {
	return Config("app.name", "go-admin")
}

// Env 获取环境
func Env() string {
	return Config("app.env", "local")
}

// Key 如果敏感信息需要加密, 可以使用这个函数获取盐值
func Key() string {
	return Config("app.key", "go-admin-key")
}

// Config 获取config，格式必须是group.key，第二个可选参数为默认值
func Config[T int | string | bool | []int | []string | interface{} | []interface{}](key string, def T) T {
	val := app.GetBean("config").(app.Bean).GetBean(key)

	switch val.(type) {
	case *interface{}:
		if *val.(*interface{}) == nil {
			return def
		}
	}

	return *val.(*T)
}

// GetGoId
// 获取跟踪ID, 严禁非开发模式使用
// github.com/bigwhite/experiments/blob/master/trace-function-call-chain/trace3/trace.go
func GetGoId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// StringToHump 蛇形转驼峰
func StringToHump(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && (d == '_' || d == '-') && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
