package rdbhoot

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (rdbHoot *HootRedisClient) GetCoreCache(guid int64) (*HootCoreCache, bool) {
	// Fetch the JSON struct for the given user from Redis
	userCacheJSON, err := rdbHoot.JSONGet(ctx, "corecache:"+strconv.FormatInt(guid, 10)).Result()
	if err != nil {
		log.Printf("error fetching user trade data from Redis: %v", err)
		return nil, false
	}
	// Unmarshal the JSON data into UserTrade struct
	var coreCache HootCoreCache
	err = json.Unmarshal([]byte(userCacheJSON), &coreCache)
	if err != nil {
		log.Printf("error unmarshaling user trade data: %v", err)
		return nil, false
	}
	return &coreCache, true
}

func (rdbHoot *HootRedisClient) GetPairIndex(guid int64) (int64, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}

	return coreCache.PairIndex, nil
}

func (rdbHoot *HootRedisClient) IsLong(guid int64) (bool, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return false, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.Buy, nil
}

func (rdbHoot *HootRedisClient) GetEntryPrice(guid int64) (float64, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.OpenPrice, nil
}

func (rdbHoot *HootRedisClient) GetPositionSize(guid int64) (float64, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.PositionSizeDai, nil
}

func (rdbHoot *HootRedisClient) GetLeverage(guid int64) (float32, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.Leverage, nil
}

func (rdbHoot *HootRedisClient) GetTakeProfit(guid int64) (float64, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.TP, nil
}

func (rdbHoot *HootRedisClient) GetStopLoss(guid int64) (float64, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return 0, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.SL, nil
}

func (rdbHoot *HootRedisClient) GetActiveCharts(guid int64) (bool, error) {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return false, fmt.Errorf("couldn't find core cache for guid %d", guid)
	}
	return coreCache.ActiveCharts, nil
}
