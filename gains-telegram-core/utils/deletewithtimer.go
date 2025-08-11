package utils

import (
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"time"
)

func DeleteMessageWithTimer(bot *tgbotapi.BotAPI, chatID int64, msgID int, seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
	DeleteMessage(bot, chatID, msgID)
}
