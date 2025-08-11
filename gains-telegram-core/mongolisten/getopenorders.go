package mongolisten

import (
	"HootCore/altseasoncore"
	"HootCore/altseasontransformers"
	"HootCore/constants"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOpenOrders returns an array of OpenLimitOrderConverted based on the provided pubkey and isArb parameters.
func GetOpenOrders(client *mongo.Client, pubkey string, chain uint64) ([]altseasoncore.TradeItemCore, error) {
	var dbName string
	switch chain {
	case 137:
		// Polygon
		dbName = constants.DbNamePoly
	case 42161:
		// Arbitrum
		dbName = constants.DbNameArb
	case 421614:
		// Sepolia
		dbName = constants.DbNameSepolia
	default:
		return nil, errors.New("invalid chain")
	}

	// Specify the MongoDB collection and database
	collection := client.Database(dbName).Collection(constants.OpenOrders)

	// Define filter criteria
	filter := bson.M{"trade.user": pubkey}

	// Fetch documents from MongoDB
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from MongoDB: %v", err)
	}
	defer cur.Close(context.Background())

	// Parse MongoDB documents into OpenLimitOrderBigString structs
	var openOrdersRaw []altseasoncore.TradeItemMongo
	for cur.Next(context.Background()) {
		var order altseasoncore.TradeItemMongo
		err := cur.Decode(&order)
		if err != nil {
			return nil, fmt.Errorf("error decoding MongoDB document: %v", err)
		}
		openOrdersRaw = append(openOrdersRaw, order)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error during MongoDB cursor iteration: %v", err)
	}
	return altseasontransformers.MongoArrIntoCoreArr(openOrdersRaw), nil
}
