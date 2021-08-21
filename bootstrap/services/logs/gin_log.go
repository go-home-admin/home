package logs

import (
	"bytes"
)

// @Bean
type GinLogrus struct {
}

func (g *GinLogrus) Write(p []byte) (n int, err error) {
	i := bytes.Index(p, []byte("[GIN-debug] "))
	if i == 0 {
		Debug(string(p))
	} else {
		Error(string(p))
	}
	return 0, nil
}
