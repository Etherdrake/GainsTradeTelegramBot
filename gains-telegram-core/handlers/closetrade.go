package handlers

import (
	"HootCore/api/tradinginteractions"
	"HootCore/database"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"context"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"time"
)

var ctx = context.Background()

func CloseTradeHandler(
	bot *tgbotapi.BotAPI,
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	chatID int64,
	guid int64) {

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return
	}

	pair := pairmaps.IndexToPair[int(trade.ActivePerp.ActiveTradeOrOrder.Trade.PairIndex)]

	currentPrice, err := rdbHoot.GetPrice(int(trade.ActivePerp.ActiveTradeOrOrder.Trade.PairIndex))
	if err != nil {
		log.Println("Market to entry price error", err)
	}

	// Check
	perpCode := trade.ActivePerp.IsTradeOrOrder
	if perpCode != 1 {
		log.Println("Closetrade called but ActivePerpetual is not set as trade")
		return
	}

	// Get the activeCacheTrade
	activeCacheTrade, err := rdbHoot.GetSelectedTrade(guid)
	if err != nil {
		log.Println("Unable to retrieve active trade from cache in BuildActiveTradeBoardGNSV2")
		return
	}

	activeCacheTradeDisplay := activeCacheTrade.IntoDisplay()

	positionSizeConv, _ := strconv.ParseFloat(activeCacheTradeDisplay.Trade.CollateralAmount, 64)
	openPriceConv, _ := strconv.ParseFloat(activeCacheTradeDisplay.Trade.OpenPrice, 64)
	leverageConv, _ := strconv.ParseFloat(activeCacheTradeDisplay.Trade.Leverage, 64)
	getPnlUSD := utils.CalculateDollarProfitOrLossWithLeverage(positionSizeConv, openPriceConv, currentPrice, leverageConv, activeCacheTrade.Trade.Long)

	msgNewStr := fmt.Sprintf("Closing trade: \n*%s* @ *%.2f*, your PNL $*%.2f*", pair, currentPrice, getPnlUSD)
	msgNew := tgbotapi.NewMessage(chatID, msgNewStr)
	msgNew.ParseMode = tgbotapi.ModeMarkdown
	msgFirst, err := bot.Send(msgNew)
	if err != nil {
		log.Println("Error sending add message in /add: ", err)
		return
	}

	_, err = tradinginteractions.CloseTradeGNS(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
	if err != nil {
		log.Println("Error closing trade: ", err)
		return
	}

	var msgChattable tgbotapi.Chattable
	msgSecondText := fmt.Sprintf("Total account credited including collateral *+$%.2f*", float64(trade.InitialPosToken)+getPnlUSD)
	msgChattable = tgbotapi.NewEditMessageText(msgFirst.Chat.ID, msgFirst.MessageID, msgSecondText)
	msgEdit, ok := msgChattable.(tgbotapi.EditMessageTextConfig)
	if !ok {
		log.Println("Cannot convert msg to type EditMessageTextConfig")
		return
	}

	msgEdit.ParseMode = tgbotapi.ModeMarkdown
	msgChattable = msgEdit

	msgTwo, errTg := bot.Send(msgChattable)
	if errTg != nil {
		log.Println("Error sending message: ", err)
	}

	experienceAdd := database.CalculateXP(trade.ActivePerp.ActiveTradeOrOrder.Trade)
	rdbErr := rdbHoot.IncrementExperience(ctx, guid, uint64(experienceAdd))
	if rdbErr != nil {
		log.Println("Error incrementing experience: ", err)
		return
	}

	time.Sleep(3 * time.Second)

	msgThreeText := fmt.Sprintf("🎉 You earned *%.2f HoiSin!* Congrats!", experienceAdd)
	msgChattable = tgbotapi.NewEditMessageText(msgTwo.Chat.ID, msgTwo.MessageID, msgThreeText)
	msgEdit, ok = msgChattable.(tgbotapi.EditMessageTextConfig)
	if !ok {
		log.Println("Cannot convert msg to type EditMessageTextConfig")
		return
	}

	msgEdit.ParseMode = tgbotapi.ModeMarkdown
	msgChattable = msgEdit

	msgThree, errTg := bot.Send(msgChattable)
	if errTg != nil {
		log.Println("Error sending message: ", err)
	}

	//var msgTx tgbotapi.MessageConfig
	//switch trade.Chain {
	//case 137:
	//	msgTxString := fmt.Sprintf("Trade %s closed successfully @ *$%.2f*: [TxReceipt](https://polygonscan.io/tx/%s)", pair, currentPrice, strings.Trim(txReceipt, "\""))
	//	msgTx = tgbotapi.NewMessage(chatID, msgTxString)
	//case 42161:
	//	msgTxString := fmt.Sprintf("Trade %s closed successfully @ *$%.2f*: [TxReceipt](https://arbiscan.io/tx/%s)", pair, currentPrice, strings.Trim(txReceipt, "\""))
	//	msgTx = tgbotapi.NewMessage(chatID, msgTxString)
	//case 421614:
	//	msgTxString := fmt.Sprintf("Trade %s closed successfully @ *$%.2f*: [TxReceipt](https://sepolia.arbiscan.io/tx/%s)", pair, currentPrice, strings.Trim(txReceipt, "\""))
	//	msgTx = tgbotapi.NewMessage(chatID, msgTxString)
	//}
	//msgTx.ParseMode = tgbotapi.ModeMarkdown
	//txMsg, err := bot.Send(msgTx)
	//if err != nil {
	//	log.Printf("Error sending message: %v", err)
	//}

	// TODO: This needs to be changed for Redis
	//strSetErr := tradecachemanagement_legacy.SetStartTradeOrOrder(client, tradeCache, guid, trade.Chain)
	//if strSetErr != nil {
	//	log.Println("Error in SetStartTradeOrOrder inside /closetrade: ", strSetErr)
	//}
	go utils.DeleteMessageWithTimer(bot, chatID, msgThree.MessageID, 5)
	return
}
