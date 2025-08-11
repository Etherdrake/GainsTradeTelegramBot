package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func GetPublicKey(client *mongo.Client, userID int64) (string, error) {
	// Get the 'users' collection
	users := client.Database("hooterdb").Collection("users")

	// Find the user with the given guid
	var result struct {
		PublicKey string `bson:"public_key"`
	}
	filter := bson.D{{"_id", userID}}
	err := users.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return "", err
	}

	return strings.ToLower(result.PublicKey), nil
}
