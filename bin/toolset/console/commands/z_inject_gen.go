// gen for home toolset
package commands

var _BeanCommandSingle *BeanCommand
var _RouteCommandSingle *RouteCommand

func GetAllProvider() []interface{} {
	return []interface{}{
		NewBeanCommand(),
		NewRouteCommand(),
	}
}

func NewBeanCommand() *BeanCommand {
	if _BeanCommandSingle == nil {
		BeanCommand := &BeanCommand{}
		_BeanCommandSingle = BeanCommand
	}
	return _BeanCommandSingle
}
func NewRouteCommand() *RouteCommand {
	if _RouteCommandSingle == nil {
		RouteCommand := &RouteCommand{}
		_RouteCommandSingle = RouteCommand
	}
	return _RouteCommandSingle
}
