package rdbhoot

import (
	"HootCore/hooterrors"
	"HootCore/pairmaps"
	"fmt"
	"strconv"
)

func (rdbHoot *HootRedisClient) SetOrderTypeToMarket(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OrderType = 0

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetOrderTypeToLimit(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OrderType = 1

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetOrderTypeToStop(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OrderType = 2

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementPage(guid int64, class string) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PairPage++

	switch class {
	case "crypto":
		if len(pairmaps.FilteredIndexToCryptoPage(int(coreCacheJSON.PairPage), 10)) == 0 {
			coreCacheJSON.PairPage = 0
		}
	case "forex":
		if len(pairmaps.FilteredIndexToFxPage(int(coreCacheJSON.PairPage), 10)) == 0 {
			coreCacheJSON.PairPage = 0
		}
	}
	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) ResetPage(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PairPage = 0

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementPage(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	if coreCacheJSON.PairPage == 0 {
		return nil
	}

	coreCacheJSON.PairPage--

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementPositionSize(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PositionSizeDai = coreCacheJSON.PositionSizeDai - 100

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetPositionSize(guid int64, positionSize float64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PositionSizeDai = positionSize

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementPositionSize(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PositionSizeDai = coreCacheJSON.PositionSizeDai + 100

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetPairIndex(pairIdx, guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PairIndex = pairIdx

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementPairIndex(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	if coreCacheJSON.PairIndex == 0 {
		return hooterrors.ErrIndexZero
	}

	coreCacheJSON.PairIndex--

	for {
		if !pairmaps.IndexToDelisted[int(coreCacheJSON.PairPage)] {
			break
		}
		coreCacheJSON.PairPage--
	}

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementPairIndex(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.PairIndex++

	for {
		if !pairmaps.IndexToDelisted[int(coreCacheJSON.PairIndex)] {
			break
		}
		coreCacheJSON.PairIndex++
	}
	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}

	return nil
}

func (rdbHoot *HootRedisClient) DecrementLeverage(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Leverage = coreCacheJSON.Leverage - 10

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetLeverage(guid int64, leverage float32) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Leverage = leverage

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementLeverage(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Leverage = coreCacheJSON.Leverage + 10

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) ToggleLongShort(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Buy = !coreCacheJSON.Buy

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetEntryPrice(guid int64, entryPrice float64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OpenPrice = entryPrice

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementEntryPrice(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OpenPrice = coreCacheJSON.OpenPrice * 0.9975

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementEntryPrice(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.OpenPrice = coreCacheJSON.OpenPrice * 1.0025

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}

	return nil
}

func (rdbHoot *HootRedisClient) DecrementStopLoss(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.SL = coreCacheJSON.SL * 0.9975

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementStopLoss(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.SL = coreCacheJSON.SL * 1.0025

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetTakeProfit(guid int64, takeProfit float64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.TP = takeProfit

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) DecrementTakeProfit(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.SL = coreCacheJSON.SL * 0.9975

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) IncrementTakeProfit(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.TP = coreCacheJSON.TP * 1.0025

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetStopLoss(guid int64, stopLoss float64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.SL = stopLoss

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetActiveWindow(guid int64, activeWindow uint8) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.ActiveWindow = activeWindow

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetLiquidation(guid int64, liq float64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Liq = liq

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetChain(guid int64, chain uint64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Chain = chain

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetCollateralToNative(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Collateral = 0

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetCollateralToDai(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Collateral = 1

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetCollateralToWeth(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Collateral = 2

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}

func (rdbHoot *HootRedisClient) SetCollateralToUsdc(guid int64) error {
	coreCacheJSON, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return fmt.Errorf("guid %d does not exist", guid)
	}

	coreCacheJSON.Collateral = 3

	// Store the JSON data in Redis using RedisJSON
	_, err := rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}
	return nil
}
