package handlers

import (
	"HootCore/constants"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"strings"
	"time"
)

func HandleMarketString(
	bot *tgbotapi.BotAPI,
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	update *tgbotapi.Update,
	boardMsgChatId int64,
	boardMsgMsgId int,
	guid int64) {

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	priceReal, err := rdbHoot.GetPrice(int(trade.PairIndex))
	if err != nil {
		log.Println("Market to entry price error", err)
	}
	err = rdbHoot.SetEntryPrice(guid, priceReal)
	if err != nil {
		fmt.Println(err)
	}
	err = rdbHoot.SetOrderTypeToMarket(guid)

	switch trade.Buy {
	case true:
		err = rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 800))
		if err != nil {
			fmt.Println("SetTakeProfit error inside provide position size: ", err)
		}
		err = rdbHoot.SetStopLoss(guid, 0)
		if err != nil {
			fmt.Println("SetStopLoss error inside provide position size: ", err)
		}

	case false:
		err = rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 800))
		if err != nil {
			fmt.Println("SetTakeProfit error inside provide position size: ", err)
		}
		err = rdbHoot.SetStopLoss(guid, 0)
		if err != nil {
			fmt.Println("SetStopLoss error inside provide position size: ", err)
		}

	}

	responseText := "Entry price set to MARKET"
	msgSet := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	sleepMsg, err := bot.Send(msgSet)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
	time.Sleep(1500 * time.Millisecond)

	// Delete the initial ForceReply message
	utils.DeleteMessage(bot, update.Message.ReplyToMessage.Chat.ID, update.Message.ReplyToMessage.MessageID)
	// Delete the user entry
	utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
	// Delete ENTRY PRICE SET TO MARKET
	utils.DeleteMessage(bot, sleepMsg.Chat.ID, sleepMsg.MessageID)

	err = UpdateNewTradeOrActiveTrades(
		bot,
		client,
		rdbHoot,
		boardMsgChatId,
		boardMsgMsgId,
		guid)
	if err != nil {
		log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
	}
	return
}

func HandleCustomPrice(
	bot *tgbotapi.BotAPI,
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	update *tgbotapi.Update,
	boardMsgChatId int64,
	boardMsgMsgId int, guid int64) {

	// Refresh the board
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist.")
	}

	marketStr := "market"
	if update.Message.Text == marketStr || strings.ToLower(update.Message.Text) == marketStr {
		HandleMarketString(
			bot,
			client,
			rdbHoot,
			update,
			boardMsgChatId,
			boardMsgMsgId, guid)
		return
	}

	livePrice, errRdb := rdbHoot.GetPrice(int(trade.PairIndex)) // Assuming you have the GetPrice function implemented to fetch price from Redis
	if errRdb != nil {
		fmt.Println("Error fetching price from Redis in Stringbuilder:", errRdb)
	}

	entryPriceStr := update.Message.Text

	// We are going to see if it's an int
	entryPriceInt, entryPriceIntErr := strconv.Atoi(entryPriceStr)
	if entryPriceIntErr != nil {
		// We are going to use the float method
		entryPriceFloat, entryPriceFloatErr := strconv.ParseFloat(entryPriceStr, 64)
		if entryPriceFloatErr != nil {
			log.Println("entryPriceFloat Custom Error: ", entryPriceFloatErr)
			invalidMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid input! Hoot no comprendo.")
			send, err := bot.Send(invalidMsg)
			if err != nil {
				log.Println("Error with message Hoot! Hoot no comprendo")
			}
			time.Sleep(3000 * time.Millisecond)
			// Delete the no comprendo
			utils.DeleteMessage(bot, update.Message.ReplyToMessage.Chat.ID, send.MessageID)
			// Delete the initial ForceReply message
			utils.DeleteMessage(bot, update.Message.ReplyToMessage.Chat.ID, update.Message.ReplyToMessage.MessageID)
			// Delete the user entry
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			return
		}

		log.Printf("Entry price as float: %2.f\n", entryPriceFloat)

		switch trade.Buy {
		case true:
			// Entry price is a float
			err := rdbHoot.SetEntryPrice(guid, entryPriceFloat)
			if err != nil {
				fmt.Println(err)
			}
			err = rdbHoot.SetStopLoss(guid, 0.0)
			if err != nil {
				log.Println("SetStopLoss error inside provide position size: ", err)
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(entryPriceFloat, float64(trade.Leverage), 800))
			if err != nil {
				log.Println("SetTakeProfit error inside provide position size: ", err)
			}
			switch entryPriceFloat > livePrice {
			case true:
				cacheErr := rdbHoot.SetOrderTypeToStop(guid)
				if cacheErr != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			case false:
				cacheErr := rdbHoot.SetOrderTypeToLimit(guid)
				if cacheErr != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			}
		case false:
			// Entry price is a float
			err := rdbHoot.SetEntryPrice(guid, entryPriceFloat)
			if err != nil {
				fmt.Println(err)
			}
			err = rdbHoot.SetStopLoss(guid, 0.0)
			if err != nil {
				log.Println("SetStopLoss error inside provide position size: ", err)
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(entryPriceFloat, float64(trade.Leverage), 800))
			if err != nil {
				log.Println("SetTakeProfit error inside provide position size: ", err)
			}
			switch entryPriceFloat < livePrice {
			case true:
				cacheErr := rdbHoot.SetOrderTypeToStop(guid)
				if cacheErr != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			case false:
				cacheErr := rdbHoot.SetOrderTypeToLimit(guid)
				if cacheErr != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			}
		}
	} else {
		switch trade.Buy {
		case true:
			switch float64(entryPriceInt) > livePrice {
			case true:
				err := rdbHoot.SetOrderTypeToStop(guid)
				if err != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			case false:
				err := rdbHoot.SetOrderTypeToLimit(guid)
				if err != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			}
			// Entry price is an integer
			err := rdbHoot.SetEntryPrice(guid, float64(entryPriceInt))
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Entry price as integer: %d\n", entryPriceInt)
			err = rdbHoot.SetStopLoss(guid, 0.0)
			if err != nil {
				fmt.Println("SetStopLoss error inside provide position size: ", err)
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(float64(entryPriceInt), float64(trade.Leverage), 900))
			if err != nil {
				fmt.Println("SetTakeProfit error inside provide position size: ", err)
			}
		case false:
			switch float64(entryPriceInt) < livePrice {
			case true:
				err := rdbHoot.SetOrderTypeToStop(guid)
				if err != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			case false:
				err := rdbHoot.SetOrderTypeToLimit(guid)
				if err != nil {
					log.Println("SetOrderTypeToLimit error inside provide position size: ", err)
				}
			}
			// Entry price is an integer
			err := rdbHoot.SetEntryPrice(guid, float64(entryPriceInt))
			if err != nil {
				log.Println(err)
			}
			log.Printf("Entry price as integer: %d\n", entryPriceInt)
			err = rdbHoot.SetStopLoss(guid, 0.0)
			if err != nil {
				log.Println("SetStopLoss error inside provide position size: ", err)
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(float64(entryPriceInt), float64(trade.Leverage), 800))
			if err != nil {
				log.Println("SetTakeProfit error inside provide position size: ", err)
			}
		}
	}

	// Delete the old message "Please provide entry price
	msgToDelete := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.ReplyToMessage.MessageID)
	_, err := bot.Request(msgToDelete)
	if err != nil {
		fmt.Println(err)
	}

	// Delete the old reply to "Please provide entry price
	replyToDelete := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
	_, err = bot.Request(replyToDelete)
	if err != nil {
		fmt.Println(err)
	}

	responseText := "Entry price set"
	msgSet := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	sleepMsg, err := bot.Send(msgSet)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
	time.Sleep(constants.UserInputDelayActiveTradesBoard)

	// Delete the old message
	msgToDelete = tgbotapi.NewDeleteMessage(sleepMsg.Chat.ID, sleepMsg.MessageID)
	_, err = bot.Request(msgToDelete)
	if err != nil {
		fmt.Println(err)
	}

	err = UpdateNewTradeOrActiveTrades(
		bot,
		client,
		rdbHoot,
		boardMsgChatId,
		boardMsgMsgId,
		guid)
	if err != nil {
		log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
	}
	return
}
