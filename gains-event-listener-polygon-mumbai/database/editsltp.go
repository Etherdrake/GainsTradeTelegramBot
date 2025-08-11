package database

import (
	"GainsListenerMumbai/transformers"
	"context"
	"errors"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// EditTp updates the TP (take profit) value for a given trade in MongoDB.
func EditTp(client *mongo.Client, trader string, pairIndex, index int64, newTp float64) error {
	// First, try to find the trade in the market collection
	tradeMarket, err := FindTradeInMarket(client, trader, pairIndex, index)
	if err != nil {
		// If not found in the market collection, try the limit collection
		tradeLimit, err := FindTradeInLimit(client, trader, pairIndex, index)
		if err != nil {
			return err
		}
		// If found in the limit collection, update TP and return
		if tradeLimit != (transformers.LimitExecutedTransform{}) {
			return updateTpOrSl(client, gnsconstants.DbNameMumbai, gnsconstants.LimitTrades, tradeLimit.Trade.Trader, tradeLimit.Trade.PairIndex, tradeLimit.Trade.Index, newTp, "tp")
		}
		return errors.New("trade not found")
	}

	// If found in the market collection, update TP and return
	return updateTpOrSl(client, gnsconstants.DbNamePoly, gnsconstants.MarketTrades, tradeMarket.Trade.Trader, tradeMarket.Trade.PairIndex, tradeMarket.Trade.Index, newTp, "tp")
}

// EditSl updates the SL (stop loss) value for a given trade in MongoDB.
func EditSl(client *mongo.Client, trader string, pairIndex, index int64, newSl float64) error {
	// First, try to find the trade in the market collection
	tradeMarket, err := FindTradeInMarket(client, trader, pairIndex, index)
	if err != nil {
		// If not found in the market collection, try the limit collection
		tradeLimit, err := FindTradeInLimit(client, trader, pairIndex, index)
		if err != nil {
			return err
		}
		// If found in the limit collection, update SL and return
		if tradeLimit != (transformers.LimitExecutedTransform{}) {
			return updateTpOrSl(client, gnsconstants.DbNameMumbai, gnsconstants.LimitTrades, tradeLimit.Trade.Trader, tradeLimit.Trade.PairIndex, tradeLimit.Trade.Index, newSl, "sl")
		}
		return errors.New("trade not found")
	}

	// If found in the market collection, update SL and return
	return updateTpOrSl(client, gnsconstants.DbNamePoly, gnsconstants.MarketTrades, tradeMarket.Trade.Trader, tradeMarket.Trade.PairIndex, tradeMarket.Trade.Index, newSl, "sl")
}

// Helper function to update TP or SL for a given trade in a specific collection
func updateTpOrSl(client *mongo.Client, dbName, collectionName string, trader string, pairIndex, index int64, value float64, field string) error {
	collection := client.Database(dbName).Collection(collectionName)

	filter := bson.M{
		"trade.trader":     trader,
		"trade.pair_index": pairIndex,
		"trade.index":      index,
	}

	update := bson.M{
		"$set": bson.M{
			"trade." + field: value,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
