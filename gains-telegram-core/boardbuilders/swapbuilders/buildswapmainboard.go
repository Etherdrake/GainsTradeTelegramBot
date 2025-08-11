package swapbuilders

import (
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
)

func BuildSwapBoardV1(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	//size := strconv.FormatInt(trade.PositionSizeDai, 10)

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
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
		//	tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("MAX", "/maxsrcamount"),
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/swaprefresh"),
			tgbotapi.NewInlineKeyboardButtonData("CHART", "/sendchart"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("DAI"+" 🔽", "/srctokenlist"),
			tgbotapi.NewInlineKeyboardButtonData("TO", "/swap"),
			tgbotapi.NewInlineKeyboardButtonData("USDC"+" 🔽", "/dsttokenlist"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
		//	longShortButton,
		//	tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("5000 DAI", "/customsrcamount"),
			tgbotapi.NewInlineKeyboardButtonData("FOR", "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("4999 USDC", "/customdstamount"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Swap DAI for USDC", "/submitswap"),
		),
	)
}
