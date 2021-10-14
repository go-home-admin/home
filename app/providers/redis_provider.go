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

	r.client.AddHook(&Hook{})
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

type Hook struct {
}

func (h *Hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	log.Info(cmd.String())

	return ctx, nil
}

func (h *Hook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	return nil
}

func (h *Hook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (h *Hook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}
