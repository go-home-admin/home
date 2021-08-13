// 代码由home-admin生成, 不需要编辑它

package routes

import (
	"github.com/go-home-admin/home/app/http/admin/open2"
	"github.com/go-home-admin/home/app/http/admin/public"
	"github.com/go-home-admin/home/app/http/admin/test"
)

var AdminRoutesSingle *AdminRoutes
var RoutesSingle *Routes

func NewAdminRoutesProvider(public *public.Controller, open2 *open2.Controller, test *test.Controller) *AdminRoutes {
	AdminRoutes := &AdminRoutes{}
	AdminRoutes.public = public
	AdminRoutes.open2 = open2
	AdminRoutes.test = test
	return AdminRoutes
}

func InitializeNewAdminRoutesProvider() *AdminRoutes {
	if AdminRoutesSingle == nil {
		AdminRoutesSingle = NewAdminRoutesProvider(
			public.InitializeNewControllerProvider(),

			open2.InitializeNewControllerProvider(),

			test.InitializeNewControllerProvider(),
		)
	}

	return AdminRoutesSingle
}

func NewRoutesProvider(AdminRoutes *AdminRoutes) *Routes {
	Routes := &Routes{}
	Routes.AdminRoutes = AdminRoutes
	return Routes
}

func InitializeNewRoutesProvider() *Routes {
	if RoutesSingle == nil {
		RoutesSingle = NewRoutesProvider(
			InitializeNewAdminRoutesProvider(),
		)
	}

	return RoutesSingle
}
