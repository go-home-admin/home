package provoders

import "flag"

// Ini 的文件加载, 其他服务不能直接使用ini, ini只能有config服务使用, 作为预留可替换
// @Bean
type Ini struct {
	path string
}

func NewIniProvider() *Ini {
	Ini := &Ini{}
	return Ini
}

func (i *Ini) Init() {
	flag.StringVar(&i.path, "path", "config.ini", "加载的配置文件")

	//解析命令行参数
	flag.Parse()
}
