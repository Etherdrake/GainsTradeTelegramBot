package rdbhoot

import (
	"HootCore/altseasontools"
	"HootCore/database"
	"HootCore/mongolisten"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// SetStartTradeOrOrder sets the first trade or order as active in the cache.
func (rdbHoot *HootRedisClient) SetStartTradeOrOrder(client *mongo.Client, guid int64) error {
	publicKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		fmt.Println("PublicKey Error: ", err)
		return err
	}

	trade, exist := rdbHoot.GetCoreCache(guid)
	if !exist {
		return fmt.Errorf("no guid has been found for guid %d", guid)
	}

	_, openTrade, openOrder, err := mongolisten.HasOpenTradeOrOrder(client, guid, trade.Chain)
	if err != nil {
		log.Println("Error calling HasOpenTradeOrOrder in SetStartTradeOrOrder")
		return err
	}
	if openTrade {
		log.Println("OpenTrade in SetStartTradeOrOrder detected")
		// Set trade as active
		openTrades, mongoErr := mongolisten.GetAllOpenTrades(client, publicKey, trade.Chain)
		if mongoErr != nil {
			log.Println("GetAllTraders error: ", err)
			return mongoErr
		}

		sortedOpenTrades := altseasontools.SortTradesLatestLast(openTrades)
		sortedInternalTrade := sortedOpenTrades[0].IntoInternal()

		err = rdbHoot.SetActiveSelectedTradeOrOrder(guid, sortedInternalTrade)
		if err != nil {
			log.Println("Error setting active order for: ", guid)
			return err
		}
		errSetTp := rdbHoot.SetActivePerpEditTp(guid, sortedInternalTrade.Trade.Tp)
		if errSetTp != nil {
			log.Println("Error setting active sl for: ", guid)
			return err
		}
		errSetSl := rdbHoot.SetActivePerpEditSl(guid, sortedInternalTrade.Trade.Sl)
		if errSetSl != nil {
			log.Println("Error setting active sl for: ", guid)
			return err
		}
		return nil
	}

	// Build only orders at idx0
	if openOrder {
		log.Println("OpenOrder in SetStartTradeOrOrder detected")
		activeOrders, err := mongolisten.GetOpenOrders(client, publicKey, trade.Chain)
		if err != nil {
			log.Println("GetOpenOrders error: ", err)
			return err
		}

		sortedActiveOrders := altseasontools.SortTradesLatestLast(activeOrders)
		sortedInternalOrder := sortedActiveOrders[0].IntoInternal()

		errSetOrder := rdbHoot.SetActiveSelectedOrder(guid, sortedInternalOrder)
		if errSetOrder != nil {
			log.Println("Error setting active order for: ", guid)
			return err
		}
		errSetTp := rdbHoot.SetActivePerpEditTp(guid, sortedInternalOrder.Trade.Tp)
		if errSetTp != nil {
			log.Println("Error setting active sl for: ", guid)
			return err
		}
		errSetSl := rdbHoot.SetActivePerpEditSl(guid, sortedInternalOrder.Trade.Sl)
		if errSetSl != nil {
			log.Println("Error setting active sl for: ", guid)
			return err
		}
	}
	return nil
}
