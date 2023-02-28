package filesystem

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app"
	"os"
	"strings"
)

// Local @Bean
type Local struct {
	root string
	url  string
}

func (l *Local) Init() {
	l.root = app.Config("filesystem.local.root", "/storage/")
	l.url = app.Config("filesystem.local.url", "http://127.0.0.1/web/")
}

func (l *Local) FormFile(c *gin.Context, up, to string) (string, error) {
	// 获取上传的文件
	file, err := c.FormFile(up)
	if err != nil {
		return "", err
	}
	// 创建目标文件夹，如果不存在
	dst := l.root + to
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			return "", err
		}
	}
	// 拼接目标文件路径
	dst = dst + "/" + file.Filename
	dst = strings.ReplaceAll(l.url+dst, "//", "")
	// 保存文件到目标路径
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}
	dst = strings.Trim(dst, "/")
	return l.url + dst, nil
}
