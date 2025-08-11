package rdblocker

import (
	"HootCore/rdbhoot"
	"context"
	"log"
)

// WipeCache removes all keys from Redis
func WipeCache(ctx context.Context, client *rdbhoot.HootRedisClient) error {
	// FLUSHDB command removes all keys from the currently selected Redis database
	err := client.FlushDB(ctx).Err()
	if err != nil {
		log.Printf("Error wiping cache: %v", err)
		return err
	}
	return nil
}

func WipeCacheSelectively(ctx context.Context, client *rdbhoot.HootRedisClient, key string) error {
	// Del command removes the specified keys
	err := client.Del(ctx, key).Err()
	if err != nil {
		log.Printf("Error wiping active_users set: %v", err)
		return err
	}
	return nil
}
