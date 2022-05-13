package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
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

}
