package public

import (
	"github.com/go-home-admin/home/bootstrap/utils"
	"testing"
)

func TestController_Home(t *testing.T) {
	receiver := InitializeNewControllerProvider()
	u := receiver.user.WhereId(449).Find()

	utils.Dump(u)
}
