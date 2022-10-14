package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// LogProvider @Bean
type LogProvider struct {
	*services.Config `inject:"config, app"`
}

func (l *LogProvider) Init() {
	level, err := logrus.ParseLevel(l.GetString("log.level", "trace"))
	if err == nil {
		logrus.SetLevel(level)
	} else {
		panic(err)
	}
	logrus.SetReportCaller(true)

	switch l.GetString("log.formatter", "text") {
	case "text":
		// 通常这里是单行输出好看点
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	default:
		// 通常这里是正式环境输出到文件的
		logrus.SetFormatter(&logrus.JSONFormatter{
			DisableTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	}
	// 是否输出到文件
	logPath := l.GetString("log.path", "")
	if logPath != "" {
		dir := path.Dir(logPath)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			logrus.Warning("无法创建log目录", err, dir)
		}

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
			//下面配置日志每隔 1 天轮转一个新文件，保留最近 30 天的日志文件，多余的自动清理掉。
			writer, _ := rotatelogs.New(
				logPath+".%Y%m%d",
				rotatelogs.WithLinkName(logPath),
				rotatelogs.WithMaxAge(time.Duration(24*30)*time.Hour),
				rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
			)
			log.SetOutput(writer)
		} else {
			logrus.Warnf("LOG文件无法打开记录, %v", logPath)
		}
	}
}
