package newtradebuilders

import (
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"log"
	"strconv"
	"strings"
)

func BuildNewTradeBoardCharts(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("5M", "/leveragechart5M"),
			tgbotapi.NewInlineKeyboardButtonData("15M", "/leveragechart15M"),
			tgbotapi.NewInlineKeyboardButtonData("1H", "/leveragechart1H"),
			tgbotapi.NewInlineKeyboardButtonData("4H", "/leveragechart4H"),
			tgbotapi.NewInlineKeyboardButtonData("1D", "/leveragechart1D"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}

func BuildNewTradeBoardCharts5M(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1M", "/leveragechart"),
			tgbotapi.NewInlineKeyboardButtonData("15M", "/leveragechart15M"),
			tgbotapi.NewInlineKeyboardButtonData("1H", "/leveragechart1H"),
			tgbotapi.NewInlineKeyboardButtonData("4H", "/leveragechart4H"),
			tgbotapi.NewInlineKeyboardButtonData("1D", "/leveragechart1D"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}

func BuildNewTradeBoardCharts15M(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1M", "/leveragechart"),
			tgbotapi.NewInlineKeyboardButtonData("5M", "/leveragechart5M"),
			tgbotapi.NewInlineKeyboardButtonData("1H", "/leveragechart1H"),
			tgbotapi.NewInlineKeyboardButtonData("4H", "/leveragechart4H"),
			tgbotapi.NewInlineKeyboardButtonData("1D", "/leveragechart1D"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}

func BuildNewTradeBoardCharts1H(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1M", "/leveragechart"),
			tgbotapi.NewInlineKeyboardButtonData("5M", "/leveragechart5M"),
			tgbotapi.NewInlineKeyboardButtonData("15M", "/leveragechart15M"),
			tgbotapi.NewInlineKeyboardButtonData("4H", "/leveragechart4H"),
			tgbotapi.NewInlineKeyboardButtonData("1D", "/leveragechart1D"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}

func BuildNewTradeBoardCharts4H(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1M", "/leveragechart"),
			tgbotapi.NewInlineKeyboardButtonData("5M", "/leveragechart5M"),
			tgbotapi.NewInlineKeyboardButtonData("15M", "/leveragechart15M"),
			tgbotapi.NewInlineKeyboardButtonData("1H", "/leveragechart1H"),
			tgbotapi.NewInlineKeyboardButtonData("1D", "/leveragechart1D"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}

func BuildNewTradeBoardCharts1D(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

	pairStr := pairmaps.IndexToPair[int(trade.PairIndex)]

	if len(pairStr) > 7 {
		pairStr = utils.TrimToSeven(pairStr)
	}

	var longShortButton tgbotapi.InlineKeyboardButton
	if trade.Buy {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🟢 LONG", "/toggletoshort")
	} else {
		longShortButton = tgbotapi.NewInlineKeyboardButtonData("🔴 SHORT", "/toggletolong")
	}

	size := strconv.FormatFloat(trade.PositionSizeDai, 'f', 2, 64)       // Format as decimal with 2 decimal places
	leverage := strconv.FormatFloat(float64(trade.Leverage), 'f', 2, 64) // Format as decimal with 2 decimal places

	var priceReal string
	if trade.OrderType != 0 {
		priceReal = formatters.FormatPrice(trade.OpenPrice)
		prePendStr := "$"
		priceReal = prePendStr + priceReal
	} else {
		priceReal = "MARKET"
	}

	var orderTypeStr string
	if trade.OrderType == 0 {
		orderTypeStr = "MARKET"
	}
	if trade.OrderType == 1 {
		orderTypeStr = "LIMIT"
	}
	if trade.OrderType == 2 {
		orderTypeStr = "STOP"
	}

	//hyperlinkStr := utils.ReplaceSlashWithDash(pairStr)

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1M", "/leveragechart"),
			tgbotapi.NewInlineKeyboardButtonData("5M", "/leveragechart5M"),
			tgbotapi.NewInlineKeyboardButtonData("15M", "/leveragechart15M"),
			tgbotapi.NewInlineKeyboardButtonData("1H", "/leveragechart1H"),
			tgbotapi.NewInlineKeyboardButtonData("4H", "/leveragechart4H"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
			tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/refresh"),
			tgbotapi.NewInlineKeyboardButtonData(pairStr+" 🔽", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("BALANCES", "/leverageback"), // https://gains.trade/trading#ETH-USD
			//tgbotapi.NewInlineKeyboardButtonURL("CHART", "https://gains.trade/trading#"+hyperlinkStr), // https://gains.trade/trading#ETH-USD
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/takeprofit"),
			longShortButton,
			tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/stoploss"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrprice"),
			tgbotapi.NewInlineKeyboardButtonData(priceReal, "/customprice"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrprice"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrpossize"),
			tgbotapi.NewInlineKeyboardButtonData("$"+size, "/custompossize"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrpossize"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➖", "/decrleverage"),
			tgbotapi.NewInlineKeyboardButtonData(leverage+"X", "/customleverage"),
			tgbotapi.NewInlineKeyboardButtonData("➕", "/incrleverage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("SUBMIT "+orderTypeStr+" ("+strings.Trim(longShortButton.Text, "🟢🔴 ")+")", "/submitconfirm"),
		),
	)
}
