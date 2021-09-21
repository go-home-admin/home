package public

import (
	"github.com/go-home-admin/home/app/entity/lrs_manager"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/bootstrap/utils"
	"testing"
)

func TestController_Home(t *testing.T) {
	u := &lrs_manager.OrmUsers{}
	providers.InitializeNewMysqlProvider().DB().First(u)

	utils.Dump(u.UserType)
}
