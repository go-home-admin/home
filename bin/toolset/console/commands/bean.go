package commands

import (
	"github.com/ctfang/command"
	"github.com/go-home-admin/toolset/parser"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// @Bean
type BeanCommand struct{}

func (BeanCommand) Configure() command.Configure {
	return command.Configure{
		Name:        "make:bean",
		Description: "生成依赖注入的声明源代码文件",
		Input: command.Argument{
			Has: []command.ArgParam{
				{
					Name:        "-f",
					Description: "强制更新",
				},
			},
		},
	}
}

func (BeanCommand) Execute(input command.Input) {
	path := getRootPath()

	//if input.GetHas("-f") == true {
	//	for dir, _ := range parser.NewGoParserForDir(path) {
	//		if
	//	}
	//}

	skip := map[string]bool{
		"/Users/lv/Desktop/github.com/go-hom-admin/home/bin": true,
	}

	for dir, fileParsers := range parser.NewGoParserForDir(path) {
		if _, ok := skip[dir]; ok {
			break
		}

		bc := newBeanCache()
		for _, fileParser := range fileParsers {
			bc.name = fileParser.PackageName
			for tName, goType := range fileParser.Types {
				for _, attr := range goType.Attrs {
					if attr.HasTag("inject") {
						for _, impStr := range fileParser.Imports {
							bc.imports[impStr] = impStr
						}

						break
					}
				}

				if goType.Doc.HasAnnotation("@Bean") {
					bc.structList[tName] = goType
				}
			}
		}

		genBean(dir, bc)
	}
}

type beanCache struct {
	name       string
	imports    map[string]string
	structList map[string]parser.GoType
}

func newBeanCache() beanCache {
	return beanCache{
		imports:    make(map[string]string),
		structList: make(map[string]parser.GoType),
	}
}

func genBean(dir string, bc beanCache) {
	if len(bc.structList) == 0 {
		return
	}
	context := make([]string, 0)
	context = append(context, "package "+bc.name)

	// import
	importAlias := genImportAlias(bc.imports)
	if len(importAlias) != 0 {
		context = append(context, "\nimport ("+getImportStr(bc, importAlias)+"\n)")
	}

	// Single
	context = append(context, genSingle(bc))
	// Provider
	context = append(context, genProvider(bc, importAlias))
	str := "// gen for home toolset"
	for _, s2 := range context {
		str = str + "\n" + s2
	}

	err := os.WriteFile(dir+"/z_inject_gen.go", []byte(str), 0766)
	if err != nil {
		log.Fatal(err)
	}
}

func genSingle(bc beanCache) string {
	str := ""
	allProviderStr := "\n\treturn []interface{}{"
	for s, goType := range bc.structList {
		if goType.Doc.HasAnnotation("@Bean") {
			str = str + "\nvar " + genSingleName(s) + " *" + s
			allProviderStr += "\n\t\t" + genInitializeNewStr(s) + "(),"
		}
	}
	// 返回全部的提供商
	str += "\n\nfunc GetAllProvider() []interface{} {" + allProviderStr + "\n\t}\n}"
	return str
}

func genSingleName(s string) string {
	return "_" + s + "Single"
}

func genProvider(bc beanCache, m map[string]string) string {
	str := ""
	for s, goType := range bc.structList {
		if goType.Doc.HasAnnotation("@Bean") {
			str = str + "\nfunc " + genInitializeNewStr(s) + "() *" + s + " {" +
				"\n\tif " + genSingleName(s) + " == nil {" + // if _provider == nil {
				"\n\t\t" + s + " := " + "&" + s + "{}" // provider := provider{}

			for attrName, attr := range goType.Attrs {
				pointer := ""
				if !attr.IsPointer() {
					pointer = "*"
				}

				for tagName, _ := range attr.Tag {
					if tagName == "inject" {
						str = str + "\n\t\t" +
							s + "." + attrName + " = " + pointer + getInitializeNewFunName(attr, m)
					}
				}
			}
			str = str +
				"\n\t\t" + genSingleName(s) + " = " + s +
				"\n\t}" +
				"\n\treturn " + genSingleName(s) +
				"\n}"
		}
	}

	return str
}

func getInitializeNewFunName(k parser.GoTypeAttr, m map[string]string) string {
	alias := ""
	name := k.TypeName
	if !k.InPackage {
		a := m[k.TypeImport]
		alias = a + "."
		arr := strings.Split(k.TypeName, ".")
		name = arr[len(arr)-1]
	}

	return alias + genInitializeNewStr(name) + "()"
}

// 控制对完函数名称
func genInitializeNewStr(name string) string {
	return "New" + name
}

// 生成 import => alias
func genImportAlias(m map[string]string) map[string]string {
	aliasMapImport := make(map[string]string)
	importMapAlias := make(map[string]string)
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
	for s, s2 := range aliasMapImport {
		importMapAlias[s2] = s
	}

	return importMapAlias
}

// m = import => alias
func getImportStr(bc beanCache, m map[string]string) string {
	has := map[string]bool{}
	for _, goType := range bc.structList {
		if goType.Doc.HasAnnotation("@Bean") {
			for _, attr := range goType.Attrs {
				if !attr.InPackage {
					has[attr.TypeImport] = true
				}
			}

		}
	}
	// 删除未使用的import
	for s, _ := range m {
		if _, ok := has[s]; !ok {
			delete(m, s)
		}
	}

	sk := sortMap(m)
	got := ""
	for _, k := range sk {
		got += "\n\t" + m[k] + " \"" + k + "\""
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
