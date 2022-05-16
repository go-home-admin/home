// gen for home toolset
package queues

import (
	providers "github.com/go-home-admin/home/bootstrap/providers"
)

var _ElectionCloseSingle *ElectionClose

func GetAllProvider() []interface{} {
	return []interface{}{
		NewElectionClose(),
	}
}

func NewElectionClose() *ElectionClose {
	if _ElectionCloseSingle == nil {
		_ElectionCloseSingle = &ElectionClose{}
		providers.AfterProvider(_ElectionCloseSingle, "election_close")
	}
	return _ElectionCloseSingle
}
