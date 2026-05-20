package providers

import "strconv"

// PanicFlatInjectGetBeanTypeMismatch 供 toolset make:bean 生成的扁平 inject 使用：
// Bean 分支上 GetBean 返回值与目标字段类型断言失败时调用，避免每条注入重复长字符串。
func PanicFlatInjectGetBeanTypeMismatch(configKey string) {
	if configKey == "" {
		panic("注入 空键 失败：GetBean 返回值类型与字段不一致")
	}
	panic("注入 键 " + strconv.Quote(configKey) + " 失败：GetBean 返回值类型与字段不一致")
}

// PanicFlatInjectUnknownProvider 依赖既不是 InjectValue 也不是 Bean 时调用。
func PanicFlatInjectUnknownProvider(configKey string) {
	if configKey == "" {
		panic("注入 空键 失败：依赖既不是 InjectValue 也不是 Bean")
	}
	panic("注入 键 " + strconv.Quote(configKey) + " 失败：依赖既不是 InjectValue 也不是 Bean")
}
