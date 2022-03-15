package constraint

type ServiceProvider interface {
	register()
	boot()
}

// Construct 具备结构体Bean
type Construct interface {
	// Init 初始化
	Init()
}

// AfterRegistration 注册完成后, 统一处理
type AfterRegistration interface {
	AfterRegistration(beans []interface{})
}
