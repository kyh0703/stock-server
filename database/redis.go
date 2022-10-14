package database

import (
	"sync"

	"github.com/go-redis/redis"
)

var (
	rc *redis.Client
	ro sync.Once
)

func Redis() *redis.Client {
	return rc
}

func ConnectRedis() (*redis.Client, error) {
	var err error
	ro.Do(func() {
		rc = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		})
		_, err := rc.Ping().Result()
		if err != nil {
			return
		}
	})
	return rc, err
}
