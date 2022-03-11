package commands

import (
	"bytes"
	"fmt"
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// @Bean
type ProtocCommand struct{}

func (ProtocCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:protoc",
		Description: "组装和执行protoc命令",
		Input: command.Argument{
			Option: []command.ArgParam{
				{
					Name:        "proto",
					Description: "proto文件存放的目录",
					Default:     "@root/protobuf",
				},
				{
					Name:        "proto_path",
					Description: "protoc后面拼接的proto_path, 可以传入多个",
					Default:     "@root/protobuf/common/http",
				},
				{
					Name:        "go_out",
					Description: "生成文件到指定目录",
					Default:     "@root/generate/proto",
				},
			},
		},
	}
}

func (ProtocCommand) Execute(input command.Input) {
	root := getRootPath()
	_, err := exec.LookPath("protoc")
	if err != nil {
		log.Printf("'protoc' 未安装; brew install protobuf")
		return
	}
	out := input.GetOption("go_out")
	out = strings.Replace(out, "@root", root, 1)
	outTemp, _ := filepath.Abs(out + "/../temp")
	_ = os.MkdirAll(outTemp, 0766)

	path := input.GetOption("proto")
	path = strings.Replace(path, "@root", root, 1)

	pps := make([]string, 0)
	for _, s := range input.GetOptions("proto_path") {
		s = strings.Replace(s, "@root", root, 1)
		pps = append(pps, "--proto_path="+s)
	}
	// path/*.proto 不是protoc命令提供的, 如果这里执行需要每一个文件一个命令
	for _, dir := range parser.GetChildrenDir(path) {
		for _, info := range dir.GetFiles(".proto") {
			cods := []string{"--proto_path=" + dir.Path}
			cods = append(cods, pps...)
			cods = append(cods, "--go_out="+outTemp)
			cods = append(cods, info.Path)

			Cmd("protoc", cods)
		}
	}

	// 生成后, 从temp目录拷贝到out
	_ = os.RemoveAll(out)
	rootAlias := strings.Replace(out, root+"/", "", 1)
	module := getModModule()

	for _, dir := range parser.GetChildrenDir(outTemp) {
		dir2 := strings.Replace(dir.Path, outTemp+"/", "", 1)
		dir3 := strings.Replace(dir2, module+"/", "", 1)
		if dir2 == dir3 {
			continue
		}

		if dir3 == rootAlias {
			_ = os.Rename(dir.Path, out)
			_ = os.RemoveAll(outTemp)
			break
		}
	}
}

func Cmd(commandName string, params []string) {
	// 打印真实命令
	PrintCmd(commandName, params)

	cmd := exec.Command(commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func PrintCmd(commandName string, params []string) {
	str := "\n" + commandName + " "
	for _, param := range params {
		str += param + " "
	}
	fmt.Print(str + "\n")
}
