package services

import "strings"

type Config struct {
	m map[interface{}]interface{}
}

func NewConfig(m map[interface{}]interface{}) *Config {
	return &Config{
		m: m,
	}
}

func (c *Config) Get() map[interface{}]interface{} {
	return c.m
}

func (c *Config) GetKey(key string) map[interface{}]interface{} {
	arr := strings.Split(key, ".")

	m := c.m
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := m[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case string:
					return v.(map[interface{}]interface{})
				default:
					return nil
				}
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				return nil
			}
		}
	}

	return m
}

func (c *Config) GetString(key string, def ...string) string {
	arr := strings.Split(key, ".")

	m := c.m
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := m[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case string:
					return v.(string)
				default:
					if len(def) == 0 {
						return ""
					}
					return def[0]
				}
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				if len(def) == 0 {
					return ""
				}
				return def[0]
			}
		}
	}

	if len(def) == 0 {
		return ""
	}
	return def[0]
}

func (c *Config) GetInt(key string, def ...int) int {
	arr := strings.Split(key, ".")

	m := c.m
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := m[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case string:
					return v.(int)
				default:
					if len(def) == 0 {
						return 0
					}
					return def[0]
				}
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				if len(def) == 0 {
					return 0
				}
				return def[0]
			}
		}
	}

	if len(def) == 0 {
		return 0
	}
	return def[0]
}

func (c *Config) GetBool(key string, def ...bool) bool {
	arr := strings.Split(key, ".")

	m := c.m
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := m[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case string:
					return v.(bool)
				default:
					if len(def) == 0 {
						return false
					}
					return def[0]
				}
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				if len(def) == 0 {
					return false
				}
				return def[0]
			}
		}
	}

	if len(def) == 0 {
		return false
	}
	return def[0]
}
