package commands

import (
	"bytes"
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"log"
	"os"
	"os/exec"
	"strings"
)

// @Bean
type RouteCommand struct{}

func (RouteCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:route",
		Description: "根据protoc文件定义, 生成路由信息和控制器文件",
		Input: command.Argument{
			Option: []command.ArgParam{
				{
					Name:        "path",
					Description: "只解析指定目录",
					Default:     "@root/protobuf",
				},
				{
					Name:        "out",
					Description: "生成文件到指定目录",
					Default:     "@root/routes",
				},
			},
		},
	}
}

func (RouteCommand) Execute(input command.Input) {
	root := getRootPath()
	module := getModModule()
	out := input.GetOption("out")
	out = strings.Replace(out, "@root", root, 1)
	path := input.GetOption("path")
	path = strings.Replace(path, "@root", root, 1)

	agl := map[string]*ApiGroups{}

	for _, parsers := range parser.NewProtocParserForDir(path) {
		for _, fileParser := range parsers {
			for _, service := range fileParser.Services {
				group := ""

				for _, option := range service.Opt {
					if option.Key == "http.RouteGroup" {
						group = option.Val
						if _, ok := agl[group]; !ok {
							agl[group] = &ApiGroups{
								name: group,
								imports: map[string]string{
									"home_api_1": "github.com/go-home-admin/home/bootstrap/http/api",
									"home_gin_1": "github.com/gin-gonic/gin",
								},
								controllers: make([]Controller, 0),
								servers:     make([]parser.Service, 0),
							}
						}
						break
					}
				}

				if group != "" {
					g := agl[group]
					imports := module + "/app/http/" + fileParser.PackageName + "/" + parser.StringToSnake(service.Name)
					g.imports[imports] = imports

					g.controllers = append(g.controllers, Controller{
						name:  service.Name,
						alias: imports,
					})

					g.servers = append(g.servers, service)
				}
			}
		}
	}

	for _, g := range agl {
		genRoute(g, out)
	}
	cmd := exec.Command("go", []string{"fmt", out}...)
	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = os.Stderr
	cmd.Dir = out
	_ = cmd.Run()
}

func genRoute(g *ApiGroups, out string) {
	context := make([]string, 0)
	context = append(context, "package routes")

	// import
	importAlias := genImportAlias(g.imports)
	if len(importAlias) != 0 {
		context = append(context, "\nimport ("+getImportStrForMap(importAlias)+"\n)")
	}
	// Routes struct
	context = append(context, genRoutesStruct(g, importAlias))
	// Routes struct func GetRoutes
	context = append(context, genRoutesFunc(g, importAlias))

	str := "// gen for home toolset"
	for _, s2 := range context {
		str = str + "\n" + s2
	}
	err := os.WriteFile(out+"/"+parser.StringToSnake(g.name)+"_gen.go", []byte(str), 0766)
	if err != nil {
		log.Fatal(err)
	}
}

func genRoutesFunc(g *ApiGroups, m map[string]string) string {
	homeGin := m["github.com/gin-gonic/gin"]
	homeApi := m["github.com/go-home-admin/home/bootstrap/http/api"]

	str := "\nfunc (c *" + parser.StringToHump(g.name) + "Routes) GetRoutes() map[*" + homeApi + ".Config]func(c *" + homeGin + ".Context) {" +
		"\n\treturn map[*" + homeApi + ".Config]func(c *" + homeGin + ".Context){"

	for _, server := range g.servers {
		for rName, rpc := range server.Rpc {
			for _, option := range rpc.Opt {
				if strings.Index(option.Key, "http.") == 0 {
					i := strings.Index(option.Key, ".")
					method := option.Key[i+1:]
					str += "\n\t\t" + homeApi + "." + method + "(\"" + option.Val + "\"):" +
						"c." + parser.StringToSnake(server.Name) + ".GinHandle" + parser.StringToHump(rName) + ","
				}
			}
		}
	}

	return str + "\n\t}\n}"
}

func genRoutesStruct(g *ApiGroups, m map[string]string) string {
	str := "\n// @Bean" +
		"\ntype " + parser.StringToHump(g.name) + "Routes struct {\n"
	for _, controller := range g.controllers {
		alias := m[controller.alias]
		str += "\t" + parser.StringToSnake(controller.name) + " *" + alias + ".Controller" + " `inject:\"\"`\n"
	}

	return str + "}\n"
}

type ApiGroups struct {
	name        string
	imports     map[string]string
	controllers []Controller
	servers     []parser.Service
}

type Controller struct {
	name  string
	alias string
	ty    string // *alias.Controller
}

func getImportStrForMap(m map[string]string) string {
	sk := sortMap(m)
	got := ""
	for _, k := range sk {
		got += "\n\t" + m[k] + " \"" + k + "\""
	}

	return got
}
