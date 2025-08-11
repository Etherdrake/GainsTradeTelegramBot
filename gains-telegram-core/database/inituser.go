package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitUser(client *mongo.Client, userID int64, firstName, lastName string) error {
	// Access the "users" collection in the "strikerdb" database
	collection := client.Database(databaseName).Collection(collectionName)

	var totalXp int64 = 1

	// Create the document to insert
	document := bson.M{
		"_id":         userID,
		"user_name":   "userName",
		"first_name":  firstName,
		"last_name":   lastName,
		"wallet_set":  false,
		"public_key":  "",
		"private_key": "",
		"total_xp":    totalXp,
	}

	// Insert the document into the collection
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}

//// UserDB represents the user document in the database
//type UserDB struct {
//	ID         primitive.ObjectID `bson:"_id,omitempty"`
//	UserName   string             `bson:"user_name"`
//	FirstName  string             `bson:"first_name"`
//	LastName   string             `bson:"last_name"`
//	WalletSet  bool               `bson:"wallet_set"`
//	PublicKey  string             `bson:"public_key"`
//	PrivateKey string             `bson:"private_key"`
//	TotalXP    int                `bson:"total_xp"`
//}
