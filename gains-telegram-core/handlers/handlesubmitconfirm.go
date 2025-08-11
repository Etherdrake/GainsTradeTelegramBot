package handlers

import (
	"HootCore/alertatron"
	"HootCore/api/native"
	"HootCore/api/pairsstorage"
	"HootCore/boardbuilders/newtradebuilders"
	"HootCore/boardbuilders/stringbuilders"
	"HootCore/database"
	"HootCore/rdbhoot"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

func HandleSubmitConfirm(
	bot *tgbotapi.BotAPI,
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	update *tgbotapi.Update,
	boardMsgChatId int64,
	boardMsgMsgId int,
	guid int64) {

	board := newtradebuilders.BuildPerpMainBoardSubmitCancel(rdbHoot, guid)

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist.")
	}

	var collateralSymbol string
	switch trade.Collateral {
	case 1:
		collateralSymbol = "DAI"
	case 2:
		collateralSymbol = "USDC"
	case 3:
		collateralSymbol = "WETH"
	default:
		collateralSymbol = "SOL"
	}

	pubKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		log.Println("Error retrieving pubkey from the database.")
	}

	payload := native.GetGasBalancePayload{
		PublicKey:  pubKey,
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	// First we check if the wallet even has a gastoken
	_, gasBalanceFloat, err := native.GetGasBalance(payload)
	if err != nil {
		log.Println("Unable to get gasBalance for user: ", guid)
		return
	}

	if gasBalanceFloat < 0.01 {
		alertatron.SendAlert(bot, strconv.FormatInt(update.Message.Chat.ID, 10), "You do not have enough gas to place a trade!") // TODO: Change this into sending new gas
	}

	// Check if the user has not put his TP or SL incorrectly for being long or short
	switch trade.Buy {
	case true:
		if trade.SL > trade.OpenPrice {
			alertatron.SendAlert(bot, update.CallbackQuery.ID, "Your stoploss is higher than your open price, this is not possible on a long as it would cause an instant close of your trade.")
			return
		}
		if trade.TP < trade.OpenPrice {
			alertatron.SendAlert(bot, update.CallbackQuery.ID, "Your stoploss is higher than your open price, this is not possible on a long as it would cause an instant close of your trade.")
			return
		}

	case false:
		if trade.SL < trade.OpenPrice && trade.SL != 0 {
			alertatron.SendAlert(bot, update.CallbackQuery.ID, "Your stoploss is lower than your open price, this is not possible on a short as it would cause an instant close of your trade.")
			return
		}
	}

	minPosSizeUsd, apiErr := pairsstorage.GetPairMinPosSizeUsd(uint64(trade.PairIndex), trade.Chain)
	if apiErr != nil {
		log.Println("Error GetPairMinPosSizeUsd in submitconfirm: ", err)
	}

	// TODO: Set a special place for WETH by multiplying WETH x it's current dollar price
	totalPositionSize := trade.PositionSizeDai * float64(trade.Leverage)
	if totalPositionSize < float64(minPosSizeUsd) {
		alertatron.SendAlert(bot, update.CallbackQuery.ID, fmt.Sprintf(
			"Your total position size is too low, your position size is calculated by collateral x leverage. Your leverage: %.2f Your position size: %.2f Your total position size: %.2f Minimum position size: %d",
			trade.Leverage, trade.PositionSizeDai, totalPositionSize, minPosSizeUsd))
		return
	}

	//minLeverage, apiErr := pairsstorage.GetPairMinLev(uint64(trade.PairIndex), trade.Chain)
	//if apiErr != nil {
	//	log.Println("Error GetPairMinPosSizeUsd in submitconfirm: ", err)
	//}
	//
	//// TODO: Set a special place for WETH by multiplying WETH x it's current dollar price
	//if float32(minLeverage) > trade.Leverage {
	//	alertatron.SendAlert(bot, update.CallbackQuery.ID, fmt.Sprintf(
	//		"Your leverage is set too low. Minimum leverage on this pair %d\nYour leverage: %.2f",
	//		minLeverage, trade.Leverage))
	//	return
	//}

	maxLeverage, apiErr := pairsstorage.GetPairMaxLev(uint64(trade.PairIndex), trade.Chain)
	if apiErr != nil {
		log.Println("Error GetPairMinPosSizeUsd in submitconfirm: ", err)
	}

	// TODO: Set a special place for WETH by multiplying WETH x it's current dollar price
	if float32(maxLeverage) < trade.Leverage {
		alertatron.SendAlert(bot, update.CallbackQuery.ID, fmt.Sprintf(
			"Your leverage is set too high. Maximum leverage on this pair %d\nYour leverage: %.2f",
			maxLeverage, trade.Leverage))
		return
	}

	payloadToken := native.BalanceOfPayload{
		PublicKey:  pubKey,
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	tokenBalance, err := native.GetBalanceOf(payloadToken)
	if err != nil {
		log.Println("GetBalance Error: ", err)
	}

	tokenBalanceF64, parseErr := strconv.ParseFloat(tokenBalance, 64)
	if parseErr != nil {
		log.Printf("Error parsing tokenBalance for user %d tokenbalance: %f error: %s", guid, tokenBalanceF64, parseErr.Error())
	}

	if tokenBalanceF64 < trade.PositionSizeDai {
		alertatron.SendAlert(bot, update.CallbackQuery.ID, fmt.Sprintf("You do not have enough %s to place this trade. The amount needed for this trade is %.2f and you have: %.2f", collateralSymbol, trade.PositionSizeDai, tokenBalanceF64))
		return
	}

	// Update the message content with the desired message
	updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
	// Create a new editable message with the updated text
	msg := tgbotapi.NewEditMessageText(boardMsgChatId, boardMsgMsgId, updatedMessage)
	msg.ParseMode = tgbotapi.ModeMarkdown

	// Assign the new inline keyboard to the message
	msg.ReplyMarkup = &board
	_, errTg := bot.Send(msg)
	if errTg != nil {
		log.Println("Error sending messages in submitconfirm handler: ", errTg)
	}
}
