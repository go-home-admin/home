package services

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"os"
	"os/signal"
)

// @Bean
type AppServer struct {
}

// Run 统一启动服务
func (a *AppServer) Run(servers []constraint.KernelServer) {
	for _, server := range servers {
		server.Init()
	}
	for _, server := range servers {
		go server.Run()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	for _, server := range servers {
		server.Exit()
	}
}
