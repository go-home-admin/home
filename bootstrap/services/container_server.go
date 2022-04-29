package services

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"os"
	"os/signal"
)

// Container @Bean
type Container struct{}

// Run 统一启动服务
func (a *Container) Run(servers []constraint.KernelServer) {
	app.RunBoot()

	for _, server := range servers {
		go server.Run()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	app.RunExit()
}
