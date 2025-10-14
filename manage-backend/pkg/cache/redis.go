package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg config.Redis) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &RedisClient{client: rdb}
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *RedisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

// GetClient 返回底层的 Redis 客户端
func (r *RedisClient) GetClient() *redis.Client {
	return r.client
}