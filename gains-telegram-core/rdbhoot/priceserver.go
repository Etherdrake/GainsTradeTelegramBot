package rdbhoot

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

var ctx = context.Background()

func (rdbHoot *HootRedisClient) GetPrctChange(pairIndex int) (float64, error) {
	// Generate the key used in Redis based on pairIndex
	key := fmt.Sprintf("prctchange:%d", pairIndex)

	// Fetch the value from Redis
	value, err := rdbHoot.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("REDIS ERR", err)
		if errors.Is(err, redis.Nil) {
			log.Printf("No price found for pairIndex %d", pairIndex)
			return 0, fmt.Errorf("no price found for pairIndex %d", pairIndex)
		}
		return 0, err
	}

	// Convert the string value to float64
	price, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Error parsing price value for pairIndex %d: %s", pairIndex, value)
		return 0, err
	}
	return price, nil
}

func (rdbHoot *HootRedisClient) GetPrice(pairIndex int) (float64, error) {
	// Generate the key used in Redis based on pairIndex
	key := fmt.Sprintf("price:%d", pairIndex)

	// Fetch the value from Redis
	value, err := rdbHoot.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("REDIS ERR", err)
		if errors.Is(err, redis.Nil) {
			log.Printf("No price found for pairIndex %d", pairIndex)
			return 0, fmt.Errorf("no price found for pairIndex %d", pairIndex)
		}
		return 0, err
	}

	// Convert the string value to float64
	price, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Printf("Error parsing price value for pairIndex %d: %s", pairIndex, value)
		return 0, err
	}

	return price, nil
}
