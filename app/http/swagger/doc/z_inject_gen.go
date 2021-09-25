// 代码由home-admin生成, 不需要编辑它

package doc

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var ControllerSingle *Controller

func NewControllerProvider() *Controller {
	Controller := &Controller{}
	return Controller
}

func InitializeNewControllerProvider() *Controller {
	if ControllerSingle == nil {
		ControllerSingle = NewControllerProvider()

		home_constraint.AfterProvider(ControllerSingle)
	}

	return ControllerSingle
}
