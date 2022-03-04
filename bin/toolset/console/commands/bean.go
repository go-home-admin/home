package commands

import (
	"fmt"
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// @Bean
type BeanCommand struct {
	r RouteCommand `inject:""`
}

func (BeanCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:bean",
		Description: "生成依赖注入的声明源代码文件",
		Input: command.Argument{
			Option: []command.ArgParam{
				{
					Name:        "path",
					Description: "需要处理的目录",
				},
			},
		},
	}
}

func (BeanCommand) Execute(input command.Input) {
	path := input.GetOption("path")
	if len(path) == 0 {
		path, _ = os.Getwd()
	} else {
		path, _ = filepath.Abs(path)
	}

	for dir, fileParsers := range parser.NewGoParserForDir(path) {
		bc := newBeanCache()
		for _, fileParser := range fileParsers {
			bc.name = fileParser.PackageName
			for tName, goType := range fileParser.Types {
				for _, attr := range goType.Attrs {
					if attr.HasTag("inject") {
						for _, impStr := range fileParser.Imports {
							bc.imports[impStr] = impStr
						}

						bc.structn[tName] = goType
						break
					}
				}
			}
		}

		genBean(dir, bc)
	}
}

type beanCache struct {
	name    string
	imports map[string]string
	structn map[string]parser.GoType
}

func newBeanCache() beanCache {
	return beanCache{
		imports: make(map[string]string),
		structn: make(map[string]parser.GoType),
	}
}

func genBean(dir string, bc beanCache) {
	if len(bc.structn) == 0 {
		return
	}
	context := make([]string, 0)
	str := "// gen for home toolset\n"

	context = append(context, "package "+bc.name)

	aliasMapImport := genImportAlias(bc.imports)
	context = append(context, "\nimport (\n"+getImportStr(aliasMapImport)+"\n)\n")

	for _, s2 := range context {
		str = str + "\n" + s2
	}
	fmt.Println(dir + "/z_inject_gen.test")
	err := os.WriteFile(dir+"/z_inject_gen.test", []byte(str), 0766)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

// 生成 alias => import
func genImportAlias(m map[string]string) map[string]string {
	aliasMapImport := make(map[string]string)
	for _, imp := range m {
		temp := strings.Split(imp, "/")
		key := temp[len(temp)-1]

		if _, ok := aliasMapImport[key]; ok {
			for i := 1; i < 1000; i++ {
				newKey := key + strconv.Itoa(i)
				if _, ok2 := aliasMapImport[newKey]; !ok2 {
					key = newKey
				}
			}
		}
		aliasMapImport[key] = imp
	}

	return aliasMapImport
}

func getImportStr(m map[string]string) string {
	sk := sortMap(m)
	got := ""
	for _, k := range sk {
		got = k + " \"" + m[k] + "\""
	}

	return got
}

func sortMap(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
