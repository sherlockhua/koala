package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache struct defines the cache.
type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string, dest interface{}) error
	Delete(key string) error
}

type CacheConfig struct {
	Addr           string
	TTL            time.Duration
	Host           string
	Port           int
	Password       string
	DB             int
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
}

// NewRedisCache creates a new RedisCache with a given TTL (time-to-live).
func NewCache(conf *CacheConfig) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:     conf.Password, // No password set
		DB:           conf.DB,       // use default DB
		DialTimeout:  time.Duration(conf.ConnectTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second,
	})
	return &RedisCache{client: client, ttl: conf.TTL}
}

// Set caches a value in Redis.
func (r *RedisCache) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, data, r.ttl).Err()
}

// Get retrieves a cached value from Redis.
func (r *RedisCache) Get(key string, dest interface{}) error {
	result, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil // Cache miss
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(result), dest)
}

// Delete removes a cached value from Redis.
func (r *RedisCache) Delete(key string) error {
	return r.client.Del(context.Background(), key).Err()
}
