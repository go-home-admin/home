package app

import (
	"fmt"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"strings"
	"sync"
)

var beansAlias = sync.Map{}
var beansBoot = make([]constraint.AfterRegistration, 0)
var beansRunAfter = make([]constraint.RunAfter, 0)
var exitBoot = make([]constraint.Exit, 0)

// AfterProvider 在 Initialize 函数后执行
func AfterProvider(bean interface{}, alias string) {
	if alias != "" {
		if _, ok := beansAlias.Load(alias); !ok {
			beansAlias.Store(alias, bean)
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

	runAfter, ok := bean.(constraint.RunAfter)
	if ok {
		beansRunAfter = append(beansRunAfter, runAfter)
	}

	exit, ok := bean.(constraint.Exit)
	if ok {
		exitBoot = append(exitBoot, exit)
	}
}

type Bean interface {
	// GetBean 只能返回指针的值
	GetBean(alias string) interface{}
}

type AppendRun interface {
	AppendRun(fun func())
}

// GetBean 只能返回指针的值
func GetBean(alias string) interface{} {
	arr := strings.Split(alias, ", ")
	bean, ok := beansAlias.Load(arr[0])
	if !ok {
		fmt.Printf("无法找到服务 %v\n", alias)
		fmt.Printf("如果您在测试待遇中使用请调用providers.NewApp()\n")
		panic("提供者别名方式的使用需要提前注册, 可以写到app_provider.go文件注册。")
	}
	if len(arr) == 1 {
		return bean
	}

	return bean.(Bean).GetBean(strings.Join(arr[1:], "."))
}

func HasBean(alias string) bool {
	arr := strings.Split(alias, ", ")
	_, ok := beansAlias.Load(arr[0])
	if !ok {
		return false
	}
	return true
}

// RunBoot
// 所有的 Init() 执行后, 触发Boot()
// Boot() 、Init() 是倒叙执行, 被依赖的先执行
// Boot() 如果是嵌套多个Bean, 可能被多次执行
func RunBoot() {
	for _, b := range beansBoot {
		b.Boot()
	}
}

func RunRunAfter() {
	for i := len(beansRunAfter) - 1; i >= 0; i-- {
		beansRunAfter[i].RunAfter()
	}
}

// RunExit
// 是顺序执行, 发起依赖的先执行(http->config)
func RunExit() {
	for i := len(exitBoot) - 1; i >= 0; i-- {
		exitBoot[i].Exit()
	}
}
