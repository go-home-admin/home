package utils

import (
	"io/ioutil"
	"os"
	"strings"
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

// SetEnv 对字符串内容进行替换环境变量
func SetEnv(fileContext []byte) []byte {
	str := string(fileContext)
	arr := strings.Split(str, "\n")

	for _, s := range arr {
		if strings.Index(s, " env(\"") != -1 {
			arr2 := strings.Split(s, ": ")
			if len(arr2) != 2 {
				continue
			}
			nS := arr2[1]
			st, et := GetBrackets(nS, '"', '"')
			key := nS[st : et+1]
			nS = nS[et+1:]
			st, et = GetBrackets(nS, '"', '"')
			val := nS[st : et+1]
			key = strings.Trim(key, "\"")
			val = strings.Trim(val, "\"")

			envVal := os.Getenv(key)
			if envVal != "" {
				val = envVal
			}

			str = strings.Replace(str, s, arr2[0]+": "+val, 1)
		}
	}

	return []byte(str)
}

func GetBrackets(str string, start, end int32) (int, int) {
	var startInt, endInt int

	bCount := 0
	for i, w := range str {
		if bCount == 0 {
			if w == start {
				startInt = i
				bCount++
			}
		} else {
			switch w {
			case end:
				bCount--
				if bCount <= 0 {
					endInt = i
					return startInt, endInt
				}
			case start:
				bCount++
			}
		}
	}

	return startInt, endInt
}
