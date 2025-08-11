package historybuilders

import (
	"HootCore/history"
	"fmt"
	"log"
	"strings"
)

// BuildHistoryString builds a history string for a user's trades
func BuildHistoryString(historicalTrades []history.TradeHistory, guid int64, page, activeIdx int) (string, error) {
	var historyStrings []string
	tradeCount := len(historicalTrades)
	if tradeCount > 20 {
		tradeCount = 20
	}

	log.Println(historicalTrades[0])

	historyStrings = append(historyStrings, "📕️ Trade History\n")

	for i := 0; i < tradeCount; i++ {
		if i == activeIdx {
			historySentence := BuildHistorySentence(historicalTrades[i], i, true)
			if historySentence != "" {
				historyStrings = append(historyStrings, historySentence)
				historyStrings = append(historyStrings, BuildActiveHistoryString(historicalTrades[i]))
			}
			continue
		}
		historySentence := BuildHistorySentence(historicalTrades[i], i, true)
		if historySentence != "" {
			historyStrings = append(historyStrings, BuildHistorySentence(historicalTrades[i], i, false))
		}
	}
	return strings.Join(historyStrings, "\n"), nil
}

// FormatTradeSummary formats a trade into a concise summary string
func FormatTradeSummary(trade history.TradeHistory) string {
	return fmt.Sprintf("Date: %s, Pair: %s, Action: %s, Price: %.4f, Size: %.2f, Leverage: %.2fx, PnL: %.2f",
		trade.Date.Format("2006-01-02 15:04:05"), trade.Pair, trade.Action, trade.Price, trade.Size, trade.Leverage, trade.PnlNet)
}
