package stringbuilders

import (
	"HootCore/altseasoncore"
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"fmt"
	"strconv"
)

func BuildSingleOrderStringGNS(rdbHoot *rdbhoot.HootRedisClient, openOrder altseasoncore.TradeItemInternal, arrIdx int, orderID string) (string, error) {
	dOrder := openOrder.Trade.IntoDisplay()

	openPriceF64, err := strconv.ParseFloat(dOrder.OpenPrice, 64)

	convert := int(openOrder.Trade.PairIndex)

	// Retrieve the price from Redis
	currentPrice, err := rdbHoot.GetPrice(convert) // Assuming you have the GetPrice function implemented to fetch price from Redis
	if err != nil {
		fmt.Println("Error fetching price from Redis in Stringbuilder:", err)
	}

	//leverage, err := strconv.ParseFloat(trade.Leverage, 64)

	tradeDirection := "📉"
	tradeDirectionStr := "Short"
	if openOrder.Trade.Long {
		tradeDirection = "📈"
		tradeDirectionStr = "Long"
	}

	tickerInt := int(openOrder.Trade.PairIndex)

	positionString := fmt.Sprintf("/o*%d* %s %s %s @ *$%s*, now: $%s",
		arrIdx,
		pairmaps.IndexToPair[tickerInt],
		tradeDirection,
		tradeDirectionStr,
		formatters.FormatPrice(openPriceF64),
		formatters.FormatPrice(currentPrice),
	)
	return positionString, nil
}
