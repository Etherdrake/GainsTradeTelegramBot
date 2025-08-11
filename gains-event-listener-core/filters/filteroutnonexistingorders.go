package filters

import (
	"GainsListenerV4/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// FilterOutNonExistingOrders filters out trades that are not present in the database and returns them as TradeItemMongo
func FilterOutNonExistingOrders(ctx context.Context, collection *mongo.Collection, streamTradesArr []types.TradeItemMongo) ([]types.TradeItemMongo, error) {
	var existingTradeIDs []string

	// Collect the IDs of all trades in the array
	for _, trade := range streamTradesArr {
		existingTradeIDs = append(existingTradeIDs, trade.ID)
	}

	// Query to find trades that exist in the database
	filter := bson.M{"_id": bson.M{"$in": existingTradeIDs}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Error finding trades in DB:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Create a map of existing trade IDs
	existingTradeMap := make(map[string]struct{})
	for cursor.Next(ctx) {
		var trade types.TradeItemMongo
		if err := cursor.Decode(&trade); err != nil {
			log.Println("Error decoding trade from DB:", err)
			return nil, err
		}
		existingTradeMap[trade.ID] = struct{}{}
	}

	// Filter out trades that are not in the existingTradeMap
	var missingTrades []types.TradeItemMongo
	for _, trade := range streamTradesArr {
		if _, exists := existingTradeMap[trade.ID]; !exists {
			missingTrades = append(missingTrades, trade)
		}
	}

	// Return the missing trades
	return missingTrades, nil
}
