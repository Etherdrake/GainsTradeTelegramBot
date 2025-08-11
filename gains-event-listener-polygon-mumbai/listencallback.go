package main

import (
	"GainsListenerMumbai/decoders"
	"GainsListenerMumbai/logger"
	"GainsListenerMumbai/tgsenders"
	"GainsListenerMumbai/topicmaps"
	"GainsListenerMumbai/transformers"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func listenForCallbackEvents(
	client *ethclient.Client,
	bot *tgbotapi.BotAPI,
	rdbPrice *redis.Client,
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

	// Listen for logs
	for {
		select {
		case err := <-sub.Err():
			log.Println("Error receiving sub: ", err)
		case eventLog := <-logs:
			//fmt.Println("[ ============================================================ ]")
			fmt.Println("Received TxHash: ", eventLog.TxHash)

			eventName, ok := topicmaps.TopichashToEventnameCallbacks[eventLog.Topics[0].String()]
			if !ok {
				logger.WriteToLog("Unknown event found, EventHash is: "+eventLog.Topics[0].String(), "unknown_eventhash")
				fmt.Println("TradingCallbacks: Hash not found in TopichashToEventname: \n" + eventLog.Topics[0].String() + "\n[ XXXXXXXXXXXXXXXXXXXXXXXXX ]")
			}
			fmt.Println("This is the Callback eventname: ", eventName)
			//fmt.Println("This is topic0 ", eventLog.Topics[0])
			//fmt.Println("FULL EVENT: ", eventLog.Data)
			switch eventName {
			case "MarketExecuted":
				decodedData, err := decoders.DecodeMarketExecuted(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				logger.WriteToLog(logger.GenerateMarketExecutedLog(eventLog.TxHash.String(), decodedData), "market")
				//logger.WriteToLog(logger.GenerateTransformedMarketExecutedLog(rdbPrice, eventLog.TxHash.String(), transformers.TransformMarketExecuted(eventLog, decodedData, collToken)), "market")
				TransformedEvent := transformers.TransformMarketExecuted(eventLog, decodedData, collToken)
				marketExecutedChan <- TransformedEvent
				tgsenders.SendToMarket(bot, logger.GenerateMarketExecutedLog(eventLog.TxHash.String(), decodedData)+"\n\n"+
					logger.GenerateTransformedMarketExecutedLog(rdbPrice, eventLog.TxHash.String(), transformers.TransformMarketExecuted(eventLog, decodedData, collToken)))
				fmt.Println("TRANSFORMED: ", TransformedEvent)

			case "LimitExecuted":
				decodedData, err := decoders.DecodeLimitExecuted(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				//logger.WriteToLog(logger.GenerateLimitExecutedLog(eventLog.TxHash.String(), decodedData), "limit")
				TransformedEvent := transformers.TransformLimitExecuted(eventLog, decodedData, collToken)

				limitExecutedChan <- TransformedEvent

				logger.WriteToLog(logger.GenerateLimitExecutedTransformLog(rdbPrice, eventLog.TxHash.String(), &TransformedEvent), "limit")
				fmt.Println("TRANSFORMED: ", TransformedEvent)

			case "MarketOpenCanceled":
				decodedData, err := decoders.DecodeMarketOpenCanceled(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				logger.WriteToLog(logger.GenerateMarketOpenCanceledLog(eventLog.TxHash.String(), decodedData), "marketopencanceled")

			case "MarketCloseCanceled":
				decodedData, err := decoders.DecodeMarketCloseCanceled(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				logger.WriteToLog(logger.GenerateMarketCloseCanceledLog(eventLog.TxHash.String(), decodedData), "marketclosecanceled")
			case "NftOrderCanceled":
				decodedData, err := decoders.DecodeNftOrderCanceled(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				logger.WriteToLog(logger.GenerateNftOrderCanceledLog(eventLog.TxHash.String(), decodedData), "marketclosecanceled")
			case "PairMaxLeverageUpdated":
				decodedData, err := decoders.DecodePairMaxLeverageUpdated(&contractAbi, &eventLog)
				if err != nil {
					log.Println(err)
					tgsenders.SendLogTelegram(bot, err.Error(), "HootCallbackErrors")
				}
				logger.WriteToLog(logger.GeneratePairMaxLeverageUpdatedLog(eventLog.TxHash.String(), decodedData), "PairMaxLeverageUpdated")
			default:
				fmt.Println("Callback event not found for some reason: ", topicmaps.TopichashToEventnameCallbacks[eventLog.Topics[0].String()])
			}
		}
	}
}
