package mongolisten

import (
	"HootCore/altseasoncore"
	"HootCore/constants"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func GetOpenTradeFromMongo(client *mongo.Client, address string, index int64, chain uint8) (altseasoncore.TradeItemMongo, bool, error) {
	var dbName string
	switch chain {
	case 0:
		// Polygon
		dbName = constants.DbNamePoly
	case 1:
		// Arbitrum
		dbName = constants.DbNameArb
	case 2:
		// Sepolia
		dbName = constants.DbNameSepolia
	default:
		return altseasoncore.TradeItemMongo{}, false, errors.New("invalid chain")
	}

	collection := client.Database(dbName).Collection(constants.OpenTrades)

	var result altseasoncore.TradeItemMongo

	query := bson.M{
		"_id": address + strconv.Itoa(int(index)),
	}

	err := collection.FindOne(context.TODO(), query, options.FindOne()).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("Document not found in GetOpenMarketTradeFromMongo: ", err)
			return altseasoncore.TradeItemMongo{}, false, nil
		}
		log.Println("Error does not equal nil: ", err)
		return altseasoncore.TradeItemMongo{}, false, fmt.Errorf("error finding trade: %v", err)
	}
	log.Println("Open MarketTrade query successful, this is the MarketTrade: ", result)
	return result, true, nil
}
