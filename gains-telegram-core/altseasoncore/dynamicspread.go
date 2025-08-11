package altseasoncore

import (
	"fmt"
)

// GetDynamicSpread calculates the Dynamic Spread for a given pairIndex and GainsTradeContext
func GetDynamicSpread(pairIndex, collateralIndex int, positionSize float64, direction bool, config GainsTradeContextConverted) (float64, error) {
	var openInterest float64
	var depth float64

	// Set OI and depth based on trade direction
	switch direction {
	case true:
		openInterest = config.Collaterals[collateralIndex-1].BorrowingFees.Pairs[pairIndex].Oi.Long
		depth = float64(config.PairInfos.PairDepths[pairIndex].OnePercentDepthAboveUsd)
	case false:
		openInterest = config.Collaterals[collateralIndex-1].BorrowingFees.Pairs[pairIndex].Oi.Short
		depth = float64(config.PairInfos.PairDepths[pairIndex].OnePercentDepthBelowUsd)
	default:
		return 0.0, fmt.Errorf("unknown direction: %s", direction)
	}

	// Calculate Dynamic Spread
	dynamicSpread := (openInterest + positionSize/2.0) / depth

	return dynamicSpread, nil
}
