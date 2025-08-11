package swapbuilders

import (
	"HootCore/rdbhoot"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"strconv"
)

func BuildSwapPairsBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	pairPage := strconv.Itoa(int(user.PairPage + 1))

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️", "/decrpairpage"),
			tgbotapi.NewInlineKeyboardButtonData("PAGE "+pairPage, "/leverage"),
			tgbotapi.NewInlineKeyboardButtonData("➡️", "/incrpairpage"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Forex", "/forexpairs"),
			tgbotapi.NewInlineKeyboardButtonData("Commodities", "/commoditypairs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔎 Search", "/search"),
			tgbotapi.NewInlineKeyboardButtonData("🔙 Back", "/leverage"),
		),
	)
}
