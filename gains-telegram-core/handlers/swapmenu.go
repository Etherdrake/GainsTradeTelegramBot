package handlers

import (
	"HootCore/boardbuilders/swapbuilders"
	"HootCore/errorhandling"
	"HootCore/rdbhoot"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func HandleSwap(rdbHoot *rdbhoot.HootRedisClient, bot *tgbotapi.BotAPI, client *mongo.Client, update tgbotapi.Update, guid int64) {
	msgText, builderErr := swapbuilders.BuildSwapStringV1(client, rdbHoot, guid)
	if builderErr != nil {
		log.Println("Error building swap string: ", builderErr)
		return
	}

	var msg tgbotapi.MessageConfig
	if update.Message != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	} else {
		if update.CallbackQuery != nil {
			msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, msgText)
		} else {
			log.Println("Error parsing update")
		}
	}

	board := swapbuilders.BuildSwapBoardV1(rdbHoot, guid)

	// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
	msg.ReplyMarkup = board
	msg.ParseMode = tgbotapi.ModeMarkdown

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("/leverage error: ", err)
		errorhandling.HandleTelegramError(bot, err, update.CallbackQuery.ID)
	}
	return
}
