package legacy

import "GainsListenerV4/types"

// CompareTradeItems compares two slices of TradeItemMongo and returns the identifiers of differing trades.
func CompareTradeItems(oldItems, newItems []types.TradeItemMongo) []string {
	// Create a map to store the first set of trade items for quick lookup
	tradeMap1 := make(map[string]types.TradeItemMongo)
	for _, trade := range oldItems {
		tradeMap1[trade.ID] = trade
	}

	// Create a map to store the second set of trade items for quick lookup
	tradeMap2 := make(map[string]types.TradeItemMongo)
	for _, trade := range newItems {
		tradeMap2[trade.ID] = trade
	}

	// Create a slice to store the identifiers of differing trades
	var differingIdentifiers []string

	// Compare the first set of trade items against the second set to (find any trades present in the database but not in the newTrades)
	for id, trade1 := range tradeMap1 {
		if trade2, exists := tradeMap2[id]; !exists || !areTradesEqual(trade1, trade2) {
			differingIdentifiers = append(differingIdentifiers, id)
		}
	}

	return differingIdentifiers
}

// areTradesEqual checks if two TradeItemMongo objects are equal
func areTradesEqual(trade1, trade2 types.TradeItemMongo) bool {
	return trade1.Trade == trade2.Trade &&
		trade1.TradeInfo == trade2.TradeInfo &&
		trade1.InitialAccFees == trade2.InitialAccFees
}
