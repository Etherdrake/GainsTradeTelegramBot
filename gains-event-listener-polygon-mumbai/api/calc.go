package api

import (
	"HootTelegram/pairmaps"
	"fmt"
	"strconv"
)

func calculatePairSpread(trade OpenTradeJSON, payload PayloadIndexChainGains) float64 {
	var pairSpread float64
	// Check if pair is crypto, forex major, forex minor or forex exotic
	// or Commidities tier-1 or tier-2
	indexInt, _ := strconv.Atoi(trade.PairIndex)
	if _, ok := pairmaps.IndexToCrypto[indexInt]; ok {
		// Pair is crypto
	} else {
		if _, ok = pairmaps.IndexToFx[indexInt]; ok {
			// Pair is Forex
			// Check if Major or Minor
		} else {
			if _, ok = pairmaps.IndexToCommodity[indexInt]; ok {
				// Pair is Commodity

				// Ch

			}
		}
	}

	// Fixed spread is 0.04%
	fixedSpead := 0.0001

	// Retrieve 1% Depth
	tradeDepth := 0.01
	// Dynamic Spread (%) = (Open interest {long/short} + New trade position size / 2) / 1% depth {above/below}.
	if tradeDepth != 0 {
		if trade.Buy {
			// Retrieve Open Interest Long

			// Retrieve 1% Depth above
		} else {
			// Retrieve Open Interest Short

			// Retrieve 1% Depth above
		}
	} else {
		// Fixed spread is returned
		return 0.004
	}
	fmt.Println(fixedSpead)
	return pairSpread
}
