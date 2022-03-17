package main

import (
	"embed"
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/app/providers"
	"github.com/go-home-admin/home/app/queue"
	"github.com/go-home-admin/home/bootstrap/constraint"
	fp "github.com/go-home-admin/home/bootstrap/providers"
)

//go:embed config
var config embed.FS

func init() {
	fp.SetConfigDir(&config)
}

func main() {
	app := providers.NewApp()

	app.Run([]constraint.KernelServer{
		// http服务
		http.GetServer(),
		// Job消费服务
		queue.GetServer(),
	})
}
