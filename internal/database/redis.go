package database

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnectRedis(host string) (*redis.Client, error) {
	// create new client
	Redis = redis.NewClient(&redis.Options{
		Addr: host,
	})

	// ping test
	if _, err := Redis.Ping().Result(); err != nil {
		return nil, err
	}

	return Redis, nil
}
