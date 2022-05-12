// gen for home toolset
package logs

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _GinLogrusSingle *GinLogrus

func GetAllProvider() []interface{} {
	return []interface{}{
		NewGinLogrus(),
	}
}

func NewGinLogrus() *GinLogrus {
	if _GinLogrusSingle == nil {
		_GinLogrusSingle = &GinLogrus{}
		providers.AfterProvider(_GinLogrusSingle, "")
	}
	return _GinLogrusSingle
}
