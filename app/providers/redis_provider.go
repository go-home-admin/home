package providers

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// Redis @Bean
type Redis struct {
	conf   *Config `inject:""`
	client *redis.Client
}

func (r *Redis) Init() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	cmd := r.client.Ping(context.Background())
	if cmd.Err() != nil {
		log.Error(cmd.Err())
	}
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}
