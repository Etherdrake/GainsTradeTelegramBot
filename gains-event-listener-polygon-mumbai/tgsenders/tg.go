package tgsenders

import (
	"GainsListenerMumbai/logger"
	gnsconstants "github.com/Etherdrake/gns-constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendToMarket(bot *tgbotapi.BotAPI, text string) {
	replyMsg := tgbotapi.NewMessage(gnsconstants.HootCommandControlID, "```\n"+text+"\n```")
	replyMsg.ReplyToMessageID = gnsconstants.HootMarketExecutedID
	replyMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(replyMsg)
	if err != nil {
		logger.WriteToLog(err.Error(), "sendlog_errors")
	}
}

func SendToLimit(bot *tgbotapi.BotAPI, text string) {
	replyMsg := tgbotapi.NewMessage(gnsconstants.HootCommandControlID, "```\n"+text+"\n```")
	replyMsg.ReplyToMessageID = gnsconstants.HootLimitExecuted
	replyMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(replyMsg)
	if err != nil {
		logger.WriteToLog(err.Error(), "sendlog_errors")
	}
}

// SendLogTelegram sends to the dedicated listener channel
func SendLogTelegram(bot *tgbotapi.BotAPI, text string, eventName string) {
	replyMsg := tgbotapi.NewMessage(gnsconstants.HootCommandControlID, "```\n"+text+"\n```")
	replyMsg.ReplyToMessageID = EventnameToTelegramChannel[eventName]
	replyMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(replyMsg)
	if err != nil {
		logger.WriteToLog(eventName+err.Error(), "sendlog_errors")
	}
}

//// SendToMatchFound sends a notification if a match is found to Telegram
//func SendToMatchFound(bot *tgbotapi.BotAPI, text string, evenName string) {
//	replyMsg := tgbotapi.NewMessage(constants.HootCommandControlID, "```\n"+text+"\n```")
//	replyMsg.ReplyToMessageID = EventnameToTelegramChannel[eventName]
//}
