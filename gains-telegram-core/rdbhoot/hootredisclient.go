package rdbhoot

import "github.com/redis/go-redis/v9"

type HootRedisClient struct {
	*redis.Client
}
