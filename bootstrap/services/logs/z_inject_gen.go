// gen for home toolset
package logs

import ()

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
