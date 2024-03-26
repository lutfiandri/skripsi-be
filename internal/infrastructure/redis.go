package infrastructure

import "github.com/go-redis/redis/v7"

func NewRedisClient(address string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	return client
}
