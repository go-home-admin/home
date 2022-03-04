package console

import (
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/console/commands"
)

// @Bean
type Kernel struct {
	bean  *commands.BeanCommand `inject:""`
	route commands.RouteCommand `inject:""`
}

func (k *Kernel) Run() {
	app := command.New()
	app.AddCommand(commands.BeanCommand{})
	app.AddCommand(commands.RouteCommand{})
	app.Run()
}

func (k *Kernel) Exit() {

}
