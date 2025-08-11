package legacy

import (
	"GainsListenerV4/database"
	"GainsListenerV4/filters"
	"GainsListenerV4/types"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

var ctx = context.Background()

// Op de een of andere manier worden er helemaal geen trades gevonden, ga er stap voor stap bij langs, error lijkt met cursor ergens te zijn.

func ProcessOpenTrades(mongoDB *mongo.Database, streamTrades []types.OpenTrade) {
	log.Println("Processing Open Trades")

	collectionTrades := mongoDB.Collection("open_trades")
	collectionOrders := mongoDB.Collection("open_orders")

	var streamTradesArr []types.TradeItemMongo

	// We loop over all the trades and add them to an array
	for _, trade := range streamTrades {
		trader := strings.ToLower(trade.Trade.User)
		identifier := trader + trade.Trade.Index

		tradeItem := types.TradeItemMongo{
			ID:             identifier,
			Trade:          types.TradeMongo(trade.Trade),
			TradeInfo:      trade.TradeInfo,
			InitialAccFees: trade.InitialAccFees,
		}
		streamTradesArr = append(streamTradesArr, tradeItem)
	}

	var tradesArr []types.TradeItemMongo
	var ordersArr []types.TradeItemMongo

	// We loop over all the trades and divide them into trades and into orders
	for _, tradeOrOrder := range streamTradesArr {
		switch tradeOrOrder.Trade.TradeType {
		case "0":
			// This is but a normal trade sire
			tradesArr = append(tradesArr, tradeOrOrder)
		case "1":
			if tradeOrOrder.Trade.IsOpen {
				// This is an order
				ordersArr = append(ordersArr, tradeOrOrder)
				//log.Println("Order found for identifier: ", tradeOrOrder.ID)
			} else {
				log.Println("TradeType 1 combined with IsOpen = false found.")
			}
		case "2":
			if tradeOrOrder.Trade.IsOpen {
				// This is an order
				ordersArr = append(ordersArr, tradeOrOrder)
				//log.Println("Order found for identifier: ", tradeOrOrder.ID)
			} else {
				log.Println("TradeType 2 combined with IsOpen = false found.")
			}
		}
	}

	// We retrieve all the trades that are already in the database
	allTradesInDatabase, dbErr := database.GetAllTradeItems(collectionTrades)
	if dbErr != nil {
		log.Println(dbErr)
	}

	allOrdersInDatabase, dbErr := database.GetAllOrderItems(collectionOrders)
	if dbErr != nil {
		log.Println(dbErr)
	}

	// We check which trades are in the streamTrades but not in the DB, these are new trades we have missed,
	newTrades, err := filters.FilterOutNonExistingTrades(ctx, collectionTrades, tradesArr)
	if err != nil {
		log.Println(err)
	}

	//log.Println("These are the newTrades: ", newTrades)

	// We check which orders are in the streamTrades but not in the DB, these are new orders we have missed,
	newOrders, err := filters.FilterOutNonExistingOrders(ctx, collectionOrders, ordersArr)
	if err != nil {
		log.Println(err)
	}

	//log.Println("These are the newOrders: ", newOrders)

	// We add these new trades to the database
	if len(newTrades) > 0 {
		storeErr := database.StoreTradeItems(collectionTrades, newTrades)
		if storeErr != nil {
			log.Printf("Error storing trades: %v", err)
		}

		// We retrieve the trades in the database again since we might have since added trades
		allTradesInDatabase, dbErr = database.GetAllTradeItems(collectionTrades)
		if dbErr != nil {
			log.Println(dbErr)
		}
	}

	// We add these new orders to the database
	if len(newOrders) > 0 {
		storeErr := database.StoreTradeItems(collectionOrders, newOrders)
		if storeErr != nil {
			log.Printf("Error storing trades: %v", err)
		}

		// We retrieve the trades in the database again since we might have since added trades
		allOrdersInDatabase, dbErr = database.GetAllTradeItems(collectionOrders)
		if dbErr != nil {
			log.Println(dbErr)
		}
	}

	// We now compare the trades inside the database with the new trades we received
	// and check for trades that are in the database but not in the array.
	orphanTrades := CompareTradeItems(allTradesInDatabase, tradesArr)

	if len(orphanTrades) > 0 {
		// We delete the orphanTrades
		delErr := database.DeleteTradesByIDs(ctx, collectionTrades, orphanTrades)
		if delErr != nil {
			log.Printf("Error deleting orphan trades: %v", delErr)
			return
		}

		log.Println("Orphan Trades Deleted: ", orphanTrades)
	}

	// We now compare the orders inside the database with the new trades we received
	// and check for trades that are in the database but not in the array.
	orphanOrders := CompareTradeItems(allOrdersInDatabase, ordersArr)

	if len(orphanOrders) > 0 {
		// We delete the orphanOrders
		delErr := database.DeleteTradesByIDs(ctx, collectionOrders, orphanOrders)
		if delErr != nil {
			log.Printf("Error deleting orphan orders: %v", delErr)
			return
		}
		log.Println("Orphan Orders Deleted: ", orphanOrders)
	}
	// The process should now be complete
	log.Println("Processing Complete...")
}
