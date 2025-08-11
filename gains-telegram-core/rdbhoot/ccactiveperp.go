package rdbhoot

import (
	"HootCore/altseasoncore"
	"fmt"
	"log"
	"strconv"
)

func (rdbHoot *HootRedisClient) SetIsTrade(guid int64, code uint) error {
	if code > 2 {
		log.Println("Wrong code for trade")
		return fmt.Errorf("code too large: %d", code)
	}

	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.IsTradeOrOrder = code

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetActiveSelectedTradeOrOrder(guid int64, activeTrade altseasoncore.TradeItemInternal) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.IsTradeOrOrder = 1
	user.ActivePerp.ActiveTradeOrOrder = activeTrade
	user.ActivePerp.ActiveTradeID = activeTrade.ID

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetActiveSelectedOrder(guid int64, activeOrder altseasoncore.TradeItemInternal) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.IsTradeOrOrder = 2
	user.ActivePerp.ActiveTradeOrOrder = activeOrder
	user.ActivePerp.ActiveTradeID = activeOrder.ID

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetActivePerpEditTp(guid int64, tp float64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	switch user.ActivePerp.IsTradeOrOrder {
	case 0:
		return fmt.Errorf("no trade or order is set")
	case 1:
		user.ActivePerp.ActiveTradeOrOrder.Trade.Tp = tp
	case 2:
		user.ActivePerp.ActiveTradeOrOrder.Trade.Tp = tp
	}

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetActivePerpEditSl(guid int64, sl float64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	switch user.ActivePerp.IsTradeOrOrder {
	case 0:
		return fmt.Errorf("no trade or order is set")
	case 1:
		user.ActivePerp.ActiveTradeOrOrder.Trade.Sl = sl
	case 2:
		user.ActivePerp.ActiveTradeOrOrder.Trade.Sl = sl
	}

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementActivePerpEditTp(guid int64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.ActiveTradeOrOrder.Trade.Tp *= 0.9975

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementActivePerpEditTp(guid int64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.ActiveTradeOrOrder.Trade.Tp *= 1.0025

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementActivePerpEditSl(guid int64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.ActiveTradeOrOrder.Trade.Sl *= 0.9975

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementActivePerpEditSl(guid int64) error {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist in cache: ", guid)
	}

	user.ActivePerp.ActiveTradeOrOrder.Trade.Sl *= 1.0025

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", user).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}
