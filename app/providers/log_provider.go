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
	// 设置等级
	logrus.SetLevel(logrus.TraceLevel)
	switch GetEnvironment() {
	case EnvironmentLocal:
		logrus.SetReportCaller(false)
	case EnvironmentTesting:
		logrus.SetReportCaller(false)
	default:
		// 是否打印调用
		logrus.SetReportCaller(true)
		// 生产等环境, 使用json方便采集
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
