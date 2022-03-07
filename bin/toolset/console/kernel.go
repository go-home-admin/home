package console

import (
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/console/commands"
)

// @Bean
type Kernel struct{}

func (k *Kernel) Run() {
	app := command.New()
	for _, provider := range commands.GetAllProvider() {
		if v, ok := provider.(command.CommandInterface); ok {
			app.AddCommand(v)
		}
	}
	app.Run()
}

func (k *Kernel) Exit() {

}
