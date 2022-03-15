package app

import "github.com/go-home-admin/home/bootstrap/constraint"

var beansAlias = map[string]interface{}{}

// AfterProvider 在 Initialize 函数后执行
func AfterProvider(bean interface{}, alias string) {
	if alias != "" {
		if _, ok := beansAlias[alias]; !ok {
			beansAlias[alias] = bean
		} else {
			panic("重复的 bean alias " + alias)
		}
	}

	construct, ok := bean.(constraint.Construct)
	if ok {
		construct.Init()
	}
}

type Bean interface {
	GetBean(alias string) interface{}
}

func GetBean(alias string) interface{} {
	return beansAlias[alias]
}
