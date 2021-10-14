package main

import (
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/app/queue"
	"github.com/go-home-admin/home/bootstrap/constraint"
)

func main() {
	app := providers.InitializeNewAppProvider()

	app.Run([]constraint.KernelServer{
		// http服务
		http.GetServer(),
		// Job消费服务
		queue.GetServer(),
	})
}
