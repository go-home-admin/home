package utils

import (
	"io/ioutil"
	"os"
)

// GetFiles 获取目录下所有文件
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

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
