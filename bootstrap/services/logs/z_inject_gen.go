// gen for home toolset
package logs

import (
	app "github.com/go-home-admin/home/bootstrap/services/app"
)

var _GinLogrusSingle *GinLogrus

func GetAllProvider() []interface{} {
	return []interface{}{
		NewGinLogrus(),
	}
}

func NewGinLogrus() *GinLogrus {
	if _GinLogrusSingle == nil {
		GinLogrus := &GinLogrus{}
		app.AfterProvider(GinLogrus, "")
		_GinLogrusSingle = GinLogrus
	}
	return _GinLogrusSingle
}
