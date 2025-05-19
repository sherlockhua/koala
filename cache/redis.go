package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache struct defines the cache.
type RedisCacheImpl struct {
	client *redis.Client
	ttl    time.Duration
}

// Set caches a value in Redis.
func (r *RedisCacheImpl) Set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, data, r.ttl).Err()
}

// Get retrieves a cached value from Redis.
func (r *RedisCacheImpl) Get(ctx context.Context, key string, dest interface{}) error {
	result, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil // Cache miss
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(result), dest)
}

// Delete removes a cached value from Redis.
func (r *RedisCacheImpl) Delete(ctx context.Context, key string) error {
	return r.client.Del(context.Background(), key).Err()
}
