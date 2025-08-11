package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// DeleteTradeByID deletes a trade document by its ID.
func DeleteTradeByID(ctx context.Context, collection *mongo.Collection, id string) error {
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting trade by ID:", err)
		return err
	}
	if result.DeletedCount == 0 {
		log.Println("No document found with the provided ID.")
	}
	return nil
}
