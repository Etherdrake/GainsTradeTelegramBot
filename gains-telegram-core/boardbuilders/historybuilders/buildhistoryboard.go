package historybuilders

import (
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
)

func BuildHistoryBoard(txReceipt string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("NEW TRADE", "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData("ACTIVE TRADES", "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/history"),
			tgbotapi.NewInlineKeyboardButtonURL("TX RECEIPT", "https://sepolia.arbiscan.io/tx/"+txReceipt),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Share on TG", "/history"),
			tgbotapi.NewInlineKeyboardButtonData("Share on X", "/history"),
		),
	)
}
