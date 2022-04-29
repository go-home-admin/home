package services

import (
	"github.com/go-home-admin/home/bootstrap/services/app"
	"strings"
)

// Init 单元测试辅助
// 如果一个单元测试连配置都都不使用, 想来也无需初始化了
var Init bool

type Config struct {
	m map[interface{}]interface{}
}

func NewConfig(m map[interface{}]interface{}) *Config {
	if !Init {
		app.RunBoot()
	}

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
		s := arr[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case map[interface{}]interface{}:
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
		s := arr[i]
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
		s := arr[i]
		if v, ok := m[s]; ok {
			if ll == i {
				switch v.(type) {
				case int:
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
		s := arr[i]
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

func (c *Config) GetConfig(key string) *Config {
	arr := strings.Split(key, ".")

	m := c.m
	lc := len(arr)
	ll := lc - 1
	for i := 0; i < lc; i++ {
		s := arr[i]
		if v, ok := m[s]; ok {
			if ll == i {
				return NewConfig(v.(map[interface{}]interface{}))
			}

			val, ook := v.(map[interface{}]interface{})
			if ook {
				m = val
			} else {
				return nil
			}
		} else {
			return nil
		}
	}

	return nil
}
