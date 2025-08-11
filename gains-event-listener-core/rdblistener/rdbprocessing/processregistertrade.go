package rdbprocessing

import (
	"GainsListenerV4/rdblistener"
	"GainsListenerV4/types"
	"context"
	"fmt"
	"log"
	"strings"
)

func ProcessRegisterTrade(rdbHoot rdblistener.HootRedisClient, ctx context.Context, registerTrade types.RegisterTrade) error {
	// Assign the trade
	trade := registerTrade.Value

	// Set the trade item
	trader := strings.ToLower(trade.Trade.User)
	identifier := trader + trade.Trade.Index

	tradeItem := types.TradeItemMongo{
		ID:             identifier,
		Trade:          types.TradeMongo(trade.Trade),
		TradeInfo:      trade.TradeInfo,
		InitialAccFees: trade.InitialAccFees,
	}

	tradeItem.Trade.User = strings.ToLower(tradeItem.Trade.User)

	// We check what kind of order this is
	var err error
	switch tradeItem.Trade.TradeType {
	case "0":
		// Store trade
		// Insert the trade into Redis
		rdbErr := rdbHoot.InsertTrade(tradeItem)
		if rdbErr != nil {
			log.Println("Error in case 0: ", rdbErr)
		}
		// Log the trade item details
		log.Printf("Storage successful! Trade Item ID: %s, Trade Type: %s\n", tradeItem.ID, tradeItem.Trade.TradeType)
	case "1", "2":
		// Insert the order into Redis
		rdbErr := rdbHoot.InsertOrder(tradeItem)
		if rdbErr != nil {
			log.Println("Error in case 0: ", rdbErr)
		}
		log.Printf("Storage successful! Trade Item ID: %s, Trade Type: %s\n", tradeItem.ID, tradeItem.Trade.TradeType)
	default:
		log.Println("Unknown Trade Type")
		err = fmt.Errorf("unknown trade type: %s", tradeItem.Trade.TradeType)
	}

	if err != nil {
		log.Printf("Error storing trade: %v\n", err)
		return err
	}
	return nil
}
