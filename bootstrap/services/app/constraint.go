package app

import "github.com/go-home-admin/home/bootstrap/constraint"

var beansAlias = map[string]interface{}{}
var beansBoot = make([]constraint.AfterRegistration, 0)
var exitBoot = make([]constraint.Exit, 0)

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

	exit, ok := bean.(constraint.Exit)
	if ok {
		exitBoot = append(exitBoot, exit)
	}
}

type Bean interface {
	GetBean(alias string) interface{}
}

func GetBean(alias string) interface{} {
	return beansAlias[alias]
}

// RunBoot
// 所有的 Init() 执行后, 触发Boot()
// Boot() 、Init() 是倒叙执行, 被依赖的先执行
func RunBoot() {
	for _, b := range beansBoot {
		b.Boot()
	}
}

// RunExit
// 是顺序执行, 发起依赖的先执行(http->config)
func RunExit() {
	for i := len(exitBoot) - 1; i >= 0; i-- {
		exitBoot[i].Exit()
	}
}
