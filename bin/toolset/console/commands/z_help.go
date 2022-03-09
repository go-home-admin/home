package commands

var rootPath string

func SetRootPath(root string) {
	rootPath = root
}

// 获取项目跟目录
func getRootPath() string {
	return rootPath
}
