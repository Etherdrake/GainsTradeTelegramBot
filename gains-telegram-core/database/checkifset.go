package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfSet(client *mongo.Client, userID int64) (bool, error) {
	// Access the "users" collection in the "strikerdb" database
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the filter to find the user by userID and check if {chain}Set is true
	filter := bson.M{
		"_id":                     userID,
		fmt.Sprintf("wallet_set"): true,
	}

	// Count the matching documents in the collection
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	// If count is greater than 0, it means the {chain}Set is true for the user
	isSet := count > 0

	return isSet, nil
}
