package handlers

import (
	"HootCore/rdbhoot"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func HandleSubmit(bot *tgbotapi.BotAPI, client *mongo.Client, rdbHoot *rdbhoot.HootRedisClient, guid int64, chatID int64, editMsg tgbotapi.Message, boardMsgId int) error {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist.")
	}

	if trade.OrderType == 0 {
		priceReal, err := rdbHoot.GetPrice(int(trade.PairIndex))
		if err != nil {
			log.Println("Market to entry price error", err)
		}
		err = rdbHoot.SetEntryPrice(guid, priceReal)
		if err != nil {
			log.Println(err)
		}
	}
	_, _, submitTradeErr := SubmitTrade(client, rdbHoot, guid)
	if submitTradeErr != nil {
		errMsg := tgbotapi.NewMessage(chatID, "Error encountered: "+submitTradeErr.Error())
		_, tgErr := bot.Send(errMsg)
		if tgErr != nil {
			log.Printf("Error sending message: %v", submitTradeErr)
		}
		return tgErr
	} else {
		//var msgTx tgbotapi.EditMessageTextConfig
		//switch trade.Chain {
		//case 137:
		//	msgTx = tgbotapi.NewEditMessageText(chatID, editMsg.MessageID, "Trade execution successful: [TxReceipt](https://polygonscan.com/tx/"+strings.Trim(txReceipt, "\"")+")")
		//case 42161:
		//	msgTx = tgbotapi.NewEditMessageText(chatID, editMsg.MessageID, "Trade execution successful: [TxReceipt](https://arbiscan.io/tx/"+strings.Trim(txReceipt, "\"")+")")
		//case 421614:
		//	msgTx = tgbotapi.NewEditMessageText(chatID, editMsg.MessageID, "Trade execution successful: [TxReceipt](https://sepolia.arbiscan.io/tx/"+strings.Trim(txReceipt, "\"")+")")
		//}
		//msgTx.ParseMode = tgbotapi.ModeMarkdown
		//msgTxDel, tgErr := bot.Send(msgTx)
		//if tgErr != nil {
		//	log.Printf("Error sending message: %v", tgErr)
		//}

		//go utils.DeleteMessageWithTimer(bot, chatID, msgTxDel.MessageID, 15)

		// If the trade is openen we move to activeTrades
		err := rdbHoot.SetActiveWindow(guid, 1)
		if err != nil {
			return err
		}
		//if errRemove := redislocker.RemoveUserID(ctx, rdbLocker, guid); errRemove != nil {
		//	log.Printf("Error removing UserID %d from Redis: %v", guid, errRemove)
		//}

		return nil
	}
}
