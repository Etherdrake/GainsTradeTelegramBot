package mongolisten

import (
	"GainsListenerMumbai/database"
	"context"
	"errors"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetOpenOrder(trader string, pairIndex, index int) (*OpenLimitOrderConverted, bool, error) {
	mongoDB, err := database.ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gnsconstants.OpenOrders)

	// Define a variable to hold the result.
	var result OpenLimitOrderConverted

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
