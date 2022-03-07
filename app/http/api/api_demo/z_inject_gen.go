// gen for home toolset
package api_demo

var _ControllerSingle *Controller

func GetAllProvider() []interface{} {
	return []interface{}{
		NewController(),
	}
}

func NewController() *Controller {
	if _ControllerSingle == nil {
		Controller := &Controller{}
		_ControllerSingle = Controller
	}
	return _ControllerSingle
}
