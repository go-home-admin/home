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
	config := app.GetBean("config").(app.Bean).GetBean("app").(*services.Config)
	return config.GetString("name", "go-admin")
}

// Env 获取环境
func Env() string {
	config := app.GetBean("config").(app.Bean).GetBean("app").(*services.Config)
	return config.GetString("env", "local")
}

// Key 如果敏感信息需要加密, 可以使用这个函数获取盐值
func Key() string {
	config := app.GetBean("config").(app.Bean).GetBean("app").(*services.Config)
	return config.GetString("key", "go-admin-key")
}
