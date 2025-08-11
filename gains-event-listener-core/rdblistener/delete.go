package rdblistener

import (
	"context"
	"log"
)

// DeleteTradeByID deletes a trade document by its ID.
func (rdbHoot *HootRedisClient) DeleteTradeByID(ctx context.Context, trader, id string) error {
	key := "trades:" + trader + ":" + id
	delErr := rdbHoot.Del(ctx, key).Err()
	if delErr != nil {
		log.Printf("Error deleting key %s: %v\n", key, delErr)
	}
	return nil
}

// DeleteOrderByID DeleteTradeByID deletes a trade document by its ID.
func (rdbHoot *HootRedisClient) DeleteOrderByID(ctx context.Context, trader, id string) error {
	key := "orders:" + trader + ":" + id
	delErr := rdbHoot.Del(ctx, key).Err()
	if delErr != nil {
		log.Printf("Error deleting key %s: %v\n", key, delErr)
	}
	return nil
}
