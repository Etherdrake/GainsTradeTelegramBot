package main

import (
	"GainsListenerMumbai/decoders"
	"GainsListenerMumbai/logger"
	"GainsListenerMumbai/tgsenders"
	"GainsListenerMumbai/topicmaps"
	"GainsListenerMumbai/transformers"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func listenForTradingEvents(
	client *ethclient.Client,
	bot *tgbotapi.BotAPI,
	contractAddress common.Address,
	contractAbi abi.ABI,
	eventChannel chan types.Log,
	openLimitChan chan transformers.OpenLimitPlacedEventTransform,
	limitExecutedChan chan transformers.LimitExecutedTransform,
	//nftOrderCanceledChan chan transformers.nft, this type has been removed
	openLimitCanceledChan chan transformers.OpenLimitCanceledEventTransform,
	nftOrderInitiatedChan chan transformers.NftOrderInitiatedEventTransform,
	marketOrderInitiatedChan chan transformers.MarketOrderInitiatedEventTransform,
	marketExecutedChan chan transformers.MarketExecutedTransform,
	marketOpenCanceledChan chan transformers.MarketOpenCanceledTransform,
	marketCloseCanceledChan chan transformers.MarketCloseCanceledTransform,
	TpUpdatedChan chan transformers.TpUpdatedEventTransform,
	SlUpdatedChan chan transformers.SlUpdatedEventTransform,
	collToken string,
) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		ToBlock:   nil,
	}

	// Create a channel to receive logs
	logs := make(chan types.Log)

	// Subscribe to logs
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trading Listener has started!")

	// Listen for logs
	for {
		select {
		case chanErr := <-sub.Err():
			log.Println("Error receiving sub: ", chanErr)
		case eventLog := <-logs:
			//fmt.Println("[ ============================================================ ]")
			//fmt.Println("Received TxHash: ", eventLog.TxHash)
			eventName, ok := topicmaps.TopicHashToEventnameTrading[eventLog.Topics[0].String()]
			if !ok {
				logger.WriteToLog("Unknown event found, EventHash is: "+eventLog.Topics[0].String(), "unknown_eventhash")
				fmt.Println("Trading: Hash not found in TopichashToEventname: \n" + eventLog.Topics[0].String() + "\n[ XXXXXXXXXXXXXXXXXXXXXXXXX ]")
			}

			fmt.Println("Current time: ", time.Now())
			fmt.Println("This is the Trading eventname: ", eventName)
			//fmt.Println("This is topic0 ", eventLog.Topics[0])
			//fmt.Println("FULL EVENT: ", eventLog.Data)

			switch eventName {
			case "Done":
				decodedData, err := decoders.DecodeDoneEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "Done")
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}

			case "Paused":
				decodedData, err := decoders.DecodePausedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}
			case "NumberUpdated":
				decodedData, err := decoders.DecodeNumberUpdatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}

			case "BypassTriggerLinkUpdated":
				decodedData, err := decoders.DecodeBypassTriggerLinkUpdatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					//tgsenders.SendLogTelegram(bot, err.Error(), "BypassTriggerLinkUpdated")
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}

			case "MarketOrderInitiated":
				decodedData, err := decoders.DecodeMarketOrderInitiatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {

					transformedEvent := transformers.TransformMarketOrderInitiatedEvent(*decodedData, collToken)
					marketOrderInitiatedChan <- *transformedEvent
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}

			// We disable this because it is an event that only applies to ORDERS and not TRADES
			case "OpenLimitPlaced":
				decodedData, err := decoders.DecodeOpenLimitPlacedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					transformedEvent := transformers.TransformOpenLimitPlacedEvent(*decodedData, collToken)
					//openLimitChan <- *transformedEvent
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println("Encountering OpenLimitPlaced... skipping: ", transformedEvent)
					//tgsenders.SendLogTelegram(bot, logger.GenerateOpenLimitPlacedEventLog(eventLog.TxHash.String(), decodedData), eventName)
				}

			// We disable this because it is an event that only applies to ORDERS and not TRADES
			case "OpenLimitUpdated":
				decodedData, err := decoders.DecodeOpenLimitUpdatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println("Encountering OpenLimitUpdated... skipping: ", decodedData)
					//tgsenders.SendLogTelegram(bot, logger.GenerateOpenLimitUpdatedEventLog(eventLog.TxHash.String(), decodedData), eventName)
				}

			// We disable this because it is an event that only applies to ORDERS and not TRADES
			case "OpenLimitCanceled":
				decodedData, err := decoders.DecodeOpenLimitCanceledEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					transformedEvent := transformers.TransformOpenLimitCanceledEvent(*decodedData, collToken)
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println("Encountering OpenLimitCanceled... skipping: ", transformedEvent)
					continue
					//openLimitCanceledChan <- *transformedEvent
					//tgsenders.SendLogTelegram(bot, logger.GenerateOpenLimitCanceledEventLog(eventLog.TxHash.String(), decodedData), eventName)
				}

			case "TpUpdated":
				decodedData, err := decoders.DecodeTpUpdatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					transformedData := transformers.TransformTpUpdatedEvent(*decodedData, collToken)
					tgsenders.SendLogTelegram(bot, logger.GenerateTpUpdatedEventLog(eventLog.TxHash.String(), decodedData), eventName)
					TpUpdatedChan <- *transformedData
				}

			case "SlUpdated":
				decodedData, err := decoders.DecodeSlUpdatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					transformedData := transformers.TransformSlUpdatedEvent(*decodedData, collToken)
					SlUpdatedChan <- *transformedData

				}

			case "NftOrderInitiated":
				decodedData, err := decoders.DecodeNftOrderInitiatedEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					transformedEvent := transformers.TransformNftOrderInitiatedEvent(*decodedData, collToken)
					nftOrderInitiatedChan <- *transformedEvent
					// Handle the decoded data, for example, print or use it as needed
				}

			case "ChainlinkCallbackTimeout":
				decodedData, err := decoders.DecodeChainlinkCallbackTimeoutEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}

			case "CouldNotCloseTrade":
				decodedData, err := decoders.DecodeCouldNotCloseTradeEvent(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
				} else {
					// Handle the decoded data, for example, print or use it as needed
					fmt.Println(decodedData)
				}
			default:
				log.Println(bot, "Unknown Event Received!", eventName)
			}

			//// Send the event to the channel
			//eventChannel <- eventLog
		}
	}
}
