package parser

import (
	"fmt"
	"strings"
)

/**
golang parser 非完整token实现
*/
type GoFileParser struct {
	PackageName string
	PackageDoc  string
	Imports     map[string]string
	Types       map[string]GoType
	Funds       map[string]GoFunc
}

func NewGoParserForDir(path string) map[string][]GoFileParser {
	got := make(map[string][]GoFileParser)
	for _, dir := range GetChildrenDir(path) {
		arr := make([]GoFileParser, 0)
		for _, file := range dir.GetFiles(".go") {
			gof, _ := GetFileParser(file.Path)
			arr = append(arr, gof)
		}
		got[dir.Path] = arr
	}

	return got
}

func GetFileParser(path string) (GoFileParser, error) {
	d := GoFileParser{
		PackageName: "",
		PackageDoc:  "",
		Imports:     make(map[string]string),
		Types:       make(map[string]GoType),
		Funds:       make(map[string]GoFunc),
	}

	l := getWordsWitchFile(path)
	lastDoc := ""
	for offset := 0; offset < len(l.list); offset++ {
		work := l.list[offset]
		// 原则上, 每个块级别的作用域必须自己处理完, 返回的偏移必须是下一个块的开始
		switch work.t {
		case wordT_line:
		case wordT_division:
		case wordT_doc:
			lastDoc = work.str
		case wordT_word:
			switch work.str {
			case "package":
				d.PackageDoc = lastDoc
				d.PackageName, offset = handlePackageName(l.list, offset)
				lastDoc = ""
			case "import":
				var imap map[string]string
				imap, offset = handleImports(l.list, offset)
				for k, v := range imap {
					d.Imports[k] = v
				}
				lastDoc = ""
			case "type":
				var imap GoType
				imap, offset = handleTypes(l.list, offset, d)
				imap.Doc = GoDoc(lastDoc)
				d.Types[imap.Name] = imap
				lastDoc = ""
			case "func":
				var gf GoFunc
				gf, offset = handleFunds(l.list, offset)
				d.Funds[gf.Name] = gf
				lastDoc = ""
			case "const":
				_, offset = handleCosts(l.list, offset)
				lastDoc = ""
			case "var":
				_, offset = handleVars(l.list, offset)
				lastDoc = ""
			default:
				fmt.Println("文件块作用域似乎解析有错误", path, work.str, offset)
			}
		}
	}

	return d, nil
}

func handlePackageName(l []*word, offset int) (string, int) {
	name, i := GetFistWordBehindStr(l[offset:], "package")
	return name, offset + i
}

func getImport(sl []string) (string, string) {
	if len(sl) == 2 {
		return sl[0], sl[1][1 : len(sl[1])-1]
	}

	str := sl[0][1 : len(sl[0])-1]
	temp := strings.Split(str, "/")
	key := temp[len(temp)-1]
	return key, str
}

func handleImports(l []*word, offset int) (map[string]string, int) {
	newOffset := offset
	imap := make(map[string]string)
	var key, val string

	ft, fti := GetFistStr(l[offset+1:])
	if ft != "(" {
		arr := make([]string, 0)
		for i, w := range l[offset+fti:] {
			if wordT_line == w.t {
				newOffset = offset + fti + i
				key, val = getImport(arr)
				imap[key] = val
				return imap, newOffset
			}

			if w.t == wordT_word {
				arr = append(arr, w.str)
			}
		}
	} else {
		st, et := GetBrackets(l[offset+1:], "(", ")")
		st = st + offset + 1
		et = et + offset + 1
		newOffset = et

		arr := make([]string, 0)
		for _, w := range l[st:et] {
			if wordT_line == w.t && len(arr) != 0 {
				key, val = getImport(arr)
				imap[key] = val
				arr = make([]string, 0)
			}

			if w.t == wordT_word {
				arr = append(arr, w.str)
			}
		}
	}
	return imap, newOffset
}

type GoType struct {
	Doc   GoDoc
	Name  string
	Attrs map[string]GoTypeAttr
}
type GoTypeAttr struct {
	Name       string
	TypeName   string
	TypeAlias  string
	TypeImport string
	InPackage  bool // 是否本包的引用
	Tag        map[string]string
}

type GoDoc string

// 是否存在某个注解
func (d GoDoc) HasAnnotation(check string) bool {
	ds := string(d)

	return strings.Index(ds, check) != -1
}

// 普通指针
func (receiver GoTypeAttr) IsPointer() bool {
	return receiver.TypeName[0:1] == "*"
}

func (receiver GoTypeAttr) HasTag(name string) bool {
	for s, _ := range receiver.Tag {
		if s == name {
			return true
		}
	}
	return false
}

// 组装成数组, 只限name type other\n结构
func getArrGoWord(l []*word) [][]string {
	got := make([][]string, 0)
	arr := GetArrWord(l)
	for _, i := range arr {
		lis := i[len(i)-1].str
		if lis[0:1] == "`" && len(i) >= 3 {
			ty := ""
			for in := 1; in < len(i)-1; in++ {
				if i[in].t != wordT_doc {
					ty = ty + i[in].str
				}
			}
			got = append(got, []string{i[0].str, ty, lis})
		}
	}

	return got
}

// 把go结构的tag格式化成数组 source = `inject:"" json:"orm"`
func getArrGoTag(source string) [][]string {
	tagStr := source[1 : len(source)-1]
	wl := GetWords(tagStr)
	// 每三个一组
	i := 0
	got := make([][]string, 0)
	arr := make([]string, 0)
	for _, w := range wl {
		if w.t == wordT_word {
			arr = append(arr, w.str)
			i++
			if i >= 2 {
				i = 0
				got = append(got, arr)
				arr = make([]string, 0)
			}
		}
	}

	return got
}
func handleTypes(l []*word, offset int, d GoFileParser) (GoType, int) {
	newOffset := offset
	nl := l[offset:]
	got := GoType{
		Doc:   "",
		Name:  "",
		Attrs: map[string]GoTypeAttr{},
	}
	ok, off := GetLastIsIdentifier(nl, "{")
	if ok {
		// 新结构
		var i int
		got.Name, i = GetFistWordBehindStr(nl, "type")
		nl = nl[i+1:]
		st, et := GetBrackets(nl, "{", "}")
		newOffset = offset + i + et + 1
		nl := nl[st+1 : et]
		arrLn := getArrGoWord(nl)
		for _, wordAttrs := range arrLn {
			// 获取属性信息
			// TODO 当前仅支持有tag的
			if len(wordAttrs) == 3 && strings.Index(wordAttrs[2], "`") == 0 {
				attr := GoTypeAttr{
					Name:     wordAttrs[0],
					TypeName: wordAttrs[1],
					Tag:      map[string]string{},
				}
				getTypeAlias(wordAttrs[1], d, &attr)
				// 解析 go tag
				tagArr := getArrGoTag(wordAttrs[2])

				for _, tagStrArr := range tagArr {
					attr.Tag[tagStrArr[0]] = tagStrArr[1]
				}
				got.Attrs[attr.Name] = attr
			}
		}
	} else {
		// struct 别名
		got.Name, _ = GetFistWordBehindStr(nl, "type")
		newOffset = off + offset
	}

	return got, newOffset
}

// 根据属性声明类型或者类型的引入名称
func getTypeAlias(str string, d GoFileParser, attr *GoTypeAttr) {
	wArr := GetWords(str)
	wf := wArr[0]

	if wf.t == wordT_word || wf.str == "*" {
		if len(wArr) >= 2 {
			attr.TypeAlias, _ = GetFistWord(wArr)
			attr.TypeImport = d.Imports[attr.TypeAlias]
			return
		}
	}
	// 本包
	attr.TypeAlias = d.PackageName
	attr.TypeImport = "" // TODO
	attr.InPackage = true
}

type GoFunc struct {
	Name string
	Stu  string
}

func handleFunds(l []*word, offset int) (GoFunc, int) {
	ft, _ := GetFistStr(l[offset+1:])
	name := ""
	if ft != "(" {
		// 普通函数
		var i int
		name, i = GetFistWordBehindStr(l[offset:], "func")
		offset = offset + i
		_, et := GetBrackets(l[offset:], "(", ")")
		offset = offset + et
	} else {
		// 结构函数
		_, et := GetBrackets(l[offset:], "(", ")")
		offset = offset + et
		name, _ = GetFistWord(l[offset:])
		_, et = GetBrackets(l[offset:], "(", ")")
		offset = offset + et
	}
	// 排除返回值的interface{}
	st, et := GetBrackets(l[offset:], "{", "}")
	interCount := 0
	for _, w := range l[offset : offset+st] {
		if w.str == "interface" {
			interCount++
		}
	}
	if interCount != 0 {
		for i := 0; i <= interCount; i++ {
			_, et := GetBrackets(l[offset:], "{", "}")
			offset = offset + et
		}
	} else {
		offset = offset + et
	}
	return GoFunc{Name: name}, offset
}
func handleCosts(l []*word, offset int) (map[string]string, int) {
	return handleVars(l, offset)
}

func handleVars(l []*word, offset int) (map[string]string, int) {
	ft, _ := GetFistStr(l[offset+1:])
	if ft != "(" {
		return nil, offset + NextLine(l[offset:])
	} else {
		_, et := GetBrackets(l[offset:], "(", ")")
		return nil, offset + et
	}
}
