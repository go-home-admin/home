package utils

import (
	"io/ioutil"
	"os"
)

// 获取目录下所有文件
func GetFiles(dirPath string) ([]string, error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var files []string
	pthSep := string(os.PathSeparator)
	for _, info := range dir {
		if !info.IsDir() {
			files = append(files, dirPath+pthSep+info.Name())
		}
	}
	return files, nil
}
