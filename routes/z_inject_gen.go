// 代码由home-admin生成, 不需要编辑它

package routes

import (
	"github.com/go-home-admin/home/app/http/admin/admin_user"
	"github.com/go-home-admin/home/app/http/admin/public"
	public_1 "github.com/go-home-admin/home/app/http/api/public"
	"github.com/go-home-admin/home/app/http/swagger/doc"
	home_constraint "github.com/go-home-admin/home/bootstrap/constraint"
)

var AdminRoutesSingle *AdminRoutes
var RoutesSingle *Routes
var ApiRoutesSingle *ApiRoutes
var SwaggerRoutesSingle *SwaggerRoutes

func NewAdminRoutesProvider(admin_user *admin_user.Controller, public *public.Controller) *AdminRoutes {
	AdminRoutes := &AdminRoutes{}
	AdminRoutes.admin_user = admin_user
	AdminRoutes.public = public
	return AdminRoutes
}

func InitializeNewAdminRoutesProvider() *AdminRoutes {
	if AdminRoutesSingle == nil {
		AdminRoutesSingle = NewAdminRoutesProvider(
			admin_user.InitializeNewControllerProvider(),

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

func NewRoutesProvider(AdminRoutes *AdminRoutes, ApiRoutes *ApiRoutes, SwaggerRoutes *SwaggerRoutes) *Routes {
	Routes := &Routes{}
	Routes.AdminRoutes = AdminRoutes
	Routes.ApiRoutes = ApiRoutes
	Routes.SwaggerRoutes = SwaggerRoutes
	return Routes
}

func InitializeNewRoutesProvider() *Routes {
	if RoutesSingle == nil {
		RoutesSingle = NewRoutesProvider(
			InitializeNewAdminRoutesProvider(),

			InitializeNewApiRoutesProvider(),

			InitializeNewSwaggerRoutesProvider(),
		)

		var temp interface{} = RoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return RoutesSingle
}

func NewApiRoutesProvider(public *public_1.Controller) *ApiRoutes {
	ApiRoutes := &ApiRoutes{}
	ApiRoutes.public = public
	return ApiRoutes
}

func InitializeNewApiRoutesProvider() *ApiRoutes {
	if ApiRoutesSingle == nil {
		ApiRoutesSingle = NewApiRoutesProvider(
			public_1.InitializeNewControllerProvider(),
		)

		var temp interface{} = ApiRoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return ApiRoutesSingle
}

func NewSwaggerRoutesProvider(doc *doc.Controller) *SwaggerRoutes {
	SwaggerRoutes := &SwaggerRoutes{}
	SwaggerRoutes.doc = doc
	return SwaggerRoutes
}

func InitializeNewSwaggerRoutesProvider() *SwaggerRoutes {
	if SwaggerRoutesSingle == nil {
		SwaggerRoutesSingle = NewSwaggerRoutesProvider(
			doc.InitializeNewControllerProvider(),
		)

		var temp interface{} = SwaggerRoutesSingle
		construct, ok := temp.(home_constraint.Construct)
		if ok {
			construct.Init()
		}
	}

	return SwaggerRoutesSingle
}
