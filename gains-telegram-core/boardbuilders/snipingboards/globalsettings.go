package snipingboards

import tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"

var SniperGlobalSettingsBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Global", "/"),
		tgbotapi.NewInlineKeyboardButtonData("Buy", "/refreshposition"),
		tgbotapi.NewInlineKeyboardButtonData("Sell", "/refreshposition"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("MultiWallet", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Anti-Rug", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Anti-MEV", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Auto-Slip", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Anti-BL", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Pre-Approve", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Slippage", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Max", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Approve", "/charts"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Bundle Tips", "/custom_tp_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Tip", "/custom_sl_open_position"),
		tgbotapi.NewInlineKeyboardButtonData("Block", "/charts"),
	),
)
