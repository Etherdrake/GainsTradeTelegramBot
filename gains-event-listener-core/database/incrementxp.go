package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func IncrementUserXP(ctx context.Context, collection *mongo.Collection, userID int64, increment int64) error {
	var user bson.M

	filter := bson.M{"_id": userID}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return err
	}

	if user["total_xp"] == nil {
		return fmt.Errorf("user with ID %d does not have a 'total_xp' field", userID)
	}

	totalXP, ok := user["total_xp"].(int64)
	if !ok {
		return fmt.Errorf("user with ID %d has a 'total_xp' field with an invalid type", userID)
	}

	newTotalXP := totalXP + increment
	update := bson.M{"$set": bson.M{"total_xp": newTotalXP}}

	log.Println("UserID: ", userID)

	updateOpts := options.Update().SetUpsert(false)

	result, err := collection.UpdateOne(ctx, filter, update, updateOpts)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
