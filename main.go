package main

import (
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
)

func main() {
	app := services.InitializeNewAppServerProvider()

	app.Run([]constraint.KernelServer{
		// http服务
		http.GetServer(),
	})
}
