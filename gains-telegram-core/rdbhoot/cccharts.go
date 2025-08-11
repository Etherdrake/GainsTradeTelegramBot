package rdbhoot

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (rdbHoot *HootRedisClient) SetActiveCharts(guid int64, isCharts bool) error {
	// Fetch the JSON struct for the given user from Redis
	userTradeJSON, err := rdbHoot.Get(ctx, "corecache:"+strconv.FormatInt(guid, 10)).Result()
	if err != nil {
		return fmt.Errorf("error fetching user trade data from Redis: %v", err)
	}

	// Unmarshal the JSON data into UserTrade struct
	var coreCache HootCoreCache
	err = json.Unmarshal([]byte(userTradeJSON), &coreCache)
	if err != nil {
		return fmt.Errorf("error unmarshaling user trade data: %v", err)
	}

	// Update the ActiveCharts field
	coreCache.ActiveCharts = isCharts

	// Marshal the updated UserTrade struct back to JSON
	updatedUserTradeJSON, err := json.Marshal(coreCache)
	if err != nil {
		return fmt.Errorf("error marshaling updated user trade data to JSON: %v", err)
	}

	// Store the updated JSON data in Redis
	_, err = rdbHoot.Set(ctx, "userTrade:"+strconv.FormatInt(guid, 10), updatedUserTradeJSON, 0).Result()
	if err != nil {
		return fmt.Errorf("error storing updated user trade data in Redis: %v", err)
	}

	return nil
}
