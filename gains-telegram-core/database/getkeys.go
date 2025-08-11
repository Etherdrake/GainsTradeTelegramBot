package database

import (
	"HootCore/types"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetKeys(client *mongo.Client, userID int64) (*types.Keys, error) {
	// Get the 'users' collection
	users := client.Database("hooterdb").Collection("users")

	// Find the user with the given guid
	var result struct {
		PublicKey  string `bson:"public_key"`
		PrivateKey string `bson:"private_key"`
	}
	filter := bson.D{{"_id", userID}}
	err := users.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &types.Keys{PublicKey: result.PublicKey, PrivateKey: result.PrivateKey}, nil
}
