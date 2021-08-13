// 代码由home-admin生成, 不需要编辑它

package test

var ControllerSingle *Controller

func NewControllerProvider() *Controller {
	Controller := &Controller{}
	return Controller
}

func InitializeNewControllerProvider() *Controller {
	if ControllerSingle == nil {
		ControllerSingle = NewControllerProvider()
	}

	return ControllerSingle
}
