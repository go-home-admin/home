package providers

import (
	"github.com/go-redis/redis/v8"
)

// Redis @Bean
type Redis struct {
	conf   *Config `inject:""`
	client *redis.Client
}

func (r *Redis) Init() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
