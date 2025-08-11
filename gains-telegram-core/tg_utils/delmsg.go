package tg_utils

import tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"

// DeleteMessage attempts to delete a message and returns any error encountered
func DeleteMessage(bot *tgbotapi.BotAPI, msgToDelete tgbotapi.DeleteMessageConfig) error {
	_, err := bot.Request(msgToDelete)
	return err
}
