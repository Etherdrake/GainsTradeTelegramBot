package database

import (
	"GainsListenerV4/types"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var ctx = context.Background()

// StoreTradeItem inserts a TradeItemMongo object into MongoDB
func StoreTradeItem(ctx context.Context, collection *mongo.Collection, tradeItem types.TradeItemMongo) error {
	_, err := collection.InsertOne(ctx, tradeItem)
	if err != nil {
		log.Println("Error in StoreTradeItem", err)
		return err
	}
	return nil
}

func StoreTradeItems(collection *mongo.Collection, tradeItems []types.TradeItemMongo) error {
	// Create a slice to store the documents to insert
	var documents []interface{}

	// Iterate over each trade item and convert it to BSON format
	for _, tradeItem := range tradeItems {
		documents = append(documents, tradeItem)
	}

	// Insert all documents into the collection at once
	_, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
