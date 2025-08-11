package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPrivateKey(client *mongo.Client, userID int64, key string) error {
	// Access the "users" collection in the "strikerdb" database
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the update filter to find the user by userID
	filter := bson.M{"_id": userID}

	// Define the update fields
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("wallet_set"):  true,
			fmt.Sprintf("private_key"): key,
		},
	}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func AddPublicKey(client *mongo.Client, userID int64, key string) error {
	// Access the "users" collection in the "strikerdb" database
	collection := client.Database(databaseName).Collection(collectionName)

	// Define the update filter to find the user by userID
	filter := bson.M{"_id": userID}

	// Define the update fields
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("wallet_set"): true,
			fmt.Sprintf("public_key"): key,
		},
	}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
