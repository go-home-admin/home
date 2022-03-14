package services

import "gopkg.in/ini.v1"

type Config struct {
	session *ini.Section
}

func (c *Config) GetString(key string) string {
	return c.session.Key(key).String()
}

func (c *Config) GetInt(key string, def int) int {
	i, err := c.session.Key(key).Int()
	if err != nil {
		return def
	}
	return i
}

func (c *Config) GetBool(key string) bool {
	s := c.session.Key(key).String()
	if s == "true" {
		return true
	}
	return false
}
