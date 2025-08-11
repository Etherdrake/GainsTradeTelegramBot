package swaps

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func CheckSwapConfigSet(ctx context.Context, rdb *redis.Client, guid string) (bool, error) {
	log.Println("Checking if swap configuration exists...")

	// Check if guid exists in Redis
	val, err := rdb.Exists(ctx, guid).Result()
	if err != nil {
		return false, fmt.Errorf("error checking if guid exists in Redis: %v", err)
	}

	// If val is 1, the guid exists in Redis
	return val == 1, nil
}
