package providers

import (
	"gopkg.in/ini.v1"
)

type Environment string

var (
	EnvironmentNow        Environment = ""
	EnvironmentLocal      Environment = "local"
	EnvironmentDev        Environment = "dev"
	EnvironmentTesting    Environment = "testing"
	EnvironmentStaging    Environment = "staging"
	EnvironmentProduction Environment = "production"
)

// Config 外部其他服务的配置依赖提供
// @Bean
type Config struct {
	iniConfig *Ini `inject:""`
}

func GetEnvironment() Environment {
	if EnvironmentNow == "" {
		conf := InitializeNewConfigProvider().GetServiceConfig("app")
		EnvironmentNow = Environment(conf.GetString("environment"))
	}
	return EnvironmentNow
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
		return 0
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
