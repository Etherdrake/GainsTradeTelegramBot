package mongolisten

import (
	"GainsListenerMumbai/database"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func HasOpenTradeOrOrder(client *mongo.Client, guid int64, isArb bool) (bool, bool, bool, error) {
	publicKey, err := database.GetPublicKey(guid)
	if err != nil {
		log.Println("Error getting pubkey")
		return false, false, false, err
	}

	marketTrades, limitTrades, err := GetAllTrades(client, publicKey, isArb)
	if err != nil {
		log.Println("GetAllTraders error: ", err)
		return false, false, false, err
	}

	activeOrders, err := GetOpenOrders(client, publicKey, isArb)
	if err != nil {
		log.Println("GetOpenOrders error: ", err)
		return false, false, false, err
	}

	// We have a trade and an open order
	if ((len(marketTrades) > 0) || (len(limitTrades) > 0)) && len(activeOrders) > 0 {
		return true, true, true, nil
	}

	// We have no trades and an open order
	if ((len(marketTrades) == 0) || (len(limitTrades) == 0)) && len(activeOrders) > 0 {
		return true, false, true, nil
	}

	// We have a trade and no open order
	if ((len(marketTrades) > 0) || (len(limitTrades) > 0)) && len(activeOrders) == 0 {
		return true, true, false, nil
	}

	// We have no trades and no open orders
	return false, false, false, nil
}
