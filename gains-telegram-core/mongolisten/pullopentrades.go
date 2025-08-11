package mongolisten

import (
	"HootCore/altseasoncore"
	"HootCore/altseasontransformers"
	"HootCore/constants"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllOpenTrades(client *mongo.Client, trader string, chain uint64) ([]altseasoncore.TradeItemCore, error) {
	openTrades, err := PullOpenTrades(client, trader, chain)
	if err != nil {
		return nil, err
	}
	openTradesCore := altseasontransformers.MongoArrIntoCoreArr(openTrades)
	return openTradesCore, err
}

// PullOpenTrades retrieves all trades for a given user from MongoDB.
func PullOpenTrades(client *mongo.Client, trader string, chain uint64) ([]altseasoncore.TradeItemMongo, error) {
	var dbName string
	switch chain {
	case 137:
		// Polygon
		dbName = constants.DbNamePoly
	case 42161:
		// Arbitrum
		dbName = constants.DbNameArb
	case 421614:
		// Sepolia
		dbName = constants.DbNameSepolia
	default:
		return nil, errors.New("invalid chain")
	}

	collection := client.Database(dbName).Collection(constants.OpenTrades)

	filter := bson.M{"trade.user": trader}

	// You can customize the options as needed, e.g., sort, limit, etc.
	findOptions := options.Find()

	// Find documents matching the filter
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var openTrades []altseasoncore.TradeItemMongo
	for cursor.Next(context.Background()) {
		var trade altseasoncore.TradeItemMongo
		if err := cursor.Decode(&trade); err != nil {
			return nil, err
		}
		openTrades = append(openTrades, trade)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return openTrades, nil
}
