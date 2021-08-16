package constraint

type ServiceProvider interface {
	register()
	boot()
}

// 具备结构体Bean
type Construct interface {
	// 初始化
	Init()
}
