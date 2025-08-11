package utils

import (
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"log"
)

func DeleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	msgToDelete := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Request(msgToDelete)
	if err != nil {
		log.Println("Failed to delete message: ", err)
	}
}

func CreateEditMessage(board tgbotapi.InlineKeyboardMarkup, updatedMessage string, chatID int64, messageID int, parseMode string) (tgbotapi.Chattable, error) {
	var msgChattable tgbotapi.Chattable
	msgChattable = tgbotapi.NewEditMessageText(chatID, messageID, updatedMessage)

	// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
	msgEdit, ok := msgChattable.(tgbotapi.EditMessageTextConfig)
	if !ok {
		return nil, fmt.Errorf("cannot convert msg to type EditMessageTextConfig")
	}

	// Assign the new inline keyboard to the message
	msgEdit.ReplyMarkup = &board
	msgEdit.ParseMode = parseMode

	return msgEdit, nil
}
