package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	ClusterClient *redis.ClusterClient
}

func NewCache(clusterClient *redis.ClusterClient) *Cache {
	return &Cache{
		ClusterClient: clusterClient,
	}
}

// set single value with a ttl
func (c *Cache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return c.ClusterClient.Set(ctx, key, value, ttl).Err()
}

// get single value
func (c *Cache) Get(ctx context.Context, key string) error {
	return c.ClusterClient.Get(ctx, key).Err()
}

// set multiple values without a ttl
func (c *Cache) MSet(ctx context.Context, values ...any) error {
	return c.ClusterClient.MSet(ctx, values...).Err()
}

// get multiple values
func (c *Cache) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return c.ClusterClient.MGet(ctx, keys...).Result()
}
