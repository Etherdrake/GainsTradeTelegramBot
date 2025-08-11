package stringbuilders

import (
	"HootCore/altseasoncore"
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	"log"
	"strconv"
)

func BuildSingleTradeStringGNS(rdbHoot *rdbhoot.HootRedisClient, trade altseasoncore.TradeItemInternal, index int, isTrade bool) (string, error) {
	dTrade := trade.Trade.IntoDisplay()

	convert := int(trade.Trade.PairIndex)

	// TODO: Be aware that this needs to be reworked for WETH and native gas tokens

	// Retrieve the price from Redis
	currentPrice, err := rdbHoot.GetPrice(convert) // Assuming you have the GetPrice function implemented to fetch price from Redis
	if err != nil {
		log.Println("Error fetching price from Redis in string builder:", err)
	}

	openPriceF64, err := strconv.ParseFloat(dTrade.OpenPrice, 64)
	leverageF32, err := strconv.ParseFloat(dTrade.Leverage, 32)
	collateralAmount, err := strconv.ParseFloat(dTrade.CollateralAmount, 64)

	// initialInvestment, entryPrice, currentPrice, leverage float64, isLong bool
	getPnlUSD := utils.CalculateDollarProfitOrLossWithLeverage(collateralAmount, openPriceF64, currentPrice, float64(leverageF32), trade.Trade.Long)
	positionSizeF64, err := strconv.ParseFloat(dTrade.CollateralAmount, 64)
	getPnlPrcnt := utils.CalculatePercentageProfitOrLoss(positionSizeF64, getPnlUSD, trade.Trade.Long)
	pnlPrcntFormatted, err := formatters.FormatFinancialUS(getPnlUSD)
	if err != nil {
		log.Println("Error in buildsinglepositionstring: ", err)
	}

	//createdBlock := trade.TradeInfo.CreatedBlock

	tradeDirection := "📉"
	if trade.Trade.Long {
		tradeDirection = "📈"
	}

	var plusOrMinus string
	var theDot string
	if getPnlUSD > 0 {
		plusOrMinus = "+"
		theDot = "🟢"
	} else {
		plusOrMinus = ""
		theDot = "🔴"
	}

	var prefix string
	if isTrade {
		prefix = "t"
	} else {
		prefix = "o"
	}

	positionString := fmt.Sprintf("/%s%d %s %s %s%.2f%% %s %s*%s*",
		prefix,
		index,
		pairmaps.IndexToPair[int(trade.Trade.PairIndex)],
		tradeDirection,
		plusOrMinus,
		getPnlPrcnt,
		theDot,
		plusOrMinus,
		pnlPrcntFormatted,
	)
	return positionString, nil
}
