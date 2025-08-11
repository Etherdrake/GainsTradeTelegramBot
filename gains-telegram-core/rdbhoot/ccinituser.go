package rdbhoot

import (
	"HootCore/altseasoncore"
	"HootCore/database"
	"HootCore/utils"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

func (rdbHoot *HootRedisClient) InitUser(client *mongo.Client, guid int64) error {
	userPubKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		return err
	}

	OpenPrice, err := rdbHoot.GetPrice(0)
	if err != nil {
		return err
	}

	percTP := 800.0
	percSL := 0.0

	TP := utils.CalculateTakeProfit(OpenPrice, 100.0, percTP)
	SL := utils.CalculateStopLoss(OpenPrice, 100.0, percSL)

	Liq := utils.CalculateLiquidationPrice(OpenPrice, 100.0, true)

	activePerp := ActivePerpetual{
		IsTradeOrOrder:     0, // | 0 == none |  1 == trade | 2 == order |
		ActiveTradeOrOrder: altseasoncore.TradeItemInternal{},
		ActiveTradeID:      "",
	}

	// Replace with your actual default values
	coreCache := HootCoreCache{
		ID:              guid,
		Trader:          userPubKey,
		PairIndex:       0,
		Index:           0,
		InitialPosToken: 0,
		PositionSizeDai: 1000,
		OpenPrice:       OpenPrice,
		Buy:             true,
		Leverage:        100,
		TP:              TP,
		SL:              SL,
		Liq:             Liq,
		PercentageTP:    percTP,
		PercentageSL:    percSL,
		OrderType:       0,
		OrderStatus:     0,
		PairPage:        0,
		PnL:             0,
		Chain:           421614, // Arbitrum Sepolia
		ActiveCharts:    false,
		ActiveWindow:    0,
		ActiveOpenTrade: 0,
		ActivePerp:      activePerp,
		Collateral:      1, // 1 = DAI | 2 = WETH | 3 = USDC // TODO: Check so this can not be 0 when ++ or --
	}

	// Marshal the coreCache struct to JSON
	coreCacheJSON, err := json.Marshal(coreCache)
	if err != nil {
		return fmt.Errorf("error marshaling coreCache struct to JSON: %v", err)
	}

	// Store the JSON data in Redis using RedisJSON
	_, err = rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}

	return nil
}

func (rdbHoot *HootRedisClient) ResetTrade(client *mongo.Client, guid int64) error {
	userPubKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		return err
	}

	OpenPrice, err := rdbHoot.GetPrice(0)
	if err != nil {
		return err
	}

	percTP := 800.0
	percSL := 0.0

	TP := utils.CalculateTakeProfit(OpenPrice, 100.0, percTP)
	SL := utils.CalculateStopLoss(OpenPrice, 100.0, percSL)

	Liq := utils.CalculateLiquidationPrice(OpenPrice, 100.0, true)

	activePerp := ActivePerpetual{
		IsTradeOrOrder:     0, // | 0 == none |  1 == trade | 2 == order |
		ActiveTradeOrOrder: altseasoncore.TradeItemInternal{},
		ActiveTradeID:      "",
	}

	// Replace with your actual default values
	coreCache := HootCoreCache{
		ID:              guid,
		Trader:          userPubKey,
		PairIndex:       0,
		Index:           0,
		InitialPosToken: 0,
		PositionSizeDai: 1000,
		OpenPrice:       OpenPrice,
		Buy:             true,
		Leverage:        100,
		TP:              TP,
		SL:              SL,
		Liq:             Liq,
		PercentageTP:    percTP,
		PercentageSL:    percSL,
		OrderType:       0,
		OrderStatus:     0,
		PairPage:        0,
		PnL:             0,
		Chain:           421614, // Arbitrum Sepolia
		ActiveCharts:    false,
		ActiveWindow:    0,
		ActiveOpenTrade: 0,
		ActivePerp:      activePerp,
		Collateral:      1, // 1 = DAI | 2 = WETH | 3 = USDC // TODO: Check so this can not be 0 when ++ or --
	}

	// Marshal the coreCache struct to JSON
	coreCacheJSON, err := json.Marshal(coreCache)
	if err != nil {
		return fmt.Errorf("error marshaling coreCache struct to JSON: %v", err)
	}

	// Store the JSON data in Redis using RedisJSON
	_, err = rdbHoot.JSONSet(ctx, "corecache:"+strconv.FormatInt(guid, 10), ".", coreCacheJSON).Result()
	if err != nil {
		return fmt.Errorf("error storing coreCache struct in Redis: %v", err)
	}

	return nil
}
