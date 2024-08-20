package utils

import (
    "context"
    "github.com/redis/go-redis/v9"
    "time"
)

var (
    Ctx               = context.Background()
    RateLimitDuration = time.Minute * 1
)

func NewRedisClient() *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}
