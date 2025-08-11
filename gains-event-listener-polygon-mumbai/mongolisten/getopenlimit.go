package mongolisten

import (
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	"fmt"
	gns_constants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetOpenLimitTradeFromMongo(address string, pairIndex int64, index int64, buy bool, leverage int64, collateralToken string) (*transformers.LimitExecutedTransform, bool, error) {
	mongoDB, err := database.ConnectToMongoDB(gns_constants.ConnectionString, gns_constants.DbNameMumbai)
	if err != nil {
		log.Println("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gns_constants.LimitTrades)

	// Define a variable to hold the result.
	var result transformers.LimitExecutedTransform

	query := bson.M{
		"trade.trader":     address,
		"trade.pair_index": pairIndex,
		"trade.index":      index,
		"trade.buy":        buy,
		"trade.leverage":   leverage,
		"collateral_token": collateralToken,
	}
	err = collection.FindOne(context.TODO(), query, options.FindOne()).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("Document not found in GetOpenLimitTradeFromMongo: ", err)
			return nil, false, nil
		}
		log.Println("Error does not equal nil: ", err)
		return nil, false, fmt.Errorf("error finding order: %v", err)
	}
	log.Println("Open LimitTrade query successful, this is the MarketTrade: ", result)
	return &result, true, nil
}
