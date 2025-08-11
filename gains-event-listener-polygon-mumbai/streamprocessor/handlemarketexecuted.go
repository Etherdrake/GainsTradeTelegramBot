package streamprocessor

import (
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/mongolisten"
	"GainsListenerMumbai/transformers"
	"fmt"
	gns_constants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// handleMarketExecuted handles the marketExecuted event, in this case it means a trade has
// been opened or closed through a manual interaction. This can be a marketOrder or limitOrder originally.
func handleMarketExecuted(mongoDb *mongo.Database, event transformers.MarketExecutedTransform) error {
	// Check if OPEN or CLOSE
	if event.Open {
		// Trade is being opened
		err := database.WriteTradeOrOrderToMongoMarket(gns_constants.MarketTrades, event, transformers.MarketOpenCanceledTransform{}, transformers.MarketCloseCanceledTransform{})
		if err != nil {
			log.Println("Error WriteTradeOrOrderToMongoMarket in handleMarketExecuted: ", err)
		}
		return nil
	} else {
		fmt.Println("EVENT IS NOT OPEN")
		// Check first if this was a market order originally
		mktTrade, mktExist, err := mongolisten.GetOpenMarketTradeFromMongo(event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
		if err != nil {
			log.Println("Error GetOpenMarketTradeFromMongo: ", err)
		}
		// This was a trade opened with market that is now being closed:
		if mktExist {
			fmt.Println("MARKET EXISTS")
			// Remove the trade from limit or market
			dbErr := mongolisten.RemoveMarketTrade(mongoDb, event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
			if dbErr != nil {
				log.Println("Error RemoveMarketTrade: ", dbErr)
			}
			fmt.Println("MARKET REMOVED")
			// Write the trade to closed trade
			writeDbErr := database.WriteTradeOrOrderToMongo("closed_trades", mktTrade)
			if writeDbErr != nil {
				log.Println("Error writing to closed_trades")
			}
			return nil
		}

		lmtTrade, lmtExist, err := mongolisten.GetOpenLimitTradeFromMongo(event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
		if err != nil {
			log.Println("Error GetOpenLimitTradeFromMongo: ", err)
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
		errWrite := database.WriteTradeOrOrderToMongoMarket(gns_constants.ClosedTrades, event, transformers.MarketOpenCanceledTransform{}, transformers.MarketCloseCanceledTransform{})
		if errWrite != nil {
			log.Println("Error WriteTradeOrOrderToMongoMarket in handleMarketExecuted: ", err)
		}
		// Remove the trade from market_trades
		err = mongolisten.RemoveMarketTrade(mongoDb, event.Trade.Trader, event.Trade.PairIndex, event.Trade.Index, event.Trade.Buy, event.Trade.Leverage, event.Trade.CollateralToken)
		if err != nil {
			log.Println("Error removing market trading in handleMarketExecuted: ", err)
			return err
		}
		return nil
	}
}
