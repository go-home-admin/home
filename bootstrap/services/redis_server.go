package services

import "github.com/go-redis/redis/v8"

type Redis struct {
	*redis.Client
}
