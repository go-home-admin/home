package providers

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
	"strings"
)

// GetBean 只能返回指针的值
func GetBean(alias string) interface{} {
	if !app.HasBean(alias) {
		NewFrameworkProvider()

		if !app.HasBean(alias) {
			arr := strings.Split(alias, ", ")
			// 如果是系统级服务, 并且默认不启动的
			// 继续注册自动
			switch arr[0] {
			case "mysql":
				NewMysqlProvider()
			case "redis":
				NewRedisProvider()
			case "config":
				NewConfigProvider()
			}
		}
	}

	return app.GetBean(alias)
}

func AfterProvider(bean interface{}, alias string) {
	app.AfterProvider(bean, alias)
}

type Bean interface {
	// GetBean 只能返回指针的值
	GetBean(alias string) interface{}
}
