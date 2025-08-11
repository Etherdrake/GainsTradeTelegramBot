package historybuilders

import (
	"HootCore/history"
	"fmt"
	"log"
)

func BuildHistorySentence(hTrade history.TradeHistory, index int, isActive bool) string {
	// Determine if the PnL is positive or negative and set appropriate symbol and color
	var symbol, color string
	if hTrade.PnlNet >= 0 {
		symbol = "📈"
		color = "🟢"
	} else {
		symbol = "📉"
		color = "🔴"
	}

	switch hTrade.Action {
	case "TradeOpenedMarket":
		return ""
		//// Calculate percentage change
		//percentageChange := history.CalculateHistoryPnl(hTrade)
		//tradeAction := "MKT"
		//if isActive {
		//	sentence := fmt.Sprintf("/h%d %s %s %s\n", index, hTrade.Pair, symbol, tradeAction)
		//	return sentence
		//} else {
		//	sentence := fmt.Sprintf("/h%d %s %s %s", index, hTrade.Pair, symbol, tradeAction)
		//	return sentence
		//}
	case "TradeClosedMarket":
		tradeAction := "MKT"
		// Calculate percentage change
		percentageChange := history.CalculateHistoryPnl(hTrade)
		if isActive {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		} else {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		}
	case "TradeOpenedLimit":
		return ""
		//tradeAction := "LMT"
		//// Calculate percentage change
		//percentageChange := history.CalculateHistoryPnl(hTrade)
		//if isActive {
		//	sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
		//	return sentence
		//} else {
		//	sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
		//	return sentence
		//}
	case "TradeClosedLimit":
		tradeAction := "LMT"
		// Calculate percentage change
		percentageChange := history.CalculateHistoryPnl(hTrade)
		if isActive {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		} else {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		}
	case "TradeClosedTP":
		tradeAction := "TP"
		// Calculate percentage change
		percentageChange := history.CalculateHistoryPnl(hTrade)
		if isActive {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		} else {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		}
	case "TradeClosedSL":
		tradeAction := "SL"
		// Calculate percentage change
		percentageChange := history.CalculateHistoryPnl(hTrade)
		if isActive {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		} else {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, percentageChange, color, hTrade.PnlNet)
			return sentence
		}
	case "TradeClosedLIQ":
		tradeAction := "LIQ"
		// Calculate percentage change
		if isActive {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f\n", index, hTrade.Pair, symbol, tradeAction, "-100%", color, hTrade.PnlNet)
			return sentence
		} else {
			sentence := fmt.Sprintf("/h%d %s %s %s %s%% %s %0.2f", index, hTrade.Pair, symbol, tradeAction, "-100%", color, hTrade.PnlNet)
			return sentence
		}
	default:
		log.Println("Unknown trade action")
		return ""
	}
}
