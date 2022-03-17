package app

import "github.com/go-home-admin/home/bootstrap/constraint"

var beansAlias = map[string]interface{}{}
var beansBoot = make([]constraint.AfterRegistration, 0)

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

	boot, ok := bean.(constraint.AfterRegistration)
	if ok {
		beansBoot = append(beansBoot, boot)
	}
}

type Bean interface {
	GetBean(alias string) interface{}
}

func GetBean(alias string) interface{} {
	return beansAlias[alias]
}

func RunBoot() {
	for _, b := range beansBoot {
		b.Boot()
	}
}
