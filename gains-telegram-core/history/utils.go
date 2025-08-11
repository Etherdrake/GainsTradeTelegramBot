package history

import (
	"fmt"
	"math/big"
)

// CalculateHistoryPnl calculates the PnL as a percentage string
func CalculateHistoryPnl(history TradeHistory) string {
	// Use big.Float for precision
	pnlNetBig := big.NewFloat(history.PnlNet)
	collateralPriceUsdBig := big.NewFloat(history.CollateralPriceUsd)

	// Perform the division and multiplication with precision
	percentageChangeBig := new(big.Float).Quo(pnlNetBig, collateralPriceUsdBig)
	percentageChangeBig.Mul(percentageChangeBig, big.NewFloat(100))
	percentageChange, _ := percentageChangeBig.Float64()

	// Format the percentage change as a string
	pnlPercentageString := fmt.Sprintf("%.2f%%", percentageChange)

	return pnlPercentageString
}
