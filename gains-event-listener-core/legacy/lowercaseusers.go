package legacy

import (
	"GainsListenerV4/types"
	"strings"
)

func LowercaseUsers(openTrades []types.OpenTrade) []types.OpenTrade {
	var newOpenTrades []types.OpenTrade

	for _, trade := range openTrades {
		trade.Trade.User = strings.ToLower(trade.Trade.User)
		newOpenTrades = append(newOpenTrades, trade)
	}

	return newOpenTrades
}
