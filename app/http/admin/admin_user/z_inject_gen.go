// gen for home toolset
package admin_user

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
