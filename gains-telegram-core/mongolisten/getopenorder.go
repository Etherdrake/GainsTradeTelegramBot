package mongolisten

import (
	"HootCore/altseasoncore"
	"HootCore/constants"
	"HootCore/database"
	"context"
	"errors"
	"fmt"
	"github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

func GetOpenOrder(trader string, index int, chain uint8) (altseasoncore.TradeItemMongo, bool, error) {
	var dbName string
	switch chain {
	case 0:
		// Polygon
		dbName = constants.DbNamePoly
	case 1:
		// Arbitrum
		dbName = constants.DbNameArb
	case 2:
		// Sepolia
		dbName = constants.DbNameSepolia
	default:
		return altseasoncore.TradeItemMongo{}, false, errors.New("invalid chain")
	}

	mongoDB, err := database.ConnectToMongoDB(gns_constants.ConnectionString, dbName)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	// Connect to the "your_database" database and "orders" collection.
	collection := mongoDB.Collection(gns_constants.OpenOrders)

	// Define a variable to hold the result.
	var result altseasoncore.TradeItemMongo

	// Create a filter based on the trader, pairIndex, and index fields.
	filter := bson.M{
		"_id": trader + strconv.Itoa(index),
	}

	// Find the document with the specified filter.
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		// Check if the error is due to not finding the document.
		if errors.Is(err, mongo.ErrNoDocuments) {
			return altseasoncore.TradeItemMongo{}, false, nil
		}
		return altseasoncore.TradeItemMongo{}, false, fmt.Errorf("error finding order: %v", err)
	}
	return result, true, nil
}
