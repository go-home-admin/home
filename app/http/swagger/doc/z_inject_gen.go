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

		var temp interface{} = ControllerSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return ControllerSingle
}
