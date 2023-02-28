package filesystem

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/app"
)

type FileDisk interface {
	FormFile(c *gin.Context, up, to string) string
}

func Disk(t string) FileDisk {
	return app.GetBean("t").(FileDisk)
}
