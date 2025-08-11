package database

import (
	"GainsListenerMumbai/api/orders"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetOpenOrderFromMongoKsuid(ksuid string, isArb bool) (*orders.OpenOrdersHuman, bool, error) {
	var dbName string
	if isArb {
		dbName = gnsconstants.DbNameArbi
	} else {
		dbName = gnsconstants.DbNamePoly
	}

	mongoDB, err := ConnectToMongoDB(gnsconstants.ConnectionString, dbName)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.OpenOrders)

	// Define a variable to hold the result.
	var result orders.OpenOrdersHuman

	// Find the document with the specified _id.
	err = collection.FindOne(context.TODO(), bson.M{"_id": ksuid}).Decode(&result)
	if err != nil {
		// Check if the error is due to not finding the document.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("error finding order: %v", err)
	}

	return &result, true, nil
}

func GetOpenOrderFromMongo(trader string, pairIndex, index int, isArb bool) (*orders.OpenOrdersHuman, bool, error) {
	var dbName string
	if isArb {
		dbName = gnsconstants.DbNameArbi
	} else {
		dbName = gnsconstants.DbNamePoly
	}

	mongoDB, err := ConnectToMongoDB(gnsconstants.ConnectionString, dbName)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.OpenOrders)

	// Define a variable to hold the result.
	var result orders.OpenOrdersHuman

	// Create a filter based on the trader, pairIndex, and index fields.
	filter := bson.M{
		"trader":    trader,
		"pairIndex": pairIndex,
		"index":     index,
	}

	// Find the document with the specified filter.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// Check if the error is due to not finding the document.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("error finding order: %v", err)
	}

	return &result, true, nil
}
