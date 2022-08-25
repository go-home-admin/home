package app

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
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
func Config[T int | string | bool | *services.Config](key string, def T) T {
	val := app.GetBean("config").(app.Bean).GetBean(key)

	if val == nil {
		return def
	}

	return *val.(*T)
}
