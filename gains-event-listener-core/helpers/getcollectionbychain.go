package helpers

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoCollection(mongoDB *mongo.Database, trade, order bool) (*mongo.Collection, error) {
	if trade {
		return mongoDB.Collection("open_trades"), nil
	}
	if order {
		return mongoDB.Collection("open_orders"), nil
	}
	return &mongo.Collection{}, fmt.Errorf("neither trade or order")
}
