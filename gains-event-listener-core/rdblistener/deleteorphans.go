package rdblistener

import (
	"context"
	"log"
)

func (rdbHoot *HootRedisClient) DeleteTradeKeysNotInMap(ctx context.Context, inputKeys map[string]bool) error {
	// Fetch all keys that match the pattern "trade:%s:%s" using KEYS command
	keysInRedis, err := rdbHoot.Keys(ctx, "trades:*").Result()
	if err != nil {
		return err
	}

	// Delete keys in Redis that are not in the inputKeys map
	for _, key := range keysInRedis {
		if _, ok := inputKeys[key]; !ok {
			delErr := rdbHoot.Del(ctx, key).Err()
			if delErr != nil {
				log.Printf("Error deleting key %s: %v\n", key, delErr)
			}
			log.Println("Deleted Trade: ", key)
		}
	}

	return nil
}

func (rdbHoot *HootRedisClient) DeleteOrderKeysNotInMap(ctx context.Context, inputKeys map[string]bool) error {
	// Fetch all keys that match the pattern "orders:%s:%s" using KEYS command
	keysInRedis, err := rdbHoot.Keys(ctx, "orders:*").Result()
	if err != nil {
		return err
	}

	// Delete keys in Redis that are not in the inputKeys map
	for _, key := range keysInRedis {
		if _, ok := inputKeys[key]; !ok {
			delErr := rdbHoot.Del(ctx, key).Err()
			if delErr != nil {
				log.Printf("Error deleting key %s: %v\n", key, delErr)
			}
			log.Println("Deleted Order: ", key)
		}
	}

	return nil
}
