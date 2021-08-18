// 代码由home-admin生成, 不需要编辑它

package routes

import (
	"github.com/go-home-admin/home/app/http/admin/public"
	public_1 "github.com/go-home-admin/home/app/http/toolset/public"
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var AdminRoutesSingle *AdminRoutes
var RoutesSingle *Routes
var ToolsetRoutesSingle *ToolsetRoutes

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

func NewRoutesProvider(AdminRoutes *AdminRoutes, ToolsetRoutes *ToolsetRoutes) *Routes {
	Routes := &Routes{}
	Routes.AdminRoutes = AdminRoutes
	Routes.ToolsetRoutes = ToolsetRoutes
	return Routes
}

func InitializeNewRoutesProvider() *Routes {
	if RoutesSingle == nil {
		RoutesSingle = NewRoutesProvider(
			InitializeNewAdminRoutesProvider(),

			InitializeNewToolsetRoutesProvider(),
		)

		var temp interface{} = RoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return RoutesSingle
}

func NewToolsetRoutesProvider(public *public_1.Controller) *ToolsetRoutes {
	ToolsetRoutes := &ToolsetRoutes{}
	ToolsetRoutes.public = public
	return ToolsetRoutes
}

func InitializeNewToolsetRoutesProvider() *ToolsetRoutes {
	if ToolsetRoutesSingle == nil {
		ToolsetRoutesSingle = NewToolsetRoutesProvider(
			public_1.InitializeNewControllerProvider(),
		)

		var temp interface{} = ToolsetRoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return ToolsetRoutesSingle
}
