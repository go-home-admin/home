package providers

import (
	"flag"
	"gopkg.in/ini.v1"
)

// Ini 的文件加载, 其他服务不能直接使用ini, ini只能由config服务使用, 作为预留可替换
// go get gopkg.in/ini.v1
// @Bean
type Ini struct {
	path string
	file *ini.File
}

func NewIniProvider() *Ini {
	Ini := &Ini{}
	Ini.Init()
	return Ini
}

func (i *Ini) Init() {
	flag.StringVar(&i.path, "path", "config.ini", "加载的配置文件")

	//解析命令行参数
	flag.Parse()

	var err error
	i.file, err = ini.Load(i.path)
	if err != nil {
		panic("无法加载基础配置, path=" + i.path)
	}
}

func (i *Ini) Session(name string) *ini.Section {
	return i.file.Section(name)
}
