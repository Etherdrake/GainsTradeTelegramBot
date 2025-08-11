package swaps

import (
	"HootCore/rdbhoot"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// EditSwapInfo This function retrieves the SwapInfos struct from Redis, checks if the chain exists,
// replaces the SwapInfo for that chain with the new SwapInfo, and writes the updated SwapInfos struct back to Redis.
// If there is an error at any point, it returns the error.
func EditSwapInfo(ctx context.Context, rdb *rdbhoot.HootRedisClient, guid string, chain uint64, newSwapInfo SwapInfo) error {
	log.Println("Editing swap information...")

	// Get SwapInfos struct from Redis
	swapInfos, err := GetSwapConfig(ctx, rdb, guid)
	if err != nil {
		return fmt.Errorf("error retrieving SwapInfos: %v", err)
	}

	// Check if chain exists in SwapInfos
	if _, ok := swapInfos.Chains[chain]; !ok {
		return fmt.Errorf("chain %s does not exist in SwapInfos", chain)
	}

	// Replace SwapInfo for chain with new SwapInfo
	swapInfos.Chains[chain] = newSwapInfo

	// Marshal updated SwapInfos struct to JSON
	swapInfosJSON, err := json.Marshal(swapInfos)
	if err != nil {
		return fmt.Errorf("error marshaling updated SwapInfos struct to JSON: %v", err)
	}

	// Write updated SwapInfos JSON to Redis
	if redisErr := rdb.JSONSet(ctx, guid, ".", swapInfosJSON).Err(); redisErr != nil {
		return fmt.Errorf("error writing updated SwapInfos JSON to Redis: %v", redisErr)
	}
	return nil
}
