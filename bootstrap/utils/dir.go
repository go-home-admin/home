package utils

import (
	"os"
	"strings"
)

// GetFiles 获取目录下所有文件
func GetFiles(dirPath string) ([]string, error) {
	dir, err := os.ReadDir(dirPath)
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
			key := nS[st+1 : et]
			nS = nS[et+1:]

			// 尝试获取默认值
			val := ""
			valIsStr := false
			if len(nS) > 2 && nS[0:1] == "," {
				nS = strings.TrimSpace(nS[1:])
				nS = strings.Trim(nS, ")") // 得到 "val" or val
				if nS[0:1] == "\"" {
					// 使用双引号括起来的就是字符串
					valIsStr = true
					st, et = GetBrackets(nS, '"', '"')
					val = nS[st+1 : et]
				} else {
					val = nS
				}
			}

			envVal, has := os.LookupEnv(key)
			if has {
				val = envVal
			}

			if !valIsStr {
				// 默认情况, 把值粘贴到yaml, 类型自动识别
				str = strings.Replace(str, s, arr2[0]+": "+val, 1)
			} else {
				// 如果有默认值, 根据默认值识别类型
				str = strings.Replace(str, s, arr2[0]+": \""+val+"\"", 1)
			}
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
