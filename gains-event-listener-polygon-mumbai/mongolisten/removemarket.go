package mongolisten

import (
	"context"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func RemoveMarketTrade(mongoDB *mongo.Database, address string, pairIndex int64, index int64, buy bool, leverage int64, collateralToken string) error {
	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.MarketTrades)
	res, errMon := collection.DeleteOne(
		context.TODO(),
		bson.M{
			"trade.trader":           address,
			"trade.pair_index":       pairIndex,
			"trade.index":            index,
			"trade.buy":              buy,
			"trade.leverage":         leverage,
			"trade.collateral_token": collateralToken,
		},
	)

	if errMon != nil {
		log.Println("Error", errMon)
	}

	if res.DeletedCount == 0 {
		log.Println("Error", errMon)
		return fmt.Errorf("error deleting market trade: address: %s pairIndex: %d index: %d buy: %t leverage: %d collateral_token: %s", address, pairIndex, index, buy, leverage, collateralToken)
	}

	fmt.Printf("Market Trade removed successfully:\n address: %s pairIndex: %d index: %d buy: %t leverage: %d collateral_token: %s", address, pairIndex, index, buy, leverage, collateralToken)
	return nil
}
