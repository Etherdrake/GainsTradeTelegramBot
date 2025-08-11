package boardbuilders

import (
	"HootCore/altseasoncore"
	"HootCore/formatters"
	"HootCore/mongolisten"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

func BuildActiveTradeBoardGNSV2(
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	guid int64,
	chain uint64) tgbotapi.InlineKeyboardMarkup {
	//publicKey, err := database.GetPublicKey(guid)
	//if err != nil {
	//	fmt.Println("PublicKey Error: ", err)
	//}

	hasTradeOrOrder, _, _, err := mongolisten.HasOpenTradeOrOrder(client, guid, chain)
	if err != nil {
		return tgbotapi.InlineKeyboardMarkup{}
	}

	isTradeOrderOrderCode, err := rdbHoot.GetIsTrade(guid)

	var activeCacheTrade altseasoncore.TradeItemInternal
	var activeCacheOrder altseasoncore.TradeItemInternal

	newTradeStr := "NEW TRADE"
	activeTradesStr := "⭐️ OPEN TRADES"

	var openTradeActive bool
	var openOrderActive bool
	var currentPrice float64
	if hasTradeOrOrder {
		switch isTradeOrderOrderCode {
		case 1:
			activeCacheTrade, err = rdbHoot.GetSelectedTrade(guid)
			if err != nil {
				log.Println("Unable to retrieve active trade from cache in BuildActiveTradeBoardGNSV2")
				break
			}
			openTradeActive = true
			openOrderActive = false
			// Retrieve the price from Redis
			currentPrice, err = rdbHoot.GetPrice(int(activeCacheTrade.Trade.PairIndex))
			if err != nil {
				fmt.Println("Error fetching price from Redis in Stringbuilder:", err)
			}

		case 2:
			activeCacheOrder, err = rdbHoot.GetSelectedOrder(guid)
			if err != nil {
				log.Println("Unable to retrieve active order from cache in BuildActiveTradeBoardGNSV2")
				break
			}
			openTradeActive = false
			openOrderActive = true
			currentPrice, err = rdbHoot.GetPrice(int(activeCacheOrder.Trade.PairIndex))
			if err != nil {
				fmt.Println("Error fetching price from Redis in Stringbuilder:", err)
			}
		}
	} else {
		log.Println("No active trade or order for the given user.")
		// User has no opentrade or openorder
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("NEW TRADE", "/newtrade"),
				tgbotapi.NewInlineKeyboardButtonData("⭐️ ACTIVE TRADES", "/activetrades"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Refresh", "/"),
				tgbotapi.NewInlineKeyboardButtonData("Chart", "/"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/"),
				tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Share on TG", "/"),
				tgbotapi.NewInlineKeyboardButtonData("Share on X", "/"),
			),
			tgbotapi.NewInlineKeyboardRow(
				// The position in question
				tgbotapi.NewInlineKeyboardButtonData("No trade or order", "/")),
		)
	}

	//// If trade or order /t or /o + ":" + "SYMBOL"
	//symbol := pairmaps.IndexToCrypto[int(pairIdx)]

	//collConv, _ := strconv.ParseFloat(activeCacheTradeDisplay.Trade.CollateralAmount, 64)

	if openTradeActive {
		// TODO: CollAmount needs to be defined as current price x quantity
		getPnlUSD := utils.CalculateDollarProfitOrLossWithLeverage(activeCacheTrade.Trade.CollateralAmount, activeCacheTrade.Trade.OpenPrice, currentPrice, activeCacheTrade.Trade.Leverage, activeCacheTrade.Trade.Long)

		log.Println("activeCacheTrade: ", activeCacheTrade.Trade)
		log.Println("PnL USD: ", getPnlUSD)

		var pnlStr string
		var pnlForm string
		if getPnlUSD > 0 {
			plusSymbol := "+"
			pnlStr = strconv.FormatFloat(getPnlUSD, 'f', 2, 64)
			pnlStr = plusSymbol + pnlStr
			pnlForm, err = formatters.FormatFinancialUS(pnlStr)
			if err != nil {
				log.Println("Error formatting pnl form:", err)
			}
		} else {
			pnlStr = strconv.FormatFloat(getPnlUSD, 'f', 2, 64)
			pnlForm, err = formatters.FormatFinancialUS(pnlStr)
			if err != nil {
				log.Println("Error formatting pnl form:", err)
			}
		}
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
				tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/activetrades"),
				tgbotapi.NewInlineKeyboardButtonData("CHART", "/sendchart"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/tpedit"),
				tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/sledit"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Share on TG", "/activetrades"),
				tgbotapi.NewInlineKeyboardButtonData("Share on X", "/activetrades"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("CLOSE TRADE ("+pnlForm+")", "/closetrade")),
		)
	}

	if openOrderActive {
		return tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(newTradeStr, "/newtrade"),
				tgbotapi.NewInlineKeyboardButtonData(activeTradesStr, "/activetrades"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("REFRESH", "/activetrades"),
				tgbotapi.NewInlineKeyboardButtonData("Entry", "/updateopenlimitprice"),
				tgbotapi.NewInlineKeyboardButtonData("CHART", "/sendchart"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Take Profit", "/tpedit"),
				tgbotapi.NewInlineKeyboardButtonData("Stop Loss", "/sledit"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Share on TG", "/activetrades"),
				tgbotapi.NewInlineKeyboardButtonData("Share on X", "/activetrades"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Cancel Order", "/cancelorder")),
		)
	}
	return tgbotapi.InlineKeyboardMarkup{}
}
