package snipingboards

import tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"

var SniperMainBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("New Snipe", "/"),
		tgbotapi.NewInlineKeyboardButtonData("Active Snipes", "/refreshposition"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Hoot?", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Refresh?", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("📈 Charts", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✅ BUY", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("🔴 WALLETS", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("📈 WALLET", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("-", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("ETH", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("+", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("0.1 ETH", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("0.5 ETH", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("MAX", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("SNIPE $TOKEN", "/custom_tp_open_position"),
	),
)
