package altseasontools

import (
	"HootCore/altseasoncore"
	"sort"
)

//func FindTradeIndex(openTrades []altseasontypes.TradeItemCore) int {
//	for i := 0; range openTrades {
//
//	}
//}

// SortTrades sorts a slice of TradeItemCore by the CreatedBlock field in ascending order
func SortTrades(openTrades []altseasoncore.TradeItemCore) []altseasoncore.TradeItemCore {
	sort.Slice(openTrades, func(i, j int) bool {
		return openTrades[i].TradeInfo.CreatedBlock < openTrades[j].TradeInfo.CreatedBlock
	})
	return openTrades
}

// SortTradesLatestLast sorts a slice of TradeItemCore by the CreatedBlock field in ascending order
func SortTradesLatestLast(openTrades []altseasoncore.TradeItemCore) []altseasoncore.TradeItemCore {
	sort.Slice(openTrades, func(i, j int) bool {
		return openTrades[i].TradeInfo.CreatedBlock > openTrades[j].TradeInfo.CreatedBlock
	})
	return openTrades
}
