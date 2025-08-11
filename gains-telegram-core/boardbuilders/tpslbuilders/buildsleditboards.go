package tpslbuilders

import (
	"HootCore/formatters"
	"HootCore/rdbhoot"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"log"
)

func BuildSlEditBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}
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
	switch trade.ActivePerp.IsTradeOrOrder {
	case 0:
		return tgbotapi.InlineKeyboardMarkup{}
	case 1:
		openPrice = formatters.FormatPrice(trade.ActivePerp.ActiveTradeOrOrder.Trade.OpenPrice)
	case 2:
		openPrice = formatters.FormatPrice(trade.ActivePerp.ActiveTradeOrOrder.Trade.OpenPrice)
	default:
		return tgbotapi.InlineKeyboardMarkup{}
	}

	//pairStr := pairmaps.IndexToPair[int(user.PairIndex)]

	return tgbotapi.NewInlineKeyboardMarkup(
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("⬅️", "/previouspair"),
		//	tgbotapi.NewInlineKeyboardButtonData("🔄 "+pairStr, "/refresh"),
		//	tgbotapi.NewInlineKeyboardButtonData("➡️", "/nextpair"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrtpedit"),
			tgbotapi.NewInlineKeyboardButtonData(openPrice, "/customsledit"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrsledit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("None", "/zerosledit"),
			tgbotapi.NewInlineKeyboardButtonData("-10%", "/minus10edit"),
			tgbotapi.NewInlineKeyboardButtonData("-25%", "/minus25edit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("-33%", "/minus33edit"),
			tgbotapi.NewInlineKeyboardButtonData("-50%", "/minus50edit"),
			tgbotapi.NewInlineKeyboardButtonData("-75%", "/minus75edit"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Back", "/activetrades"),
			tgbotapi.NewInlineKeyboardButtonData("💾 Save", "/savesledit"),
		),
	)
}
