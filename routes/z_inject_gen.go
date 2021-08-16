// 代码由home-admin生成, 不需要编辑它

package routes

import (
	"github.com/go-home-admin/home/app/http/admin/public"
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var AdminRoutesSingle *AdminRoutes
var RoutesSingle *Routes

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

		var temp interface{} = AdminRoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
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

		var temp interface{} = RoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return RoutesSingle
}
