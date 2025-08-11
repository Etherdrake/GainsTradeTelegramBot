package database

import (
	"GainsListenerV4/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetAllTradeItems(collection *mongo.Collection) ([]types.TradeItemMongo, error) {
	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Create a slice to store the trade items
	var tradeItems []types.TradeItemMongo

	// Iterate over the cursor and decode each document into a TradeItemMongo struct
	for cursor.Next(ctx) {
		var tradeItem types.TradeItemMongo
		if err := cursor.Decode(&tradeItem); err != nil {
			log.Println(err)
			return nil, err
		}
		tradeItems = append(tradeItems, tradeItem)
	}

	// Check if there was an error during cursor iteration
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return tradeItems, nil
}
