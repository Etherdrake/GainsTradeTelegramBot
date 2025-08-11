package mongolisten

import (
	"HootCore/database"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func HasOpenTradeOrOrder(client *mongo.Client, guid int64, chain uint64) (bool, bool, bool, error) {
	publicKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		log.Println("Error getting pubkey")
		return false, false, false, err
	}

	openTrades, err := GetAllOpenTrades(client, publicKey, chain)
	if err != nil {
		log.Println("GetAllTraders error: ", err)
		return false, false, false, err
	}

	openOrders, err := GetOpenOrders(client, publicKey, chain)
	if err != nil {
		log.Println("GetOpenOrders error: ", err)
		return false, false, false, err
	}

	// We have a trade and an open order
	if len(openTrades) > 0 && len(openOrders) > 0 {
		return true, true, true, nil
	}

	// We have no trades and an open order
	if len(openTrades) == 0 && len(openOrders) > 0 {
		return true, false, true, nil
	}

	// We have a trade and no open order
	if len(openTrades) > 0 && len(openOrders) == 0 {
		return true, true, false, nil
	}

	// We have no trades and no open orders
	return false, false, false, nil
}
