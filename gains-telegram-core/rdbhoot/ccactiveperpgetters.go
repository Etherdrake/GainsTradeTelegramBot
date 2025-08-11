package rdbhoot

import (
	"HootCore/altseasoncore"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (rdbHoot *HootRedisClient) GetSelectedTradeID(guid int64) (string, error) {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	return trade.ActivePerp.ActiveTradeOrOrder.ID, nil
}

func (rdbHoot *HootRedisClient) GetIsTrade(guid int64) (uint, error) {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	return trade.ActivePerp.IsTradeOrOrder, nil
}

func (rdbHoot *HootRedisClient) GetSelectedTrade(guid int64) (altseasoncore.TradeItemInternal, error) {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	if trade.ActivePerp.IsTradeOrOrder != 1 {
		return altseasoncore.TradeItemInternal{}, fmt.Errorf("trade is not active or found")
	}

	return trade.ActivePerp.ActiveTradeOrOrder, nil
}

func (rdbHoot *HootRedisClient) GetSelectedOrder(guid int64) (altseasoncore.TradeItemInternal, error) {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	if trade.ActivePerp.IsTradeOrOrder != 2 {
		return altseasoncore.TradeItemInternal{}, fmt.Errorf("order is not active or found")
	}

	return trade.ActivePerp.ActiveTradeOrOrder, nil
}

// GetSelectedTradeOrOrderOrNone check whether there is a trade or order loaded into the ActivePerpCache
func (rdbHoot *HootRedisClient) GetSelectedTradeOrOrderOrNone(guid int64) (bool, uint, altseasoncore.TradeItemInternal, error) {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	// 0 is nothing set, 1 trade, 2 is order
	if trade.ActivePerp.IsTradeOrOrder == 0 {
		return false, trade.ActivePerp.IsTradeOrOrder, altseasoncore.TradeItemInternal{}, nil
	}

	// 0 is nothing set, 1 trade, 2 is order
	if trade.ActivePerp.IsTradeOrOrder == 1 {
		return true, trade.ActivePerp.IsTradeOrOrder, trade.ActivePerp.ActiveTradeOrOrder, nil
	}

	// 0 is nothing set, 1 trade, 2 is order
	if trade.ActivePerp.IsTradeOrOrder == 2 {
		return true, trade.ActivePerp.IsTradeOrOrder, trade.ActivePerp.ActiveTradeOrOrder, nil
	}

	return false, trade.ActivePerp.IsTradeOrOrder, trade.ActivePerp.ActiveTradeOrOrder, fmt.Errorf("end of function and nothing found")
}

func (rdbHoot *HootRedisClient) SetSelectedTradeID(guid int64, TradeID string) error {
	coreCache, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	coreCache.ActivePerp.ActiveTradeID = TradeID

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
