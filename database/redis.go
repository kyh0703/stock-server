package database

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnectRedis() (*redis.Client, error) {
	var err error
	// create new client
	Redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	// ping test
	_, err = Redis.Ping().Result()
	if err != nil {
		return nil, err
	}
	return Redis, err
}
