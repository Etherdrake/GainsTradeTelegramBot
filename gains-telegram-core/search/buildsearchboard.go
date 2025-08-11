package search

import "github.com/Etherdrake/telegram-bot-api/v7"

func BuildSearchBoard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 BACK", "/pairs"),
		),
	)
}
