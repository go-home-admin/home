package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"github.com/sirupsen/logrus"
)

// @Bean
type Log struct {
	ginLog *logs.GinLogrus `inject:""`
}

func (receiver *Log) Init() {
	// 统重定向gin的log
	gin.DefaultWriter = receiver.ginLog

	logs.Logger.SetLevel(logrus.TraceLevel)
}
