package config

import "github.com/go-redis/redis"

func ConnectRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}
