package logs

import (
	"bytes"
	"github.com/sirupsen/logrus"
)

type GinLogrus struct{}

func (g *GinLogrus) Write(p []byte) (n int, err error) {
	i := bytes.Index(p, []byte("[GIN-debug] "))
	if i == 0 {
		logrus.Debug(string(p))
	} else {
		logrus.Error(string(p))
	}
	return 0, nil
}
