package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetXP(hooterdbClient *mongo.Client, userID int64) (int64, error) {
	// Get the 'users' collection
	users := hooterdbClient.Database("hooterdb").Collection("users")

	// Find the user with the given guid
	var result struct {
		TotalXP int64 `bson:"total_xp"`
	}
	filter := bson.D{{"_id", userID}}
	err := users.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.TotalXP, nil
}
