package streamprocessor

import (
	"GainsListenerMumbai/logger"
	"GainsListenerMumbai/transformers"
	"GainsListenerMumbai/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

// matchLimitOrders matches events for limit orders
func MatchLimitOrders(
	bot *tgbotapi.BotAPI,
	mongoDb *mongo.Database,
	openLimitChan chan transformers.OpenLimitPlacedEventTransform,
	limitExecutedChan chan transformers.LimitExecutedTransform,
	openLimitCanceledChan chan transformers.OpenLimitCanceledEventTransform,
	nftOrderInitiatedChan chan transformers.NftOrderInitiatedEventTransform,
	nftOrderCanceledChan chan transformers.NftOrderCanceledTransform,
	diamondCollat string,
) {
	fmt.Println("LimitOrderMatcher has been started and channels are live!")

	go func() {
		for {
			select {
			// Bot trigger, can be both a market or limit order
			case nftOrderCanceledEvent, ok := <-nftOrderCanceledChan:
				if !ok {
					fmt.Println("Error")
				}
				go func() {
					err := handleNftOrderCanceled(nftOrderCanceledEvent)
					if err != nil {
						logger.WriteToLog("Error with handler: "+err.Error(), "handlernfterrors")
					}
				}()

			// We don't need to check for this as it is already handled by the orderchecker.
			case openLimitPlacedEvent, ok := <-openLimitChan:
				if !ok {
					fmt.Println("Error")
				}
				//go func() {
				//	err := handleOpenLimitPlaced(*mongoDB, openLimitPlacedEvent, isArb)
				//	if err != nil {
				//		logger.WriteToLog("Error with handler: "+err.Error(), "handlerdb_errors")
				//	}
				//}()

				log.Println("openLimitPlaced encountered which is handled by our other service already: ", openLimitPlacedEvent)

			// LimitExecutedEvent -> the order is now an actual trade as it has been executed, we check for this event.
			case limitExecutedEvent, ok := <-limitExecutedChan:
				if !ok {
					fmt.Println("Error")
				}

				go func() {
					err := handleLimitExecuted(mongoDb, limitExecutedEvent)
					if err != nil {
						logger.WriteToLog("Error handling limit exec: "+err.Error(), "handlelimitexec")
					}
				}()
			// OpenLimitCanceled -> Trading Event -> Write to storage
			case openLimitCanceledEvent, ok := <-openLimitCanceledChan:
				if !ok {
					fmt.Println("Error")
				}

				fmt.Println("Streamprocessor received openLimitCanceledEvent")

				// Handle openLimitCanceledEvent
				// Write to memory

				trader := utils.LowercaseString(openLimitCanceledEvent.Trader)

				logger.WriteToLog("openLimitCanceledEvent: "+
					"Trader: "+trader, "openLimitCanceledEventStream")

			// NftOrderInitiated -> Callback Event -> Write to Storage
			case nftOrderInitiatedEvent, ok := <-nftOrderInitiatedChan:
				if !ok {
					fmt.Println("Error")
				}

				go func() {
					err := handleNftOrderInitiated(nftOrderInitiatedEvent)
					if err != nil {
						log.Println("Error with handleNftOrderInitiated: ", err)
					}
				}()

				trader := utils.LowercaseString(nftOrderInitiatedEvent.Trader)

				orderIdStr := strconv.Itoa(int(nftOrderInitiatedEvent.OrderID))

				fmt.Println("Initiated Received in matcher: ", orderIdStr)

				logger.WriteToLog("nftOrderInitiatedEvent: "+
					"Trader: "+trader, "nftOrderInitiatedEventStream")
			}
		}
	}()
}
