package altseasoncore

func GetClosingFee(posDai float64, leverage float64, pairIndex int, context GainsTradeContextConverted) float64 {
	if posDai == 0 || leverage == 0 {
		return 0
	}

	pairFee := context.Fees[pairIndex] // Lenth of Fees is 10

	return (pairFee.CloseFeeP + pairFee.TriggerOrderFeeP) * posDai * leverage
}
