package swaps

import (
	"HootCore/pairmaps"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"strings"
)

func GetStarterList(chain uint64, src bool) tgbotapi.InlineKeyboardMarkup {
	var starterList []pairmaps.TokenInfo
	switch chain {
	case 0:
		starterList = pairmaps.SwapStarterPolygon
	case 1:
		starterList = pairmaps.SwapStarterArbitrum
	default:
		return tgbotapi.InlineKeyboardMarkup{}
	}

	var keyboard [][]tgbotapi.InlineKeyboardButton
	var firstRow []tgbotapi.InlineKeyboardButton
	if src {
		firstRow = append(firstRow, tgbotapi.NewInlineKeyboardButtonData("Select your input token 🔽", "/srctokenlist"))
	} else {
		firstRow = append(firstRow, tgbotapi.NewInlineKeyboardButtonData("Select your output token 🔽", "/dsttokenlist"))
	}
	keyboard = append(keyboard, firstRow)
	for _, token := range starterList {
		row := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData(token.Ticker, "$"+strings.ToLower(token.Address)+"$"),
		}
		keyboard = append(keyboard, row)
	}

	lastRow := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Back 🔙", "/backswap"),
	}
	keyboard = append(keyboard, lastRow)

	return tgbotapi.NewInlineKeyboardMarkup(keyboard...)
}
