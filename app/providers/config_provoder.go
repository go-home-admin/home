package providers

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
	"gopkg.in/ini.v1"
)

// Config 外部其他服务的配置依赖提供
// @Bean
type Config struct {
	iniConfig *Ini `inject:""`
}

func (c *Config) Init() {
	app.SetEnvironment(c.GetString("environment"))
}

func (c *Config) GetString(key string) string {
	return c.iniConfig.Session("app").Key(key).String()
}

func (c *Config) GetInt(key string) int {
	i, err := c.iniConfig.Session("app").Key(key).Int()
	if err != nil {
		return 0
	}
	return i
}

func (c *Config) GetBool(key string, def bool) bool {
	s := c.iniConfig.Session("app").Key(key).String()
	if s == "true" {
		return true
	}
	return def
}

func (c *Config) GetServiceConfig(service string) *SessionService {
	return &SessionService{
		session: c.iniConfig.Session(service),
	}
}

type SessionService struct {
	session *ini.Section
}

func (c *SessionService) GetString(key string) string {
	return c.session.Key(key).String()
}

func (c *SessionService) GetInt(key string, def int) int {
	i, err := c.session.Key(key).Int()
	if err != nil {
		return def
	}
	return i
}

func (c *SessionService) GetBool(key string) bool {
	s := c.session.Key(key).String()
	if s == "true" {
		return true
	}
	return false
}
