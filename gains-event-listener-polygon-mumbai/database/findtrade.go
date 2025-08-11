package database

import (
	"GainsListenerMumbai/api/orders"
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindTrade looks for a trade with specified parameters in two MongoDB collections.
func FindTrade(client *mongo.Client, trader string, pairIndex, index int64) (transformers.MarketExecutedTransform, transformers.LimitExecutedTransform, error) {
	// First, try to find the trade in the market collection
	tradeMarket, err := FindTradeInMarket(client, trader, pairIndex, index)
	if err != nil {
		return transformers.MarketExecutedTransform{}, transformers.LimitExecutedTransform{}, err
	}

	// If the trade is not found in the market collection, try the limit collection
	if tradeMarket == (transformers.MarketExecutedTransform{}) {
		tradeLimit, err := FindTradeInLimit(client, trader, pairIndex, index)
		if err != nil {
			return transformers.MarketExecutedTransform{}, transformers.LimitExecutedTransform{}, err
		}
		return transformers.MarketExecutedTransform{}, tradeLimit, nil
	}

	return tradeMarket, transformers.LimitExecutedTransform{}, nil
}

// Helper function to find a trade in a specific collection
func FindTradeInMarket(client *mongo.Client, trader string, pairIndex, index int64) (transformers.MarketExecutedTransform, error) {
	collection := client.Database(gnsconstants.DbNameMumbai).Collection(gnsconstants.MarketTrades)

	filter := bson.M{
		"trade.trader":     trader,
		"trade.pair_index": pairIndex,
		"trade.index":      index,
	}

	var result transformers.MarketExecutedTransform
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if errors.Is(mongo.ErrNoDocuments, err) {
		// Trade not found in this collection
		return transformers.MarketExecutedTransform{}, nil
	} else if err != nil {
		return transformers.MarketExecutedTransform{}, err
	}

	return result, nil
}

// Helper function to find a trade in a specific collection
func FindTradeInLimit(client *mongo.Client, trader string, pairIndex, index int64) (transformers.LimitExecutedTransform, error) {
	collection := client.Database(gnsconstants.DbNameMumbai).Collection(gnsconstants.LimitTrades)

	filter := bson.M{
		"trade.trader":     trader,
		"trade.pair_index": pairIndex,
		"trade.index":      index,
	}

	var result transformers.LimitExecutedTransform
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if errors.Is(mongo.ErrNoDocuments, err) {
		// Trade not found in this collection
		return transformers.LimitExecutedTransform{}, nil
	} else if err != nil {
		return transformers.LimitExecutedTransform{}, err
	}

	return result, nil
}

// Helper function to find a trade in a specific collection
func FindTradeInOrder(client *mongo.Client, trader string, pairIndex, index int64) (orders.OpenOrdersHuman, error) {
	collection := client.Database(gnsconstants.DbNameMumbai).Collection(gnsconstants.OpenOrders)

	filter := bson.M{
		"trade.trader":     trader,
		"trade.pair_index": pairIndex,
		"trade.index":      index,
	}

	var result orders.OpenOrdersHuman
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if errors.Is(mongo.ErrNoDocuments, err) {
		// Trade not found in this collection
		return orders.OpenOrdersHuman{}, nil
	} else if err != nil {
		return orders.OpenOrdersHuman{}, err
	}

	return result, nil
}
