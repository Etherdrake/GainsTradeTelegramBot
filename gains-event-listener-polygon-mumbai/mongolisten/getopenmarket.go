package mongolisten

import (
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetOpenMarketTradeFromMongo(address string, pairIndex int64, index int64, buy bool, leverage int64, collateralToken string) (*transformers.MarketExecutedTransform, bool, error) {
	dbName := gnsconstants.DbNameMumbai

	mongoDB, err := database.ConnectToMongoDB(gnsconstants.ConnectionString, dbName)
	if err != nil {
		log.Println("MongoDB could not be initialized")
	}

	collection := mongoDB.Collection(gnsconstants.MarketTrades)

	var result transformers.MarketExecutedTransform

	query := bson.M{
		"trade.trader":           address,
		"trade.pair_index":       pairIndex,
		"trade.index":            index,
		"trade.buy":              buy,
		"trade.leverage":         leverage,
		"trade.collateral_token": collateralToken,
	}

	err = collection.FindOne(context.TODO(), query, options.FindOne()).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("Document not found in GetOpenMarketTradeFromMongo: ", err)
			return nil, false, nil
		}
		log.Println("Error does not equal nil: ", err)
		return nil, false, fmt.Errorf("error finding trade: %v", err)
	}
	log.Println("Open MarketTrade query successful, this is the MarketTrade: ", result)
	return &result, true, nil
}

func GetOpenMarketTradeFromMongoOld(orderID int64) (*transformers.MarketExecutedTransform, bool, error) {
	mongoDB, err := database.ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.MarketTrades)

	// Define a variable to hold the result.
	var result transformers.MarketExecutedTransform

	// Find the document with the specified _id.
	err = collection.FindOne(context.TODO(), bson.M{"_id": orderID}).Decode(&result)
	if err != nil {
		// Check if the error is due to not finding the document.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("error finding order: %v", err)
	}

	return &result, true, nil
}
