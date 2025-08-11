package rdblistener

import "github.com/redis/go-redis/v9"

type HootRedisClient struct {
	*redis.Client
}

// NewHootRedisClient is a constructor function that initializes the redis.Client once.
func NewHootRedisClient(addr string) *HootRedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		// Other options
	})

	return &HootRedisClient{Client: client}
}
