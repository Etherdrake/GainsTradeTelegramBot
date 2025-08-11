package streamprocessor

import (
	"GainsListenerMumbai/transformers"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
)

// MatchMarketOrders matches events for market orders
func MatchMarketOrders(
	bot *tgbotapi.BotAPI,
	mongoDb *mongo.Database,
	marketOrderInitiatedChan chan transformers.MarketOrderInitiatedEventTransform,
	marketExecutedChan chan transformers.MarketExecutedTransform,
	marketOpenCanceledChan chan transformers.MarketOpenCanceledTransform,
	marketCloseCanceledChan chan transformers.MarketCloseCanceledTransform,
	diamondCollat string,
) {

	//orderIdToMarketInitiated := make(map[string]transformers.MarketOrderInitiatedEventTransform)

	go func() {
		for {
			select {
			case marketOrderInitiatedEvent, ok := <-marketOrderInitiatedChan:
				fmt.Println("MarketInitiated Received: ", marketOrderInitiatedEvent.OrderID)
				if !ok {
					fmt.Println("Error")
				}

			// marketExecuted is the callback event for the MarketOrderInitiatedEvent
			case marketExecutedEvent, ok := <-marketExecutedChan:
				if !ok {
					fmt.Println("Error")
				}

				go func() {
					// Check if open or close
					err := handleMarketExecuted(mongoDb, marketExecutedEvent)
					if err != nil {
						fmt.Println("Error in MarketExecuted Handler: ", err)
					}
				}()

			// For this it means we write to canceled orders
			case marketOpenCanceledEvent, ok := <-marketOpenCanceledChan:
				if !ok {
					return
				}

				orderID := marketOpenCanceledEvent.OrderID

				fmt.Println("MarketOpenCanceled: ", orderID)

				//if _, ok = orderIdToMarketInitiated[strconv.FormatUint(uint64(orderID), 10)]; ok {
				//	if !ok {
				//		fmt.Println("Not found")
				//	} else {
				//		tgsenders.SendLogTelegram(bot, "Match found between MarketExecuted and MarketOpenCanceled", "AmalgamatedConsolidated")
				//	}
				//}

			// For this it means we write to canceled orders
			case marketCloseCanceledEvent, ok := <-marketCloseCanceledChan:
				if !ok {
					return
				}

				orderID := marketCloseCanceledEvent.OrderID

				fmt.Println("MarketCloseCanceled: ", orderID)
				//if _, ok = orderIdToMarketInitiated[strconv.FormatUint(uint64(orderID), 10)]; ok {
				//	if !ok {
				//		fmt.Println("Not found")
				//	} else {
				//		tgsenders.SendLogTelegram(bot, "Match found between MarketExecuted and marketCloseCanceled", "AmalgamatedConsolidated")
				//	}
				//}
			}
		}
	}()
}
