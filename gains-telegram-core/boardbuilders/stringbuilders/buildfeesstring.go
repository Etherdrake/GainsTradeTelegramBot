package stringbuilders

//func BuildFeesString(
//	client *mongo.Client,
//	rdbHoot *hootrdb.HootRedisClient,
//	guid int64,
//	tradeCache *tradecache.TradeCache,
//	chain uint64,
//) string {
//	log.Println(time.Now())
//	trade, exists := tradeCache.Get(guid)
//	if !exists {
//		log.Println("ERROR: trade not exists")
//	}
//
//	context, err := altseasoncore.GetGainsTradeContext(rdbHoot, 421614)
//	if err != nil {
//		return ""
//	}
//
//	convContext := context.ToConverted()
//
//	//Opening Fee: 0
//	//Closing Fee: 0,
//	//💰 Fixed Spread: 0
//	//🤸 Dynamic Spread: 0
//	//Price Impact:  0
//	//Borrowing Fee / H: 0
//
//	trade, exists = tradeCache.Get(guid)
//
//	// Opening Fee
//
//	// Closing Fee
//
//	closingFee := altseasoncore.GetClosingFee(trade.PositionSizeDai, float64(trade.Leverage), int(trade.PairIndex), convContext)
//
//	// Fixed Spread
//
//	fixedSpread, _ := altseasoncore.GetFixedSpreadPercentage(trade.PairIndex, convContext)
//
//	// Dynamic Spread
//
//	dynSpread, _ := altseasoncore.GetDynamicSpread(int(trade.PairIndex), int(trade.Collateral), trade.PositionSizeDai, trade.Buy, convContext)
//
//	// Price Impact
//
//	//altseasoncore.GetSpreadWithPriceImpactP()
//
//	// Borrowing Fee
//
//	// We need the PairFee and the GroupFee
//
//	initialFees := altseasoncore.InitialAccFeesConverted{
//		AccPairFee:  0,
//		AccGroupFee: 0,
//		Block:       0,
//	}
//
//	borrowingFee := altseasoncore.GetBorrowingFee(trade.PositionSizeDai, int(trade.PairIndex), trade.Buy, initialFees, context.ToConverted())
//
//	msg := fmt.Sprintf(`
//💲 Selected Instrument: *%s*
//
//Opening Fee: %s
//Closing Fee: %.10f,
//💰 Fixed Spread: *%.10f*
//🤸 Dynamic Spread: %.10f
//Price Impact:  %s
//Borrowing Fee / H: %.10f`,
//		utils.EscapeMarkdownV2(pairmaps.IndexToPair[int(trade.PairIndex)]),
//		"0",
//		closingFee,
//		fixedSpread,
//		dynSpread, // Notice this change
//		"0",
//		borrowingFee,
//	)
//	return msg
//}
