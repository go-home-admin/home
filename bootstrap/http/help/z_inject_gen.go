// gen for home toolset
package help

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _RouteHelpSingle *RouteHelp

func GetAllProvider() []interface{} {
	return []interface{}{
		NewRouteHelp(),
	}
}

func NewRouteHelp() *RouteHelp {
	if _RouteHelpSingle == nil {
		_RouteHelpSingle = &RouteHelp{}
		app.AfterProvider(_RouteHelpSingle, "")
	}
	return _RouteHelpSingle
}
