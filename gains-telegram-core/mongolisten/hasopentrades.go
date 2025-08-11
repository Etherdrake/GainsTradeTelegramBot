package mongolisten

import (
	"HootCore/constants"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// HasOpenMarketTrades checks if there are any documents in MongoDB matching the specified Trader string.
func HasOpenTrades(client *mongo.Client, trader string, chain uint64) (bool, error) {
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
		return false, errors.New("invalid chain")
	}

	collection := client.Database(dbName).Collection(constants.OpenTrades)

	filter := bson.M{"trade.user": trader}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
