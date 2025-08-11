package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetIdWithPublicKey(ctx context.Context, collection *mongo.Collection, pubKey string) (float64, error) {
	var result struct {
		ID float64 `bson:"_id"`
	}

	filter := bson.M{"public_key": pubKey}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}
