package boardbuilders

import (
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"log"
)

func BuildFeesBoard(rdb *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	trade, exists := rdb.GetCoreCache(guid)
	if !exists {
		log.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	var newTradeStr string

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	// Check on which Window we are:
	if trade.ActiveWindow == 0 { // We are on new trade
		newTradeStr = "⭐️ NEW TRADE"
	} else {
		if trade.ActiveWindow == 1 {
			newTradeStr = "NEW TRADE"
		} else {
			log.Println("Error during BuildPerpMainBoard: ActiveWindow not equal to 0 or 1")
		}
	}

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
		),
	)
}
