// gen for home toolset
package commands

var _BeanCommandSingle *BeanCommand
var _ProtocCommandSingle *ProtocCommand
var _RouteCommandSingle *RouteCommand

func GetAllProvider() []interface{} {
	return []interface{}{
		NewBeanCommand(),
		NewProtocCommand(),
		NewRouteCommand(),
	}
}

func NewProtocCommand() *ProtocCommand {
	if _ProtocCommandSingle == nil {
		ProtocCommand := &ProtocCommand{}
		_ProtocCommandSingle = ProtocCommand
	}
	return _ProtocCommandSingle
}
func NewRouteCommand() *RouteCommand {
	if _RouteCommandSingle == nil {
		RouteCommand := &RouteCommand{}
		_RouteCommandSingle = RouteCommand
	}
	return _RouteCommandSingle
}
func NewBeanCommand() *BeanCommand {
	if _BeanCommandSingle == nil {
		BeanCommand := &BeanCommand{}
		_BeanCommandSingle = BeanCommand
	}
	return _BeanCommandSingle
}
