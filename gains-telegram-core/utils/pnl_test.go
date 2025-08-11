package utils

import "testing"

func TestCalculateDollarProfitOrLossWithLeverage(t *testing.T) {
	initialInvestment := 928.00
	entryPrice := 3166.18
	currentPrice := 3161.73
	leverage := 100.0
	isLong := true

	expectedProfit := -130.506516

	result := CalculateDollarProfitOrLossWithLeverage(initialInvestment, entryPrice, currentPrice, leverage, isLong)

	if result != expectedProfit {
		t.Errorf("Expected profit: %f, but got: %f", expectedProfit, result)
	}
}
