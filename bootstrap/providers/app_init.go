package providers

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
)

// GetBean 只能返回指针的值
func GetBean(alias string) interface{} {
	if !app.HasBean(alias) {
		NewFrameworkProvider()
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
