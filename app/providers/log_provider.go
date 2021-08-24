package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"github.com/sirupsen/logrus"
)

// @Bean
type Log struct {
	ginLog *logs.GinLogrus `inject:""`
	conf   *Config         `inject:""`
}

func (receiver *Log) Init() {
	switch app.GetEnvironment() {
	case app.EnvironmentLocal:
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(false)
	case app.EnvironmentTesting:
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(false)
	default:
		// 设置等级
		logrus.SetLevel(logrus.WarnLevel)
		// 是否打印调用
		logrus.SetReportCaller(true)
		// 生产等环境, 使用json方便采集
		logrus.SetFormatter(&logrus.JSONFormatter{})
		// gin的log, 非调试环境下, 也统一同一个log输出
		gin.DefaultWriter = receiver.ginLog
	}
}
