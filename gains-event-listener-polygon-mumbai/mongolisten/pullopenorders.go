package mongolisten

//// GetOpenOrders retrieves all open orders for a given user from MongoDB.
//func GetOpenOrders(client *mongo.Client, trader string, isArb bool) ([]orders.OpenOrdersHuman, error) {
//	var dbName string
//	if isArb {
//		dbName = constants.DbNameArb
//	} else {
//		dbName = constants.DbNamePoly
//	}
//
//	collection := client.Database(dbName).Collection(constants.OpenOrders)
//
//	filter := bson.M{"trader": trader}
//
//	// You can customize the options as needed, e.g., sort, limit, etc.
//	findOptions := options.Find()
//
//	// Find documents matching the filter
//	cursor, err := collection.Find(context.Background(), filter, findOptions)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(context.Background())
//
//	var openOrders []orders.OpenOrdersHuman
//	for cursor.Next(context.Background()) {
//		var order orders.OpenOrdersHuman
//		if err := cursor.Decode(&order); err != nil {
//			return nil, err
//		}
//		openOrders = append(openOrders, order)
//	}
//
//	if err := cursor.Err(); err != nil {
//		return nil, err
//	}
//
//	return openOrders, nil
//}
