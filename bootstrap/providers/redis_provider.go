package providers

import (
	"context"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/logs"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

// RedisProvider @Bean("redis")
type RedisProvider struct {
	*services.Config `inject:"config, database"`
	dbs              map[string]*services.Redis
}

func (m *RedisProvider) Init() {
	m.dbs = make(map[string]*services.Redis)

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
			Addr:        config.GetString("host", "127.0.0.1") + ":" + config.GetString("port", "6379"),
			Password:    config.GetString("password", ""), // no password set
			DB:          config.GetInt("database", 0),     // use default DB
			ReadTimeout: time.Duration(config.GetInt("read_timeout", -1)),
		})

		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
		defer cancel()
		cmd := db.Ping(ctx)
		if cmd.Err() != nil {
			log.Errorf("redis connect err, %v", cmd.Err())
			panic(cmd.Err())
		}
		db.AddHook(&logs.Hook{})
		m.dbs[name.(string)] = &services.Redis{
			Client: db,
		}
	}
}

func (m *RedisProvider) GetBean(alias string) interface{} {
	return m.dbs[alias]
}
