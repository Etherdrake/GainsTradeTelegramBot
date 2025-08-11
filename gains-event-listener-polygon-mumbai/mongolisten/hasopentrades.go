package mongolisten

import (
	"context"
	"github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// HasOpenTrades checks if there are any documents in MongoDB matching the specified Trader string.
func HasOpenTrades(client *mongo.Client, trader string, isArb bool) (bool, error) {
	hasMarket, err := HasOpenMarketTrades(client, trader, isArb)
	if err != nil {
		return false, err
	}
	hasLimit, err := HasOpenLimitTrades(client, trader, isArb)
	if err != nil {
		return false, err
	}

	if hasMarket || hasLimit {
		return true, nil
	}

	return false, nil
}

// HasOpenMarketTrades checks if there are any documents in MongoDB matching the specified Trader string.
func HasOpenMarketTrades(client *mongo.Client, trader string, isArb bool) (bool, error) {
	var dbName string
	if isArb {
		dbName = gns_constants.DbNameArbi
	} else {
		dbName = gns_constants.DbNamePoly
	}

	collection := client.Database(dbName).Collection(gns_constants.MarketTrades)

	filter := bson.M{"trade.trader": trader}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// HasOpenLimitTrades checks if there are any documents in MongoDB matching the specified Trader string.
func HasOpenLimitTrades(client *mongo.Client, trader string, isArb bool) (bool, error) {
	var dbName string
	if isArb {
		dbName = gns_constants.DbNameArbi
	} else {
		dbName = gns_constants.DbNamePoly
	}

	collection := client.Database(dbName).Collection(gns_constants.LimitTrades)
	filter := bson.M{"trade.trader": trader}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
