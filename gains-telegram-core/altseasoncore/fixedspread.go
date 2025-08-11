package altseasoncore

// CalculateFixedSpread calculates the Fixed Spread for a given pairIndex and GainsTradeContext
func GetFixedSpreadPercentage(pairIndex int64, config GainsTradeContextConverted) (float64, error) {
	// Return the spread percentage
	return config.Pairs[pairIndex].SpreadP, nil
}
