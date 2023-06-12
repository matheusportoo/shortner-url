package redis_service

import (
	"github.com/redis/go-redis/v9"
)

var RdsConnection *redis.Client

func CreateConnection() *redis.Client {
	RdsConnection = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return RdsConnection
}
