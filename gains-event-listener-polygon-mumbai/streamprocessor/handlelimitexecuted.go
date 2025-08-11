package streamprocessor

import (
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/mongolisten"
	"GainsListenerMumbai/transformers"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// handleLimitExecuted needs to match with OpenLimitPlaced and NftOrderInitiated
// the LimitExecuted event can be either a marketOrder that is being closed because a LIMIT / SL / TP
// was hit or a limitOrder that is being opened or closed.
func handleLimitExecuted(mongoDb *mongo.Database, event transformers.LimitExecutedTransform) error {
	var isBuy bool
	if event.OrderType == 0 || event.OrderType == 1 || event.OrderType == 2 {
		isBuy = false
	} else {
		isBuy = true
	}

	// if the trade already exists it must be a SELL
	mktTrade, mktExist, err := mongolisten.GetOpenMarketTradeFromMongo(event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.CollateralToken)
	if err != nil {
		log.Println("Error GetOpenMarketTradeFromMongo: ", err)
	}

	lmtTrade, lmtExist, err := mongolisten.GetOpenLimitTradeFromMongo(event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.CollateralToken)
	if err != nil {
		log.Println("Error GetOpenLimitTradeFromMongo: ", err)
	}

	fmt.Printf("MarketTrade: %t LimitTrade: %t\n", mktExist, lmtExist)

	// Check if the event is a BUY or SELL
	switch isBuy {
	// An order has been triggered and opened through Limit being hit and is therefore always a Limit trade
	case true:
		dbErr := database.WriteTradeOrOrderToMongo("limit_trades", event)
		if dbErr != nil {
			log.Println("Error WriteTradeOrOrderToMongo: ", dbErr)
		}
		log.Println("BUY ENCOUNTERED: ORDERTYPE -> ", event.OrderType)
	// A trade has been triggered through a LIQ / TP / SL and NOT through a manual close
	case false:
		log.Println("SELL ENCOUNTERED: ORDERTYPE -> ", event.OrderType)
		// This was a trade opened with market that is now being closed:
		if mktExist {
			// Remove the trade from limit or market
			dbErr := mongolisten.RemoveMarketTrade(mongoDb, event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
			if dbErr != nil {
				log.Println("Error RemoveMarketTrade: ", dbErr)
			}
			// Write the trade to closed trade
			writeDbErr := database.WriteTradeOrOrderToMongo("closed_trades", mktTrade)
			if writeDbErr != nil {
				log.Println("Error writing to closed_trades")
			}
			return nil
		}
		// This was a trade opened through a limit that is now being closed
		if lmtExist {
			dbErr := mongolisten.RemoveLimitTrade(mongoDb, event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
			if dbErr != nil {
				log.Println("Error RemoveMarketTrade: ", dbErr)
			}
			// Write the trade to closed trade
			writeDbErr := database.WriteTradeOrOrderToMongo("closed_trades", lmtTrade)
			if writeDbErr != nil {
				log.Println("Error writing to closed_trades")
			}
			return nil
		}
		return nil
	}
	return nil
}
