package commands

import (
	"github.com/ctfang/command"
	"log"
)

// @Bean
type RouteCommand struct{}

func (RouteCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:route",
		Description: "",
		Input:       command.Argument{},
	}
}

func (RouteCommand) Execute(input command.Input) {
	log.Println("echo command")
}
