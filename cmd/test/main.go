package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ErikDPrince/bookmark-management/pkg/redis"
)

func main() {
	ctx := context.Background()

	redisClient, err := redis.NewClient("")
	if err != nil {
		panic(err)
	}

	if err := redisClient.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("redis ping failed: %w", err))
	}

	if err := redisClient.Set(ctx, "test", "test", time.Hour).Err(); err != nil {
		panic(fmt.Errorf("redis set failed: %w", err))
	}

	val, err := redisClient.Get(ctx, "test").Result()
	if err != nil {
		panic(fmt.Errorf("redis get after set failed: %w", err))
	}

	fmt.Printf("redis connected and wrote key: test=%s\n", val)
}
