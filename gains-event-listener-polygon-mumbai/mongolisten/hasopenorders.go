package mongolisten

import (
	"context"
	"github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// HasOpenOrders checks if there are any documents in MongoDB matching the specified Trader string.
func HasOpenOrders(client *mongo.Client, trader string, isArb bool) (bool, error) {
	var dbName string
	if isArb {
		dbName = gns_constants.DbNameArbi
	} else {
		dbName = gns_constants.DbNamePoly
	}

	collection := client.Database(dbName).Collection(gns_constants.OpenOrders)

	filter := bson.M{"trader": trader}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
