package tpslbuilders

import (
	"HootCore/formatters"
	"HootCore/rdbhoot"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"log"
)

func BuildSlBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	var newTradeStr string
	var activeTradesStr string

	// Check on which Window we are:
	if user.ActiveWindow == 0 { // We are on new user
		newTradeStr = "⭐️ NEW TRADE"
		activeTradesStr = "ACTIVE TRADES"
	} else {
		if user.ActiveWindow == 1 {
			newTradeStr = "NEW TRADE"
			activeTradesStr = "⭐️ ACTIVE TRADES"
		} else {
			log.Println("Error during BuildPerpMainBoard: ActiveWindow not equal to 0 or 1")
		}
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if user.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	//pairStr := pairmaps.IndexToPair[int(user.PairIndex)]

	stopLoss := formatters.FormatPrice(user.SL)

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
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("🛑 Stop Loss", "/leverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrsl"),
			tgbotapi.NewInlineKeyboardButtonData(stopLoss, "/customsl"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrsl"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("None", "/zerosl"),
			tgbotapi.NewInlineKeyboardButtonData("-10%", "/minus10"),
			tgbotapi.NewInlineKeyboardButtonData("-25%", "/minus25"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("-33%", "/minus33"),
			tgbotapi.NewInlineKeyboardButtonData("-50%", "/minus50"),
			tgbotapi.NewInlineKeyboardButtonData("-75%", "/minus75"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Save", "/leverage"),
		),
	)
}
