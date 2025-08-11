package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteTradesByIDs deletes multiple trade documents by their IDs.
func DeleteTradesByIDs(ctx context.Context, collection *mongo.Collection, ids []string) error {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted %d documents.\n", result.DeletedCount)
	return nil
}
