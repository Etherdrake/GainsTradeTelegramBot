package legacy

import (
	"GainsListenerV4/database"
	"GainsListenerV4/types"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

func ProcessRegisterTrade(mongoDB *mongo.Database, registerTrade types.RegisterTrade) error {
	collectionTrades := mongoDB.Collection("open_trades")
	collectionOrders := mongoDB.Collection("open_orders")

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
		err = database.StoreTradeItem(ctx, collectionTrades, tradeItem)
		if err != nil {
			log.Println("Error Store Trade Item: ", err)
			return err
		}
		// Log the trade item details
		log.Printf("Storage successful! Trade Item ID: %s, Trade Type: %s\n", tradeItem.ID, tradeItem.Trade.TradeType)
	case "1", "2":
		// Store order
		err = database.StoreTradeItem(ctx, collectionOrders, tradeItem)
		if err != nil {
			log.Println("Error Store Trade Item: ", err)
			return err
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
