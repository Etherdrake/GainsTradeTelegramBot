package tpslbuilders

import (
	"HootCore/formatters"
	"HootCore/rdbhoot"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"log"
)

func BuildTpEditBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	displayTrade := trade.ActivePerp.ActiveTradeOrOrder.IntoDisplay()

	var newTradeStr string
	var activeTradesStr string

	// Check on which Window we are:
	if trade.ActiveWindow == 0 { // We are on new trade
		newTradeStr = "⭐️ NEW TRADE"
		activeTradesStr = "ACTIVE TRADES"
	} else {
		if trade.ActiveWindow == 1 {
			newTradeStr = "NEW TRADE"
			activeTradesStr = "⭐️ ACTIVE TRADES"
		} else {
			log.Println("Error during BuildPerpMainBoard: ActiveWindow not equal to 0 or 1")
		}
	}

	var openPrice string
	var err error
	switch trade.ActivePerp.IsTradeOrOrder {
	case 0:
		return tgbotapi.InlineKeyboardMarkup{}
	case 1:
		openPrice, err = formatters.FormatFinancialUS(displayTrade.Trade.OpenPrice)
		if err != nil {
			log.Println("Error during BuildTpEditBoard:", err)
		}
	case 2:
		openPrice, err = formatters.FormatFinancialUS(displayTrade.Trade.OpenPrice)
		if err != nil {
			log.Println("Error during BuildTpEditBoard:", err)
		}
	default:
		return tgbotapi.InlineKeyboardMarkup{}
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrtpedit"),
			tgbotapi.NewInlineKeyboardButtonData(openPrice, "/customtpedit"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrtpedit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("None", "/zerotpedit"),
			tgbotapi.NewInlineKeyboardButtonData("25%", "/plus25edit"),
			tgbotapi.NewInlineKeyboardButtonData("50%", "/plus50edit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("100%", "/plus100edit"),
			tgbotapi.NewInlineKeyboardButtonData("200%", "/plus150edit"),
			tgbotapi.NewInlineKeyboardButtonData("900%", "/plus900edit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Back", "/activetrades"),
			tgbotapi.NewInlineKeyboardButtonData("💾 Save", "/savetpedit"),
		),
	)
}
