package database

import (
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetOpenMarketTradeFromMongo(orderID int64) (*transformers.MarketExecutedTransform, bool, error) {
	mongoDB, err := ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
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

func GetOpenLimitTradeFromMongo(orderID int64) (*transformers.MarketExecutedTransform, bool, error) {
	mongoDB, err := ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.LimitTrades)

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
