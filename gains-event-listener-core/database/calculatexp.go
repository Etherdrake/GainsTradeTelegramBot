package database

import "GainsListenerV4/types"

func CalculateXP(trade types.TradeInternal) float64 {
	// Determine PnLFactor based on profitLossPercent
	var PnLFactor float64
	profitLossPercent := (trade.Tp - trade.Sl) / trade.OpenPrice * 100.0

	if profitLossPercent <= -50 {
		PnLFactor = 0.5
	} else if profitLossPercent >= 80 {
		PnLFactor = 1.8
	} else {
		// Linear interpolation for profitLossPercent between -50 and 80
		PnLFactor = 0.5 + (profitLossPercent+50)*(1.3/130)
	}

	// Calculate XP
	XP := (trade.CollateralAmount / 10) * PnLFactor

	return XP
}
