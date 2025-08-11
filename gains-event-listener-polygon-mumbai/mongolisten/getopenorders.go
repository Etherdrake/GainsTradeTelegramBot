package mongolisten

import (
	"context"
	"fmt"
	"github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sort"
)

// GetOpenOrders returns an array of OpenLimitOrderConverted based on the provided pubkey and isArb parameters.
func GetOpenOrders(client *mongo.Client, pubkey string, isArb bool) ([]OpenLimitOrderConverted, error) {
	var database string
	if isArb {
		database = gns_constants.DbNameArbi
	} else {
		database = gns_constants.DbNamePoly
	}

	// Specify the MongoDB collection and database
	collection := client.Database(database).Collection("open_orders")

	// Define filter criteria
	filter := bson.M{"trader": pubkey}

	// Fetch documents from MongoDB
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from MongoDB: %v", err)
	}
	defer cur.Close(context.Background())

	// Parse MongoDB documents into OpenLimitOrderBigString structs
	var openOrdersBigString []OpenLimitOrderBigString
	for cur.Next(context.Background()) {
		var order OpenLimitOrderBigString
		err := cur.Decode(&order)
		if err != nil {
			return nil, fmt.Errorf("error decoding MongoDB document: %v", err)
		}
		openOrdersBigString = append(openOrdersBigString, order)
	}

	// Check for errors during cursor iteration
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("error during MongoDB cursor iteration: %v", err)
	}

	// Convert OpenLimitOrderBigString to OpenLimitOrderConverted
	var openOrdersConv []OpenLimitOrderConverted
	for _, bigStringOrder := range openOrdersBigString {
		convertedOrder := ConvertBigLimitToLimit(bigStringOrder)
		openOrdersConv = append(openOrdersConv, convertedOrder)
	}

	// Sort the array by Block in ascending order
	sort.Slice(openOrdersConv, func(i, j int) bool {
		return openOrdersConv[i].Block < openOrdersConv[j].Block
	})

	return openOrdersConv, nil
}
