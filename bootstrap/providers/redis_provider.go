package providers

import (
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-redis/redis/v8"
)

// RedisProvider @Bean("redis")
type RedisProvider struct {
	*services.Config `inject:"config, database"`
	dbs              map[string]*redis.Client
}

func (m *RedisProvider) Init() {
	m.dbs = make(map[string]*redis.Client)

	for name, dataT := range m.GetKey("connections") {
		data, ok := dataT.(map[interface{}]interface{})

		if !ok {
			continue
		}
		driver := data["driver"].(string)
		if driver != "redis" {
			continue
		}
		config := services.NewConfig(data)

		db := redis.NewClient(&redis.Options{
			Addr:     config.GetString("ip", "127.0.0.1") + ":" + config.GetString("port", "6379"),
			Password: config.GetString("password", ""), // no password set
			DB:       config.GetInt("database", 0),     // use default DB
		})

		m.dbs[name.(string)] = db
	}
}

func (m *RedisProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
