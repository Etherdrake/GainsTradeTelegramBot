package mongolisten

import (
	"GainsListenerMumbai/transformers"
	"context"
	"github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTrades(client *mongo.Client, trader string, isArb bool) ([]transformers.MarketExecutedTransform, []transformers.LimitExecutedTransform, error) {
	marketTrades, err := PullOpenMarketTrades(client, trader, isArb)
	if err != nil {
		return nil, nil, err
	}
	limitTrades, err := PullOpenLimitTrades(client, trader, isArb)
	if err != nil {
		return nil, nil, err
	}

	return marketTrades, limitTrades, nil
}

// PullOpenMarketTrades retrieves all trades for a given user from MongoDB.
func PullOpenMarketTrades(client *mongo.Client, trader string, isArb bool) ([]transformers.MarketExecutedTransform, error) {
	var dbName string
	if isArb {
		dbName = gns_constants.DbNameArbi
	} else {
		dbName = gns_constants.DbNamePoly
	}
	collection := client.Database(dbName).Collection(gns_constants.MarketTrades)

	filter := bson.M{"trade.trader": trader}

	// You can customize the options as needed, e.g., sort, limit, etc.
	findOptions := options.Find()

	// Find documents matching the filter
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var marketTrades []transformers.MarketExecutedTransform
	for cursor.Next(context.Background()) {
		var trade transformers.MarketExecutedTransform
		if err := cursor.Decode(&trade); err != nil {
			return nil, err
		}
		marketTrades = append(marketTrades, trade)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return marketTrades, nil
}

// PullOpenTrades retrieves all trades for a given user from MongoDB.
func PullOpenLimitTrades(client *mongo.Client, trader string, isArb bool) ([]transformers.LimitExecutedTransform, error) {
	var dbName string
	if isArb {
		dbName = gns_constants.DbNameArbi
	} else {
		dbName = gns_constants.DbNamePoly
	}

	collection := client.Database(dbName).Collection(gns_constants.LimitTrades)

	filter := bson.M{"Trader": trader}

	// You can customize the options as needed, e.g., sort, limit, etc.
	findOptions := options.Find()

	// Find documents matching the filter
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var limitTrades []transformers.LimitExecutedTransform
	for cursor.Next(context.Background()) {
		var trade transformers.LimitExecutedTransform
		if err := cursor.Decode(&trade); err != nil {
			return nil, err
		}
		limitTrades = append(limitTrades, trade)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return limitTrades, nil
}
