package commands

import (
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"log"
)

// @Bean
type RouteCommand struct{}

func (RouteCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:route",
		Description: "根据protoc文件定义, 生成路由信息和控制器文件",
		Input:       command.Argument{},
	}
}

func (RouteCommand) Execute(input command.Input) {
	path := getRootPath()

	for s, parsers := range parser.NewProtocParserForDir(path) {
		println(s, parsers)
	}

	log.Println("echo command")
}
