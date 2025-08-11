package altseasoncore

//func GetPnl(rdbHoot *rdbhoot.HootRedisClient, posDai float64, trade TradeItemCore, tradeInfo TradeInfo, initialAccFees InitialAccFeesCore, useFees bool) (float64, float64, error) {
//	currentPrice, err := priceserver.GetPrice(rdbHoot, int(trade.Trade.PairIndex))
//	if err != nil {
//		log.Println("Error getting price from Redis:", err)
//	}
//
//	if currentPrice == 0 {
//		return 0, 0, nil
//	}
//
//	iacConv := initialAccFees.IntoConverted()
//
//	context, err := GetGainsTradeContext(rdbHoot, 421614)
//	if err != nil {
//		log.Println("Error getting trading system config from Redis:", err)
//	}
//
//	contextConv := context.ToConverted()
//
//	posCollat := trade.Trade.CollateralAmount
//	openPrice := trade.Trade.OpenPrice
//	leverage := trade.Trade.Leverage
//	maxGainP := context.MaxGainP
//	//feeFloat := context.Fees[trade.Trade.PairIndex]
//
//	var maxGain float64
//	if maxGainP == 0 {
//		maxGain = math.Inf(1)
//	} else {
//		maxGain = (float64(maxGainP) / 100.00) * posCollat
//	}
//
//	var pnlCollat float64
//	if trade.Trade.Long {
//		pnlCollat = ((currentPrice - openPrice) / openPrice) * float64(leverage) * posCollat
//	} else {
//		pnlCollat = ((openPrice - currentPrice) / openPrice) * float64(leverage) * posCollat
//	}
//
//	if pnlCollat > maxGain {
//		pnlCollat = maxGain
//	}
//
//	if useFees {
//		pnlCollat -= GetBorrowingFee(
//			posDai,
//			int(trade.Trade.PairIndex),
//			trade.Trade.Long,
//			iacConv,
//			contextConv,
//		)
//	}
//
//	pnlPercentage := (pnlCollat / posCollat) * 100
//
//	if pnlPercentage <= -90 {
//		pnlPercentage = -100
//	} else {
//		pnlCollat -= GetClosingFee(posCollat, float64(leverage), int(trade.Trade.PairIndex), context.ToConverted())
//		pnlPercentage = (pnlCollat / posCollat) * 100
//	}
//
//	if pnlPercentage < -100 {
//		pnlPercentage = -100
//	}
//	pnlCollat = (posCollat * pnlPercentage) / 100
//
//	return pnlCollat, pnlPercentage, nil
//}
