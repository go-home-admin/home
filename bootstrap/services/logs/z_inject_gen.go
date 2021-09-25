// 代码由home-admin生成, 不需要编辑它

package logs

import (
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var GinLogrusSingle *GinLogrus

func NewGinLogrusProvider() *GinLogrus {
	GinLogrus := &GinLogrus{}
	return GinLogrus
}

func InitializeNewGinLogrusProvider() *GinLogrus {
	if GinLogrusSingle == nil {
		GinLogrusSingle = NewGinLogrusProvider()

		home_constraint.AfterProvider(GinLogrusSingle)
	}

	return GinLogrusSingle
}
