// gen for home toolset
package doc

import ()

var _ControllerSingle *Controller

func GetAllProvider() []interface{} {
	return []interface{}{
		NewController(),
	}
}

func NewController() *Controller {
	if _ControllerSingle == nil {
		Controller := &Controller{}
		app.AfterProvider(Controller, "")
		_ControllerSingle = Controller
	}
	return _ControllerSingle
}
