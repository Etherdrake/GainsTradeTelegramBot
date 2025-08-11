package handlers

import (
	"HootCore/boardbuilders"
	"HootCore/boardbuilders/newtradebuilders"
	"HootCore/boardbuilders/stringbuilders"
	"HootCore/mongolisten"
	"HootCore/rdbhoot"
	"HootCore/utils"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func UpdateNewTradeOrActiveTrades(
	bot *tgbotapi.BotAPI,
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	boardMsgChatId int64, boardMsgMsgId int,
	guid int64) error {

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	// Let's first check if we are on a chart window or not
	switch trade.ActiveCharts {
	case true:
		board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
		// If true, we only edit the message for message with caption
		editMessage := tgbotapi.NewEditMessageReplyMarkup(boardMsgChatId, boardMsgMsgId, board)
		_, err := bot.Send(editMessage)
		if err != nil {
			log.Println("Error sending message over Telegram: ", err)
			return err
		}
		return nil
	case false:
		break
	}
	switch trade.ActiveWindow {
	// NewTrade Window
	case 0:
		board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
		updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
		finalMsg, err := utils.CreateEditMessage(board, updatedMessage, boardMsgChatId, boardMsgMsgId, "Markdown")
		if err != nil {
			log.Println("Error CreatingEditMessage message: ", err)
			return err
		}
		_, err = bot.Send(finalMsg)
		if err != nil {
			log.Println("Error sending message: ", err)
			return err
		}
	// ActiveTrades Window
	case 1:
		var finalMsg tgbotapi.Chattable
		hasTradeOrOrder, hasTrades, hasOrder, err := mongolisten.HasOpenTradeOrOrder(client, guid, trade.Chain)
		if err != nil {
			log.Println("Error calling HasOpenTradeOrOrder")
		}
		if hasTradeOrOrder {
			// We have no orders and no trades
			if !hasOrder && !hasTrades {
				// Build the string before the board to update the cache first
				updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
				board := boardbuilders.BuildActiveTradeBoardGNSV2(client, rdbHoot, guid, trade.Chain)
				finalMsg, err = utils.CreateEditMessage(board, updatedMessage, boardMsgChatId, boardMsgMsgId, "Markdown")
				if err != nil {
					log.Println("Error in CreateEditMessage in case polygon")
				}
			}

			// There is an order, but we do not have any trades
			if hasOrder || hasTrades {
				// Build the string before the board to update the cache first
				updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
				board := boardbuilders.BuildActiveTradeBoardGNSV2(client, rdbHoot, guid, trade.Chain)
				finalMsg, err = utils.CreateEditMessage(board, updatedMessage, boardMsgChatId, boardMsgMsgId, "Markdown")
				if err != nil {
					log.Println("Error in CreateEditMessage in case polygon")
				}
			}
			_, err = bot.Send(finalMsg)
			if err != nil {
				log.Println("Error sending message: ", err)
				return err
			}
		}
	}
	return nil
}
