package commands

import (
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"os"
	"path/filepath"
	"sort"
)

// @Bean
type BeanCommand struct{}

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
			for tName, goType := range fileParser.Types {
				for attName, attr := range goType.Attrs {
					if attr.HasTag("inject") {
						for _, impStr := range fileParser.Imports {
							bc.imports[impStr] = impStr
						}

						_ = attName
						_ = tName
					}
				}
			}
		}

		genBean(dir, bc)
	}
}

type beanCache struct {
	imports map[string]string
	structn map[string]string
}

func newBeanCache() beanCache {
	return beanCache{
		imports: make(map[string]string),
		structn: make(map[string]string),
	}
}

func genBean(dir string, bc beanCache) {

}

func sortMap(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
