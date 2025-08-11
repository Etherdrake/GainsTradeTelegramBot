package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetGuid(client *mongo.Client, pubKey string) (int64, error) {
	// Get the 'users' collection
	users := client.Database("hooterdb").Collection("users")

	// Find the user with the given guid
	var result struct {
		Guid int64 `bson:"_id"`
	}
	filter := bson.D{{"public_key", pubKey}}
	err := users.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println("Error users.FindOne: ", err)
		return 0, err
	}

	return result.Guid, nil
}
