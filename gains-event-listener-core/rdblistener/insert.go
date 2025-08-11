package rdblistener

import (
	"GainsListenerV4/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func (rdbHoot *HootRedisClient) InsertTrade(trade types.TradeItemMongo) error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Set timeout duration as needed
	defer cancel()                                                          // Ensure the cancel function is called to release resources

	// Serialize the trade struct to JSON
	tradeJSON, err := json.Marshal(trade)
	if err != nil {
		log.Println("Error serializing trade to JSON:", err)
		return err
	}

	// Use JSON.SET to insert the serialized trade into Redis
	key := fmt.Sprintf("trades:%s:%s", strings.ToLower(trade.Trade.User), trade.ID)
	statusCmd := rdbHoot.JSONSet(ctx, key, ".", string(tradeJSON))
	if statusCmd.Err() != nil {
		log.Println("Error inserting trade into Redis:", statusCmd.Err())
		return statusCmd.Err()
	}

	return nil
}

func (rdbHoot *HootRedisClient) InsertOrder(trade types.TradeItemMongo) error {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Set timeout duration as needed
	defer cancel()                                                          // Ensure the cancel function is called to release resources

	// Serialize the trade struct to JSON
	tradeJSON, err := json.Marshal(trade)
	if err != nil {
		log.Println("Error serializing trade to JSON:", err)
		return err
	}

	// Use JSON.SET to insert the serialized trade into Redis
	key := fmt.Sprintf("orders:%s:%s", strings.ToLower(trade.Trade.User), trade.ID)
	statusCmd := rdbHoot.JSONSet(ctx, key, ".", string(tradeJSON))
	if statusCmd.Err() != nil {
		log.Println("Error inserting trade into Redis:", statusCmd.Err())
		return statusCmd.Err()
	}

	return nil
}
