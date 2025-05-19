package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, dest interface{}) error
	Delete(ctx context.Context, key string) error
}

// NewCache creates a new RedisCache with a given TTL (time-to-live).
func NewCache(addr string, ttl time.Duration) RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCacheImpl{client: client, ttl: ttl}
}
