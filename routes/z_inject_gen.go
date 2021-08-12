// 代码由home-admin生成, 不需要编辑它

package routes

import (
	"github.com/go-home-admin/home/app/http/admin/public"
)

var AdminRoutesSingle *AdminRoutes

func NewAdminRoutesProvider(public *public.Controller) *AdminRoutes {
	AdminRoutes := &AdminRoutes{}
	AdminRoutes.public = public
	return AdminRoutes
}

func InitializeNewAdminRoutesProvider() *AdminRoutes {
	if AdminRoutesSingle == nil {
		AdminRoutesSingle = NewAdminRoutesProvider(
			public.InitializeNewControllerProvider(),
		)
	}

	return AdminRoutesSingle
}
