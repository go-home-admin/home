// gen for home toolset
package doc

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
