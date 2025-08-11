package database

import (
	"GainsListenerMumbai/api/orders"
	"GainsListenerMumbai/logger"
	"GainsListenerMumbai/utils"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// FindOpenLimitOrderForExecuteMatch finds an open limit order in MongoDB based on trader, index, and pairindex.
func FindOpenLimitOrderForExecuteMatch(trader string, index, pairindex int64) (*orders.OpenOrdersHuman, error) {
	mongoDB, err := ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.OpenOrders)

	filter := bson.M{
		"Trader":    utils.LowercaseString(trader),
		"Index":     index,
		"PairIndex": pairindex,
	}

	var result orders.OpenOrdersHuman // Replace with your actual document type
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Handle case where no matching document is found
			fmt.Println("\n\n\n\n DOCUMENT WAS NOT FOUND FOR LIMITORDERFOREXECUTE: " + trader)
			return nil, fmt.Errorf("no matching document found")
		}
		// Handle other errors
		fmt.Println("error inside FindOpenLimitOrderForExecuteMatch: ", err.Error())
		logger.WriteToLog(err.Error(), "FindOpenLimitOrderForExecuteMatch")
		return nil, err
	}

	return &result, nil
}
