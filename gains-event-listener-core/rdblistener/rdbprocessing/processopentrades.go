package rdbprocessing

import (
	"GainsListenerV4/rdblistener"
	"GainsListenerV4/types"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
)

var mu sync.Mutex

//
//func ProcessOpenTrades(rdbHoot *rdblistener.HootRedisClient, ctx context.Context, streamTrades []types.OpenTrade) {
//	log.Println("Processing Open Trades")
//
//	var streamTradesArr []types.TradeItemMongo
//
//	// We loop over all the trades and add them to an array
//	for _, trade := range streamTrades {
//		trader := strings.ToLower(trade.Trade.User)
//		identifier := trader + trade.Trade.Index
//
//		tradeItem := types.TradeItemMongo{
//			ID:             identifier,
//			Trade:          types.TradeMongo(trade.Trade),
//			TradeInfo:      trade.TradeInfo,
//			InitialAccFees: trade.InitialAccFees,
//		}
//		streamTradesArr = append(streamTradesArr, tradeItem)
//	}
//
//	tradeKeysArr := make(map[string]bool)
//	orderKeysArr := make(map[string]bool)
//
//	// We loop over all the trades and divide them into trades and into orders
//	for _, tradeOrOrder := range streamTradesArr {
//		switch tradeOrOrder.Trade.TradeType {
//		case "0":
//			// Check if the trade exists
//			exists, err := rdbHoot.CheckIfTradeExists(ctx, tradeOrOrder.ID, strings.ToLower(tradeOrOrder.Trade.User))
//			if err != nil {
//				log.Println(err)
//			}
//			switch exists {
//			case true:
//				// Update the trade if it exists
//				rdbErr := rdbHoot.UpdateTrade(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in case 0: ", rdbErr)
//				}
//			case false:
//				// Insert the trade into Redis
//				rdbErr := rdbHoot.InsertTrade(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in case 0: ", rdbErr)
//				}
//			}
//			tradeKeysArr[fmt.Sprintf("trades:%s:%s", tradeOrOrder.Trade.User, tradeOrOrder.ID)] = true
//		case "1":
//			// Check if the order exists
//			exists, err := rdbHoot.CheckIfOrderExists(ctx, tradeOrOrder.ID, strings.ToLower(tradeOrOrder.Trade.User))
//			if err != nil {
//				log.Println(err)
//			}
//			switch exists {
//			case true:
//				// Update the order if it exists
//				rdbErr := rdbHoot.UpdateOrder(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in true case 1: ", rdbErr)
//				}
//			case false:
//				// Insert the order into Redis
//				rdbErr := rdbHoot.InsertOrder(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in false case 1: ", rdbErr)
//				}
//			}
//			orderKeysArr[fmt.Sprintf("orders:%s:%s", tradeOrOrder.Trade.User, tradeOrOrder.ID)] = true
//		case "2":
//			// Check if the order exists
//			exists, err := rdbHoot.CheckIfOrderExists(ctx, tradeOrOrder.ID, strings.ToLower(tradeOrOrder.Trade.User))
//			if err != nil {
//				log.Println(err)
//			}
//			switch exists {
//			case true:
//				// Update the order if it exists
//				rdbErr := rdbHoot.UpdateOrder(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in case 0: ", rdbErr)
//				}
//			case false:
//				// Insert the order into Redis
//				rdbErr := rdbHoot.InsertOrder(tradeOrOrder)
//				if rdbErr != nil {
//					log.Println("Error in case 0: ", rdbErr)
//				}
//			}
//			orderKeysArr[fmt.Sprintf("orders:%s:%s", tradeOrOrder.Trade.User, tradeOrOrder.ID)] = true
//		}
//	}
//	err := rdbHoot.DeleteTradeKeysNotInMap(ctx, tradeKeysArr)
//	if err != nil {
//		log.Println("Error in DeleteTradeKeysNotInArray: ", err)
//	}
//
//	err = rdbHoot.DeleteOrderKeysNotInMap(ctx, orderKeysArr)
//	if err != nil {
//		log.Println("Error in DeleteOrderKeysNotInArray: ", err)
//	}
//
//	// The process should now be complete
//	log.Println("Processing Complete...")
//}

func ProcessOpenTrades(rdbHoot rdblistener.HootRedisClient, ctx context.Context, streamTrades []types.OpenTrade) {
	log.Println("Processing Open Trades")

	var streamTradesArr []types.TradeItemMongo

	// Loop over all the trades and add them to an array
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

	tradeKeysArr := make(map[string]bool)
	orderKeysArr := make(map[string]bool)

	tradeKeysCh := make(chan string)
	orderKeysCh := make(chan string)

	var wg sync.WaitGroup
	var keysWg sync.WaitGroup

	// Process trades and orders concurrently using goroutines
	for _, tradeOrOrder := range streamTradesArr {
		wg.Add(1)
		go func(tradeOrOrder types.TradeItemMongo) {
			defer wg.Done()
			processTradeOrOrder(ctx, rdbHoot, tradeOrOrder, tradeKeysCh, orderKeysCh)
		}(tradeOrOrder)
	}

	// Collect trade keys
	keysWg.Add(1)
	go func() {
		defer keysWg.Done()
		for tradeKey := range tradeKeysCh {
			tradeKeysArr[tradeKey] = true
		}
	}()

	// Collect order keys
	keysWg.Add(1)
	go func() {
		defer keysWg.Done()
		for orderKey := range orderKeysCh {
			orderKeysArr[orderKey] = true
		}
	}()

	// Wait for all processTradeOrOrder goroutines to finish
	wg.Wait()

	// Close channels after all goroutines have finished
	close(tradeKeysCh)
	close(orderKeysCh)

	// Wait for key collection goroutines to finish
	keysWg.Wait()

	log.Println("Processed Open Trades")

	// Delete keys not in the map
	err := rdbHoot.DeleteTradeKeysNotInMap(ctx, tradeKeysArr)
	if err != nil {
		log.Println("Error in DeleteTradeKeysNotInArray: ", err)
	}

	err = rdbHoot.DeleteOrderKeysNotInMap(ctx, orderKeysArr)
	if err != nil {
		log.Println("Error in DeleteOrderKeysNotInArray: ", err)
	}

	// The process should now be complete
	log.Println("Deletions complete...")
}

func processTradeOrOrder(ctx context.Context, rdbHoot rdblistener.HootRedisClient, tradeOrOrder types.TradeItemMongo, tradeKeysCh chan<- string, orderKeysCh chan<- string) {
	switch tradeOrOrder.Trade.TradeType {
	case "0":
		// Check if the trade exists
		exists, err := rdbHoot.CheckIfTradeExists(ctx, tradeOrOrder.ID, strings.ToLower(tradeOrOrder.Trade.User))
		if err != nil {
			log.Println(err)
		}

		if exists {
			// Update the trade if it exists
			rdbErr := rdbHoot.UpdateTrade(tradeOrOrder)
			if rdbErr != nil {
				log.Println("Error in case 0: ", rdbErr)
			}
		} else {
			// Insert the trade into Redis
			rdbErr := rdbHoot.InsertTrade(tradeOrOrder)
			if rdbErr != nil {
				log.Println("Error in case 0: ", rdbErr)
			}
		}
		tradeKeysCh <- fmt.Sprintf("trades:%s:%s", tradeOrOrder.Trade.User, tradeOrOrder.ID)

	case "1", "2":
		// Check if the order exists
		exists, err := rdbHoot.CheckIfOrderExists(ctx, tradeOrOrder.ID, strings.ToLower(tradeOrOrder.Trade.User))
		if err != nil {
			log.Println(err)
		}

		if exists {
			// Update the order if it exists
			rdbErr := rdbHoot.UpdateOrder(tradeOrOrder)
			if rdbErr != nil {
				log.Println("Error in true case 1: ", rdbErr)
			}
		} else {
			// Insert the order into Redis
			rdbErr := rdbHoot.InsertOrder(tradeOrOrder)
			if rdbErr != nil {
				log.Println("Error in false case 1: ", rdbErr)
			}
		}
		orderKeysCh <- fmt.Sprintf("orders:%s:%s", tradeOrOrder.Trade.User, tradeOrOrder.ID)
	}
}
