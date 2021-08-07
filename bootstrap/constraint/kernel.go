package constraint

// 每个独立服务的业务配置
type KernelServer interface {
	Init()
	Run()
	Exit()
}
