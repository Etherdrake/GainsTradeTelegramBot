package altseasoncore

import "math"

func GetLiquidationPrice(trade TradeItemCore, tradeCore TradeItemCore, context GainsTradeContextConverted) float64 {
	iacConv := trade.InitialAccFees.IntoConverted()

	closingFee := GetClosingFee(trade.Trade.CollateralAmount, float64(trade.Trade.Leverage), int(trade.Trade.PairIndex), context)
	borrowingFee := GetBorrowingFee(trade.Trade.CollateralAmount*float64(trade.Trade.Leverage), int(trade.Trade.PairIndex), trade.Trade.Long, iacConv, context)

	liqPriceDistance := (trade.Trade.OpenPrice * (trade.Trade.CollateralAmount*0.9 - (borrowingFee + closingFee))) / trade.Trade.CollateralAmount / float64(trade.Trade.Leverage)

	if trade.Trade.Long {
		return math.Max(trade.Trade.OpenPrice-liqPriceDistance, 0)
	} else {
		return math.Max(trade.Trade.OpenPrice+liqPriceDistance, 0)
	}
}
