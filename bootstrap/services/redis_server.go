package services

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

type Redis struct {
	*redis.Client
}

func (r Redis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.Client.Set(context.Background(), key, value, expiration)
}

func (r Redis) Get(key string) *redis.StringCmd {
	return r.Client.Get(context.Background(), key)
}

func (r Redis) GetString(key string) (string, bool) {
	cmd := r.Client.Get(context.Background(), key)
	err := cmd.Err()
	if errors.Is(err, redis.Nil) {
		return "", false
	} else if err != nil {
		log.Error(err)
		return "", false
	}
	return cmd.Val(), true
}

func (r Redis) GetInt(key string) (int, bool) {
	i, err := r.Client.Get(context.Background(), key).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, false
		}
		log.Errorf("GetInt %v", err)
	}
	return i, true
}

func (r Redis) GetInt64(key string) (int64, bool) {
	i, err := r.Client.Get(context.Background(), key).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, false
		}
		log.Errorf("GetInt64 %v", err)
	}
	return i, true
}

func (r Redis) GetFloat32(key string) (float32, bool) {
	i, err := r.Client.Get(context.Background(), key).Float32()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, false
		}
		log.Errorf("GetFloat32 %v", err)
	}
	return i, true
}

func (r Redis) GetFloat64(key string) (float64, bool) {
	i, err := r.Client.Get(context.Background(), key).Float64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, false
		}
		log.Errorf("GetFloat32 %v", err)
	}
	return i, true
}

func (r Redis) GetBool(key string) (bool, bool) {
	i, err := r.Client.Get(context.Background(), key).Bool()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, false
		}
		log.Errorf("GetFloat32 %v", err)
	}
	return i, true
}

func (r Redis) Incr(key string) (int64, bool) {
	cmd := r.Client.Incr(context.Background(), key)
	err := cmd.Err()
	if errors.Is(err, redis.Nil) {
		return 0, true
	} else if err != nil {
		log.Error(err)
		return 0, false
	}
	return cmd.Val(), true
}
