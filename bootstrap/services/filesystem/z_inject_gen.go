// gen for home toolset
package filesystem

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _LocalSingle *Local

func GetAllProvider() []interface{} {
	return []interface{}{
		NewLocal(),
	}
}

func NewLocal() *Local {
	if _LocalSingle == nil {
		_LocalSingle = &Local{}
		providers.AfterProvider(_LocalSingle, "")
	}
	return _LocalSingle
}
