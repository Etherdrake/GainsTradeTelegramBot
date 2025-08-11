package boardbuilders

import (
	"HootCore/rdbhoot"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"strconv"
)

func BuildPairsBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
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

func BuildForexBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	user, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		fmt.Println("User not found in cache")
		return tgbotapi.InlineKeyboardMarkup{}
	}

	user.PairPage += 1

	pairPage := strconv.Itoa(int(user.PairPage))

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️", "/decrpairpagefx"),
			tgbotapi.NewInlineKeyboardButtonData("PAGE "+pairPage, "/leverage"),
			tgbotapi.NewInlineKeyboardButtonData("➡️", "/incrpairpagefx"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Crypto", "/pairs"),
			tgbotapi.NewInlineKeyboardButtonData("Commodities", "/commoditypairs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔎 Search", "/search"),
			tgbotapi.NewInlineKeyboardButtonData("🔙 Back", "/leverage"),
		),
	)
}

func BuildCommoditiesBoard(rdbHoot *rdbhoot.HootRedisClient, guid int64) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️", "/commoditypairs"),
			tgbotapi.NewInlineKeyboardButtonData("PAGE 1", "/leverage"),
			tgbotapi.NewInlineKeyboardButtonData("➡️", "/commoditypairs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Forex", "/forexpairs"),
			tgbotapi.NewInlineKeyboardButtonData("Crypto", "/pairs"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔙 Back", "/leverage"),
		),
	)
}
