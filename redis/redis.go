package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Nodes    int
	Host     string
	Port     string
	Password string
	DB       int
}

// New initializes and returns a Redis client based on the provided configuration.
func NewRedisCluster(ctx context.Context, cfg Config) *redis.ClusterClient {
	addresses := make([]string, cfg.Nodes)
	basePort, _ := strconv.Atoi(cfg.Port)
	for i := 0; i < cfg.Nodes; i++ {
		port := strconv.Itoa(basePort + i)
		addresses[i] = fmt.Sprintf("%s:%s", cfg.Host, port)
	}

	opts := &redis.ClusterOptions{
		Addrs:        addresses,
		ClientName:   "polvol-redis-client",
		Password:     cfg.Password,
		ReadOnly:     true,
		MinIdleConns: 10,
		PoolSize:     100,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	rdb := redis.NewClusterClient(opts)

	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})

	if err != nil {
		panic(err)
	}

	return rdb
}
