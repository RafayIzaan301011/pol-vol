package redis

import (
	"context"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RateLimit struct {
	ClusterClient *redis.ClusterClient
	MaxReq        int
}

func NewRateLimiter(clusterClient *redis.ClusterClient, maxReq int) *RateLimit {
	return &RateLimit{
		ClusterClient: clusterClient,
		MaxReq:        maxReq,
	}
}

func (rl *RateLimit) Allow(ctx context.Context, key string) (bool, error) {
	limiter := redis_rate.NewLimiter(rl.ClusterClient)
	res, err := limiter.Allow(ctx, key, redis_rate.PerMinute(rl.MaxReq))
	if err != nil {
		return true, err
	}

	return res.Allowed > 0, nil
}

func (rl *RateLimit) GetUsage(ctx context.Context, key string) (int, error) {
	limiter := redis_rate.NewLimiter(rl.ClusterClient)

	res, err := limiter.AllowN(ctx, key, redis_rate.PerMinute(rl.MaxReq), 0)
	if err != nil {
		return 0, err
	}

	consumed := rl.MaxReq - res.Remaining

	return consumed, nil
}
