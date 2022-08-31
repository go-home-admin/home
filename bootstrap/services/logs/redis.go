package logs

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// Hook 记录log
type Hook struct{}

func (h *Hook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	log.WithFields(log.Fields{"type": "redis"}).Debug(cmd.String())
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
