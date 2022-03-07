// gen for home toolset
package logs

var _GinLogrusSingle *GinLogrus

func GetAllProvider() []interface{} {
	return []interface{}{
		NewGinLogrus(),
	}
}

func NewGinLogrus() *GinLogrus {
	if _GinLogrusSingle == nil {
		GinLogrus := &GinLogrus{}
		_GinLogrusSingle = GinLogrus
	}
	return _GinLogrusSingle
}
