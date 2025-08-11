package database

import (
	"GainsListenerV4/types"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TradeExistsByID(ctx context.Context, collection *mongo.Collection, identifier string) (bool, error) {
	filter := bson.M{"_id": identifier}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetTradeByID(ctx context.Context, collection *mongo.Collection, tradeID string) (types.TradeItemMongo, error) {
	filter := bson.M{"_id": tradeID}
	var trade types.TradeItemMongo
	err := collection.FindOne(ctx, filter).Decode(&trade)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return types.TradeItemMongo{}, nil
		}
		return types.TradeItemMongo{}, err
	}
	return trade, nil
}
