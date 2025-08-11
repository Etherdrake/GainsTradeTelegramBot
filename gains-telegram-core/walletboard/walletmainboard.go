package walletboard

import tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"

var WalletMainBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "/backmain"),
	),
)
