package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUserDB(client *mongo.Client, userID int64) (bool, error) {
	// Access the "users" collection in the "strikerdb" database
	collection := client.Database(databaseName).Collection(collectionName)

	// Check if the user exists in the collection
	filter := bson.M{"_id": userID}
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
