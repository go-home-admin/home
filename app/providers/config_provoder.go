package providers

import (
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"gopkg.in/ini.v1"
)

// Config 外部其他服务的配置依赖提供
// @Bean
type Config struct {
	iniConfig *Ini `inject:""`
}

func NewConfigProvider(ini *Ini) *Config {
	return &Config{
		iniConfig: ini,
	}
}

func (g *Config) GetServiceConfig(service string) SessionService {
	return SessionService{
		session: g.iniConfig.Session(service),
	}
}

type SessionService struct {
	session *ini.Section
}

func (c *SessionService) GetString(key string) string {
	return c.session.Key(key).String()
}

func (c *SessionService) GetInt(key string) int {
	i, err := c.session.Key(key).Int()
	if err != nil {
		logs.Error(err)
		return 0
	}
	return i
}

func (c *SessionService) GetBool(key string) bool {
	i, err := c.session.Key(key).Bool()
	if err != nil {
		logs.Error(err)
		return false
	}
	return i
}
