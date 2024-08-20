package config

import (
	"93HW/utils"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = utils.NewRedisClient()

	status := RedisClient.Ping(utils.Ctx)

	if status.Err() != nil {
		log.Fatalf("Could not connect to Redis: %v", status.Err())
	}
	log.Println("COnnected to Redis succesfully")
}
