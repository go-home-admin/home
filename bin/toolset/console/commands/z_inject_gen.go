// gen for home toolset
package commands

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
)

var _RouteCommandSingle *RouteCommand
var _BeanCommandSingle *BeanCommand
var _ProtocCommandSingle *ProtocCommand

func GetAllProvider() []interface{} {
	return []interface{}{
		NewRouteCommand(),
		NewBeanCommand(),
		NewProtocCommand(),
	}
}

func NewBeanCommand() *BeanCommand {
	if _BeanCommandSingle == nil {
		BeanCommand := &BeanCommand{}
		app.AfterProvider(BeanCommand, "")
		_BeanCommandSingle = BeanCommand
	}
	return _BeanCommandSingle
}
func NewProtocCommand() *ProtocCommand {
	if _ProtocCommandSingle == nil {
		ProtocCommand := &ProtocCommand{}
		app.AfterProvider(ProtocCommand, "")
		_ProtocCommandSingle = ProtocCommand
	}
	return _ProtocCommandSingle
}
func NewRouteCommand() *RouteCommand {
	if _RouteCommandSingle == nil {
		RouteCommand := &RouteCommand{}
		app.AfterProvider(RouteCommand, "")
		_RouteCommandSingle = RouteCommand
	}
	return _RouteCommandSingle
}
