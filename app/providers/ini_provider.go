package providers

import (
	"flag"
	"github.com/go-home-admin/home/bootstrap/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Ini 的文件加载, 其他服务不能直接使用ini, ini只能由config服务使用, 作为预留可替换
// go get gopkg.in/ini.v1
// @Bean
type Ini struct {
	path string
	file *ini.File
}

func (i *Ini) Init() {
	flag.StringVar(&i.path, "path", "", "加载的配置目录, 默认在./config")

	if !flag.Parsed() {
		flag.Parse()
	}

	if i.path == "" {
		pwd, _ := os.Getwd()
		for j := 0; j < 30; j++ {
			dir, _ := filepath.Abs(pwd)
			if utils.IsExist(dir) && utils.IsExist(dir+"/go.mod") && utils.IsExist(dir+"/config") {
				i.path = dir + "/config"
				break
			}
			pwd = pwd + "/../"
		}
		logrus.Info("未指定配置, 系统自动加载配置目录: " + i.path)
	}

	cfl, err := utils.GetFiles(i.path)
	if err != nil {
		logrus.Error(err)
	}

	i.file = ini.Empty()
	for _, file := range cfl {
		if path.Ext(file) == ".ini" {
			s, _ := ioutil.ReadFile(file)
			if i.file.Append(file, s) != nil {
				logrus.Error(err)
			}
		}
	}
}

func (i *Ini) Session(name string) *ini.Section {
	return i.file.Section(name)
}
