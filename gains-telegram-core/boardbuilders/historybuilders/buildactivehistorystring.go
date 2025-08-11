package historybuilders

import (
	"HootCore/history"
	"fmt"
)

func BuildActiveHistoryString(history history.TradeHistory) string {
	// Determine if the PnL is positive or negative and set appropriate symbol and color
	var symbol, color string
	if history.PnlNet >= 0 {
		symbol = "📈"
		color = "🟢"
	} else {
		symbol = "📉"
		color = "🔴"
	}

	// Calculate percentage change
	percentageChange := (history.PnlNet / history.CollateralPriceUsd) * 100

	// Format date
	dateString := history.Date.Format("15:04 UTC - Jan 02, 2006")

	// Build the active history string
	activeHistoryString := fmt.Sprintf(
		"⚖️ Type: Market\n🛠️ Leverage: %.0fx\n💲 Price: %.2f %s\n💰 Size: %.2f USD\n📅 Closed: %s\n📜 Overview: %s %+0.2f%% %s %+0.2f\n",
		history.Leverage, history.Price, history.Pair, history.Size, dateString, symbol, percentageChange, color, history.PnlNet,
	)

	return activeHistoryString
}
