package main

import (
	"HootCore/alertatron"
	"HootCore/altseasontools"
	"HootCore/api/native"
	"HootCore/api/pairsstorage"
	"HootCore/api/testnet"
	"HootCore/api/tradinginteractions"
	"HootCore/boardbuilders"
	"HootCore/boardbuilders/historybuilders"
	"HootCore/boardbuilders/newtradebuilders"
	"HootCore/boardbuilders/stringbuilders"
	"HootCore/boardbuilders/swapbuilders"
	"HootCore/boardbuilders/tpslbuilders"
	"HootCore/chartservice"
	"HootCore/concurrentmaps"
	"HootCore/constants"
	"HootCore/database"
	"HootCore/errorhandling"
	"HootCore/formatters"
	"HootCore/handlers"
	"HootCore/history"
	"HootCore/mongolisten"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/rdbhoot/rdblocker"
	"HootCore/rdbhoot/rdbstart"
	"HootCore/regexes"
	"HootCore/search"
	"HootCore/swaps"
	"HootCore/tg_utils"
	"HootCore/utils"
	"HootCore/wallet"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"strings"
	"time"
)

func handleUpdates(bot *tgbotapi.BotAPI, client *mongo.Client, update tgbotapi.Update, boardCache *concurrentmaps.BoardMessagesCache, rdbHoot *rdbhoot.HootRedisClient) {
	guidFirst := update.SentFrom().ID
	pressedStart, errCheck := rdbstart.CheckUserID(ctx, rdbHoot, guidFirst)
	if errCheck != nil {
		log.Println("Error checking UserID in RdbLocker: ", guidFirst)
	}
	if !pressedStart {
		// First we check if this is not the /start message
		if update.CallbackQuery == nil && update.Message.Text == "/start" {
			log.Println("User pressed start for the first time.")
		} else {
			if update.CallbackQuery != nil {
				alertMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Please send the /start command again to re-initialize your session.")
				_, err := bot.Send(alertMsg)
				if err != nil {
					log.Println("Error sending alert message: ", err)
				}
				return
			} else {
				alertMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please send the /start command again to re-initialize your session.")
				_, err := bot.Send(alertMsg)
				if err != nil {
					log.Println("Error sending alert message: ", err)
				}
				return
			}
		}
	}
	if update.CallbackQuery != nil {
		// Handle callback queries here
		callbackData := update.CallbackQuery.Data
		log.Printf("Callbackdata: %s", callbackData)
		//log.Println("ActivePerp: ", rdbHoot.Cache[5938441997].ActivePerp)
		boardCache.Set(update.CallbackQuery.From.ID, update.CallbackQuery.Message)
		var msg tgbotapi.Chattable
		guid := update.CallbackQuery.From.ID
		chatID := update.CallbackQuery.Message.Chat.ID
		boardMsg, ok := boardCache.Get(update.CallbackQuery.From.ID)
		if !ok {
			log.Println("Error retrieving BoardMessageNew")
		}
		switch callbackData {
		case "/srctokenlist": // TODO: This should update on click
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, boardMsg.Text)
			board := swaps.GetStarterList(trade.Chain, true)
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
			_, _ = bot.Send(msg)
			return
		case "/dsttokenlist": // TODO: This should update on click
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, boardMsg.Text)
			board := swaps.GetStarterList(trade.Chain, false)
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
			_, _ = bot.Send(msg)
			return
		case "/decrswapprice":
			return
		case "/incrswapprice":
			return
		case "/submitswap":
			return
		case "/search":
			searchMsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Please reply to this message with your query: ")
			searchMsg.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply:            true,
				InputFieldPlaceholder: "Type the symbol or part of the symbol for which you want to search.",
				Selective:             false,
			}
			_, _ = bot.Send(searchMsg)
			return
		case "/leveragefromchart":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
			board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
			// Update the message content with the desired message
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			perpMenuString := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			// Create a new editable message with the updated text
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, perpMenuString)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msg.ReplyMarkup = board
			msg.ParseMode = tgbotapi.ModeMarkdown

			spoofMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Loading...")
			spoofMsg.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: false,
				Selective:  true,
			}
			spoofDel, _ := bot.Send(spoofMsg)
			utils.DeleteMessage(bot, spoofDel.Chat.ID, spoofDel.MessageID)
			return
		case "/leverage":
			board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
			// Update the message content with the desired message
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			// Update the message content with the desired message
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			// Create a new editable message with the updated text
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msg = msgEdit
		//
		case "/leveragechart":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
			handlers.LeverageChartSender(rdbHoot, "1M", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart: ", err)
			}
			return
		case "/leveragechart5M":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, boardMsg.MessageID)
			handlers.LeverageChartSender(rdbHoot, "5M", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart5M: ", err)
			}
			return
		case "/leveragechart15M":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, boardMsg.MessageID)
			handlers.LeverageChartSender(rdbHoot, "15M", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart15M: ", err)
			}
			return
		case "/leveragechart1H":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, boardMsg.MessageID)
			handlers.LeverageChartSender(rdbHoot, "1H", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart1H: ", err)
			}
			return
		case "/leveragechart4H":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, boardMsg.MessageID)
			handlers.LeverageChartSender(rdbHoot, "4H", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart4H: ", err)
			}
			return
		case "/leveragechart1D":
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, boardMsg.MessageID)
			handlers.LeverageChartSender(rdbHoot, "1D", bot, guid, update.CallbackQuery.Message.Chat.ID)
			err := rdbHoot.SetActiveCharts(guid, true)
			if err != nil {
				log.Println("Error in /leveragechart1D: ", err)
			}
			return
		case "/leverageback":
			err := rdbHoot.SetActiveCharts(guid, false)
			if err != nil {
				log.Println("Error in SetActiveCharts /leverageback: ", err)
			}
			utils.DeleteMessage(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
			// Update the message content with the desired message
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			switch trade.ActiveWindow {
			case 0:
				board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
				perpMenuString := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
				// Create a new editable message with the updated text
				msgNew := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, perpMenuString)
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgNew.ReplyMarkup = board
				msgNew.ParseMode = tgbotapi.ModeMarkdown
				_, err := bot.Send(msgNew)
				if err != nil {
					log.Println(err)
				}
			case 1:
				board := boardbuilders.BuildActiveTradeBoardGNSV2(client, rdbHoot, guid, trade.Chain)
				perpMenuString := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
				// Create a new editable message with the updated text
				msgNew := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, perpMenuString)
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgNew.ReplyMarkup = board
				msgNew.ParseMode = tgbotapi.ModeMarkdown
				_, err := bot.Send(msgNew)
				if err != nil {
					log.Println(err)
				}
			}

			return
			// TODO: /leveragechart not implemented for active trades
		case "/sendchart":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			b64, err := chartservice.GetChart1M(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println("/sendchart GetChart1M error:  ", err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard1M(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}

			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/sendchart5M":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist. This")
			}

			b64, err := chartservice.GetChart5M(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println(err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard5M(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
			return
		case "/sendchart15M":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			b64, err := chartservice.GetChart15M(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println(err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard15M(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
			return
		case "/sendchart1H":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			b64, err := chartservice.GetChart1H(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println(err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard1H(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
			return
		case "/sendchart4H":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			b64, err := chartservice.GetChart4H(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println(err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard4H(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
			return
		case "/sendchart1D":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			b64, err := chartservice.GetChart1D(pairmaps.IndexToPair[int(trade.PairIndex)])
			if err != nil {
				log.Println(err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, boardMsg.MessageID)

			err = chartservice.SendChartWithKeyboard1D(bot, update.CallbackQuery.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
			return
		case "/newtrade":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist: ", guid)
				return
			}

			if trade.ActiveWindow == 0 {
				alertatron.SendAlert(bot, update.CallbackQuery.ID, alertatron.NewTradeDouble)
				return
			}

			err := rdbHoot.SetActiveWindow(guid, 0)
			if err != nil {
				log.Println("Error setting ActiveWindow in /newtrade")
			}

			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/newtradecancel":
			_, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist: ", guid)
				return
			}
			err := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return

		case "/activetrades":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
				return
			}
			isTrue, err := CheckIfUserHasOpenTradesOrOrdersGNS(client, guid, trade.Chain)
			if err != nil {
				log.Println("Error Checking OpenTradesOrOrders: ", err)
				return
			}
			if isTrue {
				cacheErr := rdbHoot.SetActiveWindow(guid, 1)
				if cacheErr != nil {
					log.Println("Error setting active window in active trades.")
					return
				}
				errUpdate := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID,
					boardMsg.MessageID, guid)
				if errUpdate != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
					return
				}
			} else {
				alertatron.SendAlert(bot, update.CallbackQuery.ID, "You do not have any open trades or orders.")
				return
			}
			return
		case "/updateopenlimitprice": // TODO: Check if this exists or not
			// Load the active order and update the price on-chain
			return
		case "/resettrade":
			rdbErr := rdbHoot.ResetTrade(client, guid)
			if rdbErr != nil {
				return
			}
			err := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID,
				guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
			if cacheErr != nil {
				log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
			}
		case "/market":
			rdbErr := rdbHoot.SetOrderTypeToMarket(guid)
			if rdbErr != nil {
				return
			}
			err := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID,
				guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
			if cacheErr != nil {
				log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
			}
			return
			//
		case "/submitconfirm":
			handlers.HandleSubmitConfirm(bot, client, rdbHoot, &update, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			return
		case "/submit":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.", guid)
			}
			trades, dbErr := mongolisten.PullOpenTrades(client, trade.Trader, trade.Chain)
			if dbErr != nil {
				log.Println("Error PullOpenTrades inside submit: ", dbErr)
				return
			}
			board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			finalMsg, err := utils.CreateEditMessage(board, updatedMessage, boardMsg.Chat.ID, boardMsg.MessageID, "Markdown")
			if err != nil {
				log.Println("Error CreatingEditMessage message inside /submit: ", err)
				return
			}
			_, err = bot.Send(finalMsg)
			if err != nil {
				log.Println("Error sending message inside /submit: ", err)
				return
			}
			go func() {
				pair := pairmaps.IndexToPair[int(trade.PairIndex)]

				currentPrice, err := rdbHoot.GetPrice(int(trade.PairIndex))
				if err != nil {
					log.Println("Market to entry price error", err)
				}

				var longShort string
				if trade.Buy {
					longShort = "LONG"
				} else {
					longShort = "SHORT"
				}

				currentPriceUS, _ := formatters.FormatFinancialUS(currentPrice)

				msgNewStr := fmt.Sprintf("Executing trade:\n *%.2fx %s %s @ %s*", trade.Leverage, longShort, pair, currentPriceUS)
				msgNew := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, msgNewStr)
				msgNew.ParseMode = tgbotapi.ModeMarkdown
				msgFirst, err := bot.Send(msgNew)
				if err != nil {
					log.Println("Error sending add message in /submit: ", err)
					return
				}

				submitErr := handlers.HandleSubmit(bot, client, rdbHoot, guid, boardMsg.Chat.ID, *boardMsg, boardMsg.MessageID)
				if submitErr != nil {
					log.Println("Error handling submit: ", err)
					return
				} else {
					cacheErr := rdbHoot.SetActiveWindow(guid, 1)
					if cacheErr != nil {
						log.Println("Error setting active window in active trades.")
						return
					}
					errUpdate := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID,
						boardMsg.MessageID, guid)
					if errUpdate != nil {
						log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
						return
					}
				}

				posSize, err := formatters.FormatFinancialUS(float64(trade.Leverage) * trade.PositionSizeDai)
				if err != nil {
					log.Println(err)
				}
				var msgChattable tgbotapi.Chattable
				msgTwoStr := fmt.Sprintf("Trade executed *successfully.*\n Position size *%s*", posSize)
				msgChattable = tgbotapi.NewEditMessageText(msgFirst.Chat.ID, msgFirst.MessageID, msgTwoStr)
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgEdit, ok := msgChattable.(tgbotapi.EditMessageTextConfig)
				if !ok {
					log.Println("Cannot convert msg to type EditMessageTextConfig")
					return
				}

				msgEdit.ParseMode = tgbotapi.ModeMarkdown
				msgChattable = msgEdit

				_, errTg := bot.Send(msgChattable)
				if errTg != nil {
					log.Println("Error sending message: ", err)
				}

				if len(trades) <= 1 {
					cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
					if cacheErr != nil {
						log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
					}
				}
				go utils.DeleteMessageWithTimer(bot, chatID, msgEdit.MessageID, 15)
				return
			}()
			return
		case "/wallet":
			isSet, err := database.CheckIfSet(client, guid)
			if err != nil {
				fmt.Println(err)
				return
			}
			if isSet {
				pubkey, err := database.GetPublicKey(client, guid)
				if err != nil {
					log.Println("Error retrieving public key in wallet", err)
					return
				}
				updatedMessage, err := stringbuilders.BuildSimpleWalletMainString(client, guid)
				if err != nil {
					fmt.Println("Error building string for users: " + strconv.Itoa(int(guid)))
					fmt.Println(err)
				}
				// Create a new editable message with the updated text
				msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
				qrFileName, err := wallet.CreateQR(pubkey)
				if err != nil {
					fmt.Println(err)
				}
				// Send the QR code as a photo
				photoMedia := tg_utils.NewInputMediaPhotoFromFile(qrFileName)

				photoMsg := tgbotapi.NewPhoto(guid, photoMedia.BaseInputMedia.Media)
				_, err = bot.Send(photoMsg)
				if err != nil {
					fmt.Println("Photo error")
				}

				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
				if !ok {
					log.Println("Cannot convert msg to type EditMessageTextConfig")
					return
				}
				// Assign the new inline keyboard to the message
				msgEdit.ReplyMarkup = &boardbuilders.WalletMainBoard
				msg = msgEdit

			} else {
				pubkey, pk := wallet.GenerateWallet()
				pubKeyErr := database.AddPublicKey(client, guid, pubkey)
				if pubKeyErr != nil {
					return
				}
				privKeyErr := database.AddPrivateKey(client, guid, pk)
				if privKeyErr != nil {
					return
				}

				updatedMessage, err := stringbuilders.BuildSimpleWalletMainString(client, guid)
				if err != nil {
					fmt.Println("Error building string for users: " + strconv.Itoa(int(guid)))
					fmt.Println(err)
				}
				// Create a new editable message with the updated text
				msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
				if !ok {
					log.Println("Cannot convert msg to type EditMessageTextConfig")
					return
				}
				// Assign the new inline keyboard to the message
				msgEdit.ReplyMarkup = &boardbuilders.WalletMainBoard
				msg = msgEdit

			}
		// #TODO: Swap interactions for the swap boards
		case "/backswap":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			if trade.PairPage < 0 {
				trade.PairPage = 0
			}

			board := swapbuilders.BuildSwapBoardV1(rdbHoot, guid)
			updatedMessage := swapbuilders.BuildSwapPairString(trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		// Perpmainboard Interactions
		case "/pairs":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			if trade.PairPage < 0 {
				trade.PairPage = 0
			}

			board := boardbuilders.BuildPairsBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildCryptoPairString(rdbHoot, int(trade.PairPage))
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/refresh":
			err := handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID,
				guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
			if cacheErr != nil {
				log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
			}
			return
		case "/stoploss":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/takeprofit":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/toggletolong":
			err := rdbHoot.ToggleLongShort(guid)
			if err != nil {
				log.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			price, err := rdbHoot.GetPrice(int(trade.PairIndex))
			if err != nil {
				log.Println("Error getting price")
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 800))
			if err != nil {
				fmt.Println(err)
			}
			err = rdbHoot.SetStopLoss(guid, 0)
			if err != nil {
				fmt.Println(err)
			}
			if trade.OpenPrice < price && trade.OrderType != 1 {
				err := rdbHoot.SetOrderTypeToLimit(guid)
				if err != nil {
					log.Println("User does not exist.")
				}
			}
			if trade.OpenPrice > price && trade.OrderType != 2 {
				err := rdbHoot.SetOrderTypeToStop(guid)
				if err != nil {
					log.Println("User does not exist.")
				}
			}

			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/toggletoshort":
			err := rdbHoot.ToggleLongShort(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			price, err := rdbHoot.GetPrice(int(trade.PairIndex))
			if err != nil {
				log.Println("Error getting price")
			}
			err = rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 800))
			if err != nil {
				log.Println(err)
			}
			if err != nil {
				log.Println(err)
			}
			if trade.OrderType != 0 {
				if !trade.Buy {
					if trade.OpenPrice > price && trade.OrderType != 1 {
						err := rdbHoot.SetOrderTypeToLimit(guid)
						if err != nil {
							log.Println("User does not exist.")
						}
					}
					if trade.OpenPrice < price && trade.OrderType != 2 {
						err := rdbHoot.SetOrderTypeToStop(guid)
						if err != nil {
							log.Println("User does not exist.")
						}
					}
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/decrprice":
			err := rdbHoot.DecrementEntryPrice(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			if trade.OpenPrice <= 0 {
				alertatron.SendAlert(bot, update.CallbackQuery.ID, "Entryprice is lower than or equal to zero! Open price has been reset to current live-price. Happy trading sir :-)")
				price, _ := rdbHoot.GetPrice(int(trade.PairIndex))
				err := rdbHoot.SetEntryPrice(guid, price)
				if err != nil {
					log.Println("Error encountered during SetEntryPrice")
				}
			}
			price, err := rdbHoot.GetPrice(int(trade.PairIndex))
			if err != nil {
				log.Printf("Error fetching price for pairIndex %d: %v", trade.PairIndex, err)
				// Handle this error: perhaps return or set a default price value
			}

			if trade.Buy {
				if trade.OpenPrice < price && trade.OrderType != 1 {
					err := rdbHoot.SetOrderTypeToLimit(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
				if trade.OpenPrice > price && trade.OrderType != 2 {
					err := rdbHoot.SetOrderTypeToStop(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
			} else {
				if trade.OpenPrice > price && trade.OrderType != 1 {
					err := rdbHoot.SetOrderTypeToLimit(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
				if trade.OpenPrice < price && trade.OrderType != 2 {
					err := rdbHoot.SetOrderTypeToStop(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/incrprice":
			err := rdbHoot.IncrementEntryPrice(guid)
			if err != nil {
				log.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			price, err := rdbHoot.GetPrice(int(trade.PairIndex))
			if err != nil {
				log.Printf("Error fetching price for pairIndex %d: %v", trade.PairIndex, err)
				// Handle this error: perhaps return or set a default price value
			}
			if trade.Buy {
				if trade.OpenPrice < price && trade.OrderType != 1 {
					err := rdbHoot.SetOrderTypeToLimit(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
				if trade.OpenPrice > price && trade.OrderType != 2 {
					err := rdbHoot.SetOrderTypeToStop(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
			} else {
				if trade.OpenPrice > price && trade.OrderType != 1 {
					err := rdbHoot.SetOrderTypeToLimit(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
				if trade.OpenPrice < price && trade.OrderType != 2 {
					err := rdbHoot.SetOrderTypeToStop(guid)
					if err != nil {
						log.Println("User does not exist.")
					}
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/customprice":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide your order entry price, type \"market\" for market orders or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
			_, err := bot.Send(msg)
			if err != nil {
				log.Println("This is the customprice: ", err)
			}
			return
		case "/decrpossize":
			err := rdbHoot.DecrementPositionSize(guid)
			if err != nil {
				fmt.Println(err)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/custompossize":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide collateral size or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
			// Optional: use ForceReply to show reply interface to the user
			//msg = tgbotapi.ForceReply{ForceReply: true, Selective: true}
		case "/incrpossize":
			err := rdbHoot.IncrementPositionSize(guid)
			if err != nil {
				fmt.Println(err)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/decrleverage":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			// TODO: Pull this from the API instead of hardcoding -> This exists already in Redis
			minLev := float32(2)

			if minLev == trade.Leverage-1 {
				answerTxt := fmt.Sprintf("Leverage for this pair can not be set lower than: %.2f", minLev)
				alertatron.SendAlert(bot, update.CallbackQuery.ID, answerTxt)
				return
			}

			err := rdbHoot.DecrementLeverage(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/customleverage":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide the leverage you want or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
		case "/incrleverage":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			// TODO: Pull this from the API instead of hardcoding
			maxLev := float32(150)

			if maxLev == trade.Leverage+1 {
				answerTxt := fmt.Sprintf("Leverage for this pair can not be set higher than: %.2f", maxLev)
				alertatron.SendAlert(bot, update.CallbackQuery.ID, answerTxt)
				return
			}
			err := rdbHoot.IncrementLeverage(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "/closetrade":
			go handlers.CloseTradeHandler(bot, client, rdbHoot, boardMsg.Chat.ID, guid)
			return
		case "/cancelorder":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			// Check
			perpCode := trade.ActivePerp.IsTradeOrOrder
			if perpCode != 2 {
				log.Println("Cancel order called but ActivePerpetual is not set as trade")
				return
			}
			txReceipt, err := tradinginteractions.CancelOrderGNS(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
			if err != nil {
				log.Println("Error closing trade: ", err)
				return
			}
			var msgTx tgbotapi.MessageConfig
			switch trade.Chain {
			case 137:
				msgTx = tgbotapi.NewMessage(chatID, "Order cancelled successfully: [TxReceipt](https://polygonscan.com/tx/"+strings.Trim(txReceipt, "\"")+")")
			case 42161:
				msgTx = tgbotapi.NewMessage(chatID, "Order cancelled successfully: [TxReceipt](https://arbiscan.io/tx/"+strings.Trim(txReceipt, "\"")+")")
			case 421614:
				msgTx = tgbotapi.NewMessage(chatID, "Order cancelled successfully: [TxReceipt](https://sepolia.arbiscan.io/tx/"+strings.Trim(txReceipt, "\"")+")")
			}
			cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
			if cacheErr != nil {
				log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
			}
			if errRemove := rdblocker.RemoveUserID(ctx, rdbHoot, guid); errRemove != nil {
				log.Printf("Error removing UserID %d from Redis: %v", guid, errRemove)
			}
			msgTx.ParseMode = tgbotapi.ModeMarkdown
			_, err = bot.Send(msgTx)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
			return
		case "/openchart":
			return
			// Take Profit Queries
		case "/zerotp":
			err := rdbHoot.SetTakeProfit(guid, 0)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msg = msgEdit
		case "/decrtp":
			err := rdbHoot.DecrementTakeProfit(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/customtp":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide take profit or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
			// Take Profit Queries
		case "/incrtp":
			err := rdbHoot.IncrementTakeProfit(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/decrpairpage":
			err := rdbHoot.DecrementPage(guid)
			if err != nil {
				log.Println("Error decrementing pair page for user: ", guid)
			}
			user, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := boardbuilders.BuildPairsBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildCryptoPairString(rdbHoot, int(user.PairPage))

			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)

			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/incrpairpage":
			err := rdbHoot.IncrementPage(guid, "crypto")
			if err != nil {
				log.Println("Error incrementing pair page for user: ", guid)
			}
			user, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := boardbuilders.BuildPairsBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildCryptoPairString(rdbHoot, int(user.PairPage))

			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)

			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/decrpairpagefx":
			err := rdbHoot.DecrementPage(guid)
			if err != nil {
				log.Println("Error decrementing pair page for user: ", guid)
			}
			user, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := boardbuilders.BuildForexBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildFxPairString(rdbHoot, int(user.PairPage))

			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)

			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/incrpairpagefx":
			err := rdbHoot.IncrementPage(guid, "forex")
			if err != nil {
				log.Println("Error incrementing pair page for user: ", guid)
			}
			user, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := boardbuilders.BuildForexBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildFxPairString(rdbHoot, int(user.PairPage))

			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)

			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus25":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 25))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			} else {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 25))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus50":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 50))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			} else {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 50))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus100":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 100))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			} else {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 100))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus150":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 200))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			} else {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 200))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus900":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 900))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			} else {
				err := rdbHoot.SetTakeProfit(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 900))
				if err != nil {
					log.Println("Take profit failed to set at /plus25", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/zerotpedit":
			err := rdbHoot.SetActivePerpEditSl(guid, 0)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/decrtpedit":
			err := rdbHoot.DecrementActivePerpEditTp(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/customtpedit":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide the new take-profit for this open position or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
			// Take Profit Queries
		case "/incrtpedit":
			err := rdbHoot.IncrementActivePerpEditTp(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus25edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
					return
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error plus25edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newTakeProfit := utils.CalculateTakeProfit(openPrice, leverage, 25)

			err := rdbHoot.SetActivePerpEditTp(guid, newTakeProfit)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus50edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newTakeProfit := utils.CalculateTakeProfit(openPrice, leverage, 50)

			err := rdbHoot.SetActivePerpEditTp(guid, newTakeProfit)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			if err != nil {
				log.Println("Could not GetActiveOpenTrade in /zerotpedit", err)
			}
			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus100edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newTakeProfit := utils.CalculateTakeProfit(openPrice, leverage, 100)

			err := rdbHoot.SetActivePerpEditTp(guid, newTakeProfit)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus150edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newTakeProfit := utils.CalculateTakeProfit(openPrice, leverage, 150)

			err := rdbHoot.SetActivePerpEditTp(guid, newTakeProfit)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/plus900edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newTakeProfit := utils.CalculateTakeProfit(openPrice, leverage, 900)

			err := rdbHoot.SetActivePerpEditTp(guid, newTakeProfit)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/decrsl":
			err := rdbHoot.DecrementStopLoss(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/customsl":
			msgConfig := tgbotapi.NewMessage(boardMsg.Chat.ID, "Please provide stop loss or /cancel")
			msgConfig.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			msg = msgConfig
		case "/incrsl":
			err := rdbHoot.IncrementStopLoss(guid)
			if err != nil {
				fmt.Println(err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/zerosl":
			err := rdbHoot.SetStopLoss(guid, 0)
			if err != nil {
				log.Println("Stop loss failed to set at /plus300", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msg = msgEdit
		case "/minus10":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 10))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
				fmt.Println("BUY")
			} else {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 10))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
				fmt.Println("SELL")
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus25":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 25))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			} else {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 25))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus33":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 33))
				if err != nil {
					log.Println("Stop loss failed to set at /minus33", err)
				}
			} else {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 33))
				if err != nil {
					log.Println("Stop loss failed to set at /minus33", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus50":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 50))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			} else {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 50))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus75":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if exists {
				log.Println("User does not exist.")
			}
			if trade.Buy {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 75))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			} else {
				err := rdbHoot.SetStopLoss(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 75))
				if err != nil {
					log.Println("Stop loss failed to set at /plus300", err)
				}
			}
			trade, exists = rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/zerosledit":
			err := rdbHoot.SetActivePerpEditSl(guid, 0)
			if err != nil {
				log.Println("Stop loss failed to set at /plus300", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msg = msgEdit
		case "/minus10edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newStopLoss := utils.CalculateStopLoss(openPrice, leverage, 10)

			err := rdbHoot.SetActivePerpEditSl(guid, newStopLoss)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = "MarkdownV2"
			msg = msgEdit
		case "/minus25edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = openOrder.Trade.Leverage
			}
			newStopLoss := utils.CalculateStopLoss(openPrice, leverage, 33)
			err := rdbHoot.SetActivePerpEditSl(guid, newStopLoss)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus33edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus33edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus33edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newStopLoss := utils.CalculateStopLoss(openPrice, leverage, 33)

			err := rdbHoot.SetActivePerpEditSl(guid, newStopLoss)
			if err != nil {
				log.Println("Take profit failed to set at /minus33edit", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus50edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newStopLoss := utils.CalculateStopLoss(openPrice, leverage, 50)

			err := rdbHoot.SetActivePerpEditSl(guid, newStopLoss)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/minus75edit":
			perpCacheCode, cacheErrGet := rdbHoot.GetIsTrade(guid)
			if cacheErrGet != nil {
				return
			}
			var openPrice float64
			var leverage float64
			switch perpCacheCode {
			case 0:
				return
			case 1:
				openTrade, cacheErr := rdbHoot.GetSelectedTrade(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openTrade.Trade.OpenPrice
				leverage = float64(openTrade.Trade.Leverage)
			case 2:
				openOrder, cacheErr := rdbHoot.GetSelectedOrder(guid)
				if cacheErr != nil {
					log.Println("Error minus75edit")
				}
				openPrice = openOrder.Trade.OpenPrice
				leverage = float64(openOrder.Trade.Leverage)
			}

			newStopLoss := utils.CalculateStopLoss(openPrice, leverage, 75)

			err := rdbHoot.SetActivePerpEditSl(guid, newStopLoss)
			if err != nil {
				log.Println("Take profit failed to set at /plus25", err)
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/savesledit":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			switch trade.ActivePerp.IsTradeOrOrder {
			case 0:
				log.Println("No ActivePerp could be found in /savesledit.")
				return
			case 1:
				_, err := tradinginteractions.UpdateSlGns(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.ActivePerp.ActiveTradeOrOrder.Trade.Sl, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
				if err != nil {
					log.Println("Error in UpdateSlGns /savesledit", err)
				}
				return
			case 2:
				log.Println(trade.ActivePerp.ActiveTradeOrOrder)
				err := tradinginteractions.UpdateOpenLimitOrder(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.ActivePerp.ActiveTradeOrOrder.Trade.OpenPrice, trade.ActivePerp.ActiveTradeOrOrder.Trade.Sl, trade.ActivePerp.ActiveTradeOrOrder.Trade.Tp, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
				if err != nil {
					log.Println("Error in UpdateOpenLimitOrder /savesledit", err)
				}
				return
			default:
				log.Println("No ActivePerp could be found in /savesledit. No int was even set!")
				return
			}
		case "/savetpedit":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}
			switch trade.ActivePerp.IsTradeOrOrder {
			case 0:
				log.Println("No ActivePerp could be found in /savetpedit.")
				return
			case 1:
				txReceipt, err := tradinginteractions.UpdateTpGns(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.ActivePerp.ActiveTradeOrOrder.Trade.Sl, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
				if err != nil {
					log.Println("Error in UpdateTpGns /savetpedit", err)
				} else {
					log.Println("Successfully updated TP: ", txReceipt)
				}
				return
			case 2:
				err := tradinginteractions.UpdateOpenLimitOrder(guid, trade.ActivePerp.ActiveTradeOrOrder.Trade.Index, trade.ActivePerp.ActiveTradeOrOrder.Trade.OpenPrice, trade.ActivePerp.ActiveTradeOrOrder.Trade.Tp, trade.ActivePerp.ActiveTradeOrOrder.Trade.Sl, trade.Chain, trade.ActivePerp.ActiveTradeOrOrder.Trade.CollateralIndex)
				if err != nil {
					log.Println("Error in UpdateOpenLimitOrder /savetpedit", err)
				}
				return
			default:
				log.Println("No ActivePerp could be found in /savesledit. No int was even set!")
				return
			}
		case "/forexpairs":
			err := rdbHoot.ResetPage(guid)
			if err != nil {
				log.Println("Error resetting page.")
			}

			board := boardbuilders.BuildForexBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildFxPairString(rdbHoot, 1)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/cryptopairs":
			err := rdbHoot.ResetPage(guid)
			if err != nil {
				log.Println("Error resetting page.")
			}
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := boardbuilders.BuildPairsBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildCryptoPairString(rdbHoot, int(trade.PairPage))
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}

			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		case "/commoditypairs":
			err := rdbHoot.ResetPage(guid)
			if err != nil {
				log.Println("Error resetting page.")
			}
			//trade, exists := rdbHoot.GetCoreCache(guid)
			//if !exists {
			//	log.Println("User does not exist.")
			//}
			board := boardbuilders.BuildCommoditiesBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildCommoditiesPairString(rdbHoot)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		}

		if strings.Contains(callbackData, "/sledit") {
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildSlEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		}
		if strings.Contains(callbackData, "/tpedit") {
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			board := tpslbuilders.BuildTpEditBoard(rdbHoot, guid)
			updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
			msg = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msgEdit, ok := msg.(tgbotapi.EditMessageTextConfig)
			if !ok {
				log.Println("Cannot convert msg to type EditMessageTextConfig")
				return
			}
			msgEdit.LinkPreviewOptions.IsDisabled = true
			// Assign the new inline keyboard to the message
			msgEdit.ReplyMarkup = &board
			msgEdit.ParseMode = tgbotapi.ModeMarkdown
			msg = msgEdit
		}

		// Adding in the logic for the swaps UI => 13/03/24
		// when the tokenAddress is within dollar signs the users wants to change the srcToken
		// when the tokenAddress is within two ampersand symbols the user wants to change the dstToken
		if update.CallbackQuery.Data[0] == '$' && update.CallbackQuery.Data[len(update.CallbackQuery.Data)-1] == '$' {
			// We have to extract the text between the dollar signs
			extractedAddress := update.CallbackQuery.Data[1 : len(update.CallbackQuery.Data)-1]

			// We have to check what our chainId is
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			// We have to retrieve the old swap information
			cachedSwapInfo, err := swaps.GetSwapConfig(ctx, rdbHoot, strconv.FormatInt(guid, 10))
			if err != nil {
				log.Println("Error in GetSwapConfig: ", err)
			}

			log.Println("CachedSwapInfo: ", cachedSwapInfo)

			var chainId int
			var tokenInfo *swaps.TokenInfo
			var swapErr error

			// We have to retrieve the chainInfo depending on which chain the user is on
			switch trade.Chain {
			case 1:
				chainId = 42161
				tokenInfo, swapErr = swaps.GetTokenInfo(chainId, extractedAddress)
				if swapErr != nil {
					log.Println("Error in GetTokenInfo: ", err)
				}
			case 0:
				chainId = 137
				tokenInfo, swapErr = swaps.GetTokenInfo(chainId, extractedAddress)
				if swapErr != nil {
					log.Println("Error in GetTokenInfo: ", err)
				}
			default:
				log.Println("Chain not found for chain: ", trade.Chain)
				return
			}

			if tokenInfo != nil {
				if len(tokenInfo.Contracts) > 0 {
					// We set this address as the new swapAddress with the corresponding information
					newSwapInfo := swaps.SwapInfo{
						// We only change the source because the token we received was in dollar signs
						SrcSymbol:        tokenInfo.Symbol,
						SrcToken:         strings.ToLower(tokenInfo.Contracts[0].ContractAddress),
						SrcTokenDecimals: "18",   // TODO: Make sure this is correct
						SrcTokenAmt:      "1000", // TODO: This is hardcoded
						DstSymbol:        cachedSwapInfo.Chains[trade.Chain].DstSymbol,
						DstToken:         cachedSwapInfo.Chains[trade.Chain].DstToken,
						DstTokenDecimals: cachedSwapInfo.Chains[trade.Chain].DstTokenDecimals,
					}
					infoErr := swaps.EditSwapInfo(ctx, rdbHoot, strconv.FormatInt(guid, 10), trade.Chain, newSwapInfo)
					if err != nil {
						log.Println("Error in EditSwapInfo: ", infoErr)
						return
					}
					// Send the new swapboard
					handlers.HandleSwap(rdbHoot, bot, client, update, guid)
					return
				} else {
					log.Println("Tokeninfo not found for chain, here is the full struct: ", tokenInfo)
					return
				}
			} else {
				log.Printf("swaps.GetTokenInfo returned a nil pointer for chain: %d and CA: %s", chainId, extractedAddress)
				return
			}
		}
		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
			errorhandling.HandleTelegramError(bot, err, update.CallbackQuery.ID)
		}
	} else if update.Message != nil && update.Message.ReplyToMessage == nil {
		guid := update.Message.From.ID
		if update.Message.Text != "/start" {
			_, exists := rdbHoot.GetCoreCache(update.Message.From.ID)
			if !exists {
				alertMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please send the /start command again to re-initialize your session.")
				_, err := bot.Send(alertMsg)
				if err != nil {
					log.Println("Error sending alert message: ", err)
				}
				return
			}
		}
		if update.Message.Text == "/reset" {
			_, exists := rdbHoot.GetCoreCache(update.Message.From.ID)
			if !exists {
				log.Println("User does not exist:", guid)
				return
			}
			err := rdbHoot.ResetUser(client, guid)
			if err != nil {
				log.Println("Error in ResetUser: ", err)
			}
			rdbHoot.InitUser(client, guid)
			log.Println("User reset: ", guid)
			return
		}

		if update.Message.Text == "/resetcache" {
			err := rdbHoot.ResetUser(client, update.Message.From.ID)
			if err != nil {
				log.Println("Error in ResetUser: ", err)
				return
			}
			rdbHoot.InitUser(client, guid)
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			return
		}

		var msg tgbotapi.MessageConfig
		boardMsg, ok := boardCache.Get(update.Message.From.ID)
		if !ok {
			log.Println("Message does not exist")
		}

		// Regexes Logic
		if regexes.HistoryPattern.MatchString(update.Message.Text) {
			match := regexes.HistoryPattern.FindStringSubmatch(update.Message.Text)
			if len(match) > 1 {
				number, err := strconv.Atoi(match[1])
				if err != nil {
					log.Println(err)
				}

				pubKey, err := database.GetPublicKey(client, guid)
				if err != nil {
					log.Println("Error getting publickey inside regex matching for /h[0-9]: ", err)
					return
				}

				historicalTrades, err := history.GetHistoricalTrades(pubKey)
				if err != nil {
					log.Printf("error getting historical trades: %v", err)
					return
				}

				historyString, err := historybuilders.BuildHistoryString(historicalTrades, guid, 0, number)
				if err != nil {
					return
				}
				board := historybuilders.BuildHistoryBoard(historicalTrades[number].Tx)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, historyString)
				msg.ReplyMarkup = board
				msg.ParseMode = tgbotapi.ModeMarkdown
			} else {
				log.Println("Did not match any known pattern")
			}
		}

		if regexes.TradePattern.MatchString(update.Message.Text) {
			match := regexes.TradePattern.FindStringSubmatch(update.Message.Text)
			if len(match) > 1 {
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				idx, err := strconv.Atoi(match[1])
				if err != nil {
					log.Println(err)
				}
				trade, exists := rdbHoot.GetCoreCache(guid)
				if !exists {
					log.Println("User does not exist.")
				}
				publicKey, err := database.GetPublicKey(client, guid)
				if err != nil {
					log.Println("Error getting PublicKey: ", err)
				}
				hasTrades, err := mongolisten.HasOpenTrades(client, publicKey, trade.Chain)
				if err != nil {
					log.Println("Error Checking OpenTradesOrOrders: ", err)
					return
				}
				var msgChattable tgbotapi.Chattable
				if hasTrades {
					openTrades, err := mongolisten.GetAllOpenTrades(client, publicKey, trade.Chain)
					if err != nil {
						log.Println("GetAllTraders error: ", err)
					}

					openTradesSorted := altseasontools.SortTrades(openTrades)

					if err != nil {
						log.Println("ExtractNumber error in /t handling", err)
					}

					for idx > len(openTradesSorted) {
						idx -= 1
					}

					// BEST USE METHOD
					// BEST USE METHOD
					errSelect := rdbHoot.SetActiveSelectedTradeOrOrder(guid, openTradesSorted[idx-1].IntoInternal())
					if errSelect != nil {
						fmt.Println("Unable to SetActiveSelectedTradeOrOrder in /t")
					}

					updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
					board := boardbuilders.BuildActiveTradeBoardGNSV2(client, rdbHoot, guid, trade.Chain)

					msgChattable = tgbotapi.NewEditMessageText(boardMsg.Chat.ID, boardMsg.MessageID, updatedMessage)
					// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
					msgEdit, ok := msgChattable.(tgbotapi.EditMessageTextConfig)
					if !ok {
						log.Println("Cannot convert msg to type EditMessageTextConfig")
						return
					}
					// Assign the new inline keyboard to the message
					msgEdit.ReplyMarkup = &board
					msgEdit.ParseMode = tgbotapi.ModeMarkdown
					msgChattable = msgEdit

					_, errTg := bot.Send(msgChattable)
					if errTg != nil {
						log.Println("Error sending message: ", err)
					}
					return
				} else {
					log.Println("No trade has been found!")
					return
				}
			}
		}

		if regexes.OrderPattern.MatchString(update.Message.Text) {
			match := regexes.OrderPattern.FindStringSubmatch(update.Message.Text)
			if len(match) > 1 {
				// First we delete
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				log.Println("RegExes: ", match)
				index, err := strconv.Atoi(match[1])
				if err != nil {
					log.Println(err)
				}
				trade, exists := rdbHoot.GetCoreCache(guid)
				if !exists {
					log.Println("User does not exist.")
				}
				publicKey, err := database.GetPublicKey(client, guid)
				if err != nil {
					log.Println(err)
				}
				hasOrders, err := mongolisten.HasOpenOrders(client, publicKey, trade.Chain)
				if err != nil {
					//TODO: Tell the user no user has been found for him specifically
					log.Println("Error Checking OpenTradesOrOrders: ", err)
					return
				}
				if hasOrders {
					// Delete the message containing the textChars
					utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)

					activeOrders, err := mongolisten.GetOpenOrders(client, publicKey, trade.Chain)
					if err != nil {
						log.Println("GetOpenOrders error: ", err)
					}

					sortedOrders := altseasontools.SortTrades(activeOrders)

					for index > len(sortedOrders) {
						index -= 1
					}

					// We decrease by one here, so we do not get a length error
					errSetActive := rdbHoot.SetActiveSelectedOrder(guid, sortedOrders[index-1].IntoInternal())
					if errSetActive != nil {
						log.Println("Error setting ActiveSelectedOrder: ", errSetActive)
					}

					updatedMessage := stringbuilders.BuildActiveGainsTradeStringV2(client, rdbHoot, guid, trade.Chain)
					board := boardbuilders.BuildActiveTradeBoardGNSV2(client, rdbHoot, guid, trade.Chain)
					finalMsg, editErr := utils.CreateEditMessage(board, updatedMessage, boardMsg.Chat.ID, boardMsg.MessageID, "Markdown")
					if editErr != nil {
						log.Println("Error in CreateEditMessage in case polygon")
					}
					_, tgErr := bot.Send(finalMsg)
					if tgErr != nil {
						log.Println("Error sending message:", err)
						return
					}
					return
				}
			}
		}

		switch update.Message.Text {
		//case "/feestest":
		//	board := boardbuilders.BuildFeesBoard(tradeCache, rdbHoot, guid)
		//	// Update the message content with the desired message
		//	trade, exists := rdbHoot.GetCoreCache(guid)
		//	if !exists {
		//		log.Println("User does not exist /leverage.")
		//		return
		//	}
		//	perpMenuString := stringbuilders.BuildFeesString(client, rdbHoot, guid, tradeCache, trade.Chain)
		//	// Create a new editable message with the updated text
		//	msg = tgbotapi.NewMessage(update.Message.Chat.ID, perpMenuString)
		//	// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
		//	msg.ReplyMarkup = board
		//	msg.ParseMode = tgbotapi.ModeMarkdown
		//
		//	//msgToDelete := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
		//	//
		//	//_, err := bot.Request(msgToDelete)
		//	//if err != nil {
		//	//	errorhandling.HandleTelegramError(bot, err, update.CallbackQuery.ID)
		//	//}
		case "/swap":
			handlers.HandleSwap(rdbHoot, bot, client, update, guid)
			return
		case "/history":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist:", trade.ID)
			}
			pubKey, err := database.GetPublicKey(client, guid)
			if err != nil {
				log.Println("Error getting public key inside /history")
				return
			}

			historicalTrades, err := history.GetHistoricalTrades(pubKey)
			if err != nil {
				log.Printf("error getting historical trades: %v", err)
				return
			}

			if len(historicalTrades) > 0 {
				board := historybuilders.BuildHistoryBoard(historicalTrades[0].Tx)
				historyString, _ := historybuilders.BuildHistoryString(historicalTrades, trade.ID, 0, 0)
				// Create a new editable message with the updated text
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, historyString)
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msg.ReplyMarkup = board
				msg.ParseMode = tgbotapi.ModeMarkdown
				_, tgErr := bot.Send(msg)
				if tgErr != nil {
					log.Println(tgErr)
				}
				return
			} else {
				return
			}
		//case "/secretalert":
		//	//testStr := "400000000"
		//	//tester, _ := bigtools.ScaleBigStrToFloatString(testStr, 10)
		//	//fmt.Println(tester)
		//
		//	config, err := altseasoncore.GetGainsTradeContext(rdbHoot, 421614)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//
		//	testPosSzie := 100000.0000
		//	dynSpread, dynErr := altseasoncore.GetDynamicSpread(0, 1, testPosSzie, true, config.ToConverted())
		//	if dynErr != nil {
		//		log.Println(dynErr)
		//	}
		//
		//	log.Println("dynamicSpread: ", dynSpread)
		//
		//	kara, err := altseasoncore.GetFixedSpreadPercentage(0, config.ToConverted())
		//	if err != nil {
		//		log.Println(err)
		//	}
		//
		//	log.Println("FixedSpread", kara)
		case "/start":
			photoFilePath := "./assets/logo_wide.jpg"
			photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoFilePath))
			photoMsg.Caption = StartMessage
			photoMsg.ParseMode = tgbotapi.ModeMarkdown
			//msg.ReplyMarkup = keyboards.GetMainMenuKeyboard()

			userInMongo, err := database.CheckUserDB(client, guid)
			if err != nil {
				log.Println("Error with CheckuserDB", err)
			}
			if !userInMongo {
				log.Println("User is not in Mongo")
				initErr := database.InitUser(client, guid, update.Message.From.FirstName, update.Message.From.LastName)
				if initErr != nil {
					log.Println("Error: InitUser Database ", guid)
				}
				isSet, setErr := database.CheckIfSet(client, update.Message.From.ID)
				if setErr != nil {
					log.Println("Error checking if set.")
					return
				}
				if !isSet {
					pubkey, pk := wallet.GenerateWallet()
					errPub := database.AddPublicKey(client, guid, strings.ToLower(pubkey))
					if errPub != nil {
						log.Println("Error creating publickey in /start")
						return
					}
					errPriv := database.AddPrivateKey(client, guid, strings.ToLower(pk))
					if errPriv != nil {
						log.Println("Error creating privatekey in /start")
						return
					}
				}
			}
			_, exists := rdbHoot.GetCoreCache(guid) // Use tradeCache's Get function to check for existence
			if !exists {
				// The user does not exist inside the cache yet
				rdbErr := rdbHoot.InitUser(client, guid)
				if rdbErr != nil {
					log.Println("Error initializing user: ", err)
					return
				}

				initExpErr := rdbHoot.InitExperienceTimer(ctx, guid)
				if initExpErr != nil {
					log.Println("Error initializing ExperienceTimer: ", rdbErr)
				}

				// We check here if the experience counter exists for the user
				xpExists, xpErr := rdbHoot.CheckExperienceExists(ctx, guid)
				if xpErr != nil {
					log.Println("Error in CheckExperienceExists: ", err)
				}
				if !xpExists {
					xpRdbErr := rdbHoot.SetExperience(ctx, guid, 0)
					if xpRdbErr != nil {
						log.Println("Error in SetExperience: ", err)
					}
				}

				timerExists, expTimerErr := rdbHoot.CheckExperienceTimerExists(ctx, guid)
				if expTimerErr != nil {
					log.Println("Error in CheckExperienceTimerExists inside of /start: ", err)
				}

				if !timerExists {
					initExpErr = rdbHoot.InitExperienceTimer(ctx, guid)
					if initExpErr != nil {
						log.Println("Error initializing ExperienceTimer: ", rdbErr)
					}
				}

				// TODO: Move this stuff to /leverage with the database check
				msgNewStr := fmt.Sprintf("🎉 Hey fren, you got *$10,000 DAI* and earned *10,000 HoiSin!*")
				msgNew := tgbotapi.NewMessage(update.Message.Chat.ID, msgNewStr)
				_, tgErr := bot.Send(msgNew)
				if tgErr != nil {
					log.Println("Error sending add message in /start: ", err)
					return
				}

				// Add userID to Redis
				errAdd := rdbstart.AddUserID(ctx, rdbHoot, guidFirst)
				if errAdd != nil {
					log.Printf("Error add UserID %d to Redis: %v", guidFirst, errAdd)
				}

				cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
				if cacheErr != nil {
					log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
				}

				// Do a check whether the user has TestETH TODO: This part must be removed once the beta is different
				pubKey, dbErr := database.GetPublicKey(client, guid)
				if dbErr != nil {
					log.Println("Error getting public key inside /start")
				}

				payload := native.GetGasBalancePayload{
					PublicKey:  pubKey,
					Collateral: 1,
					Chain:      421614,
				}

				_, balanceEthF64, balErr := native.GetGasBalance(payload)
				if balErr != nil {
					log.Println("Get gas balance error: ", balErr)
				}

				// If the balance is smaller we send testEth again
				if balanceEthF64 < 0.005 {
					apiErr := testnet.GetEthSepolia(guid)
					if apiErr != nil {
						log.Println("Something went wrong GetEthSepolia in /start: ", apiErr)
					}
				}

				_, tgErr = bot.Send(photoMsg)
				if tgErr != nil {
					log.Println("Error sending start message in /start: ", err)
				}
				return
			} else {
				rdbErr := rdbHoot.SetActiveWindow(guid, 0)
				if rdbErr != nil {
					log.Println("Error setting ActiveWindow")
				}

				cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
				if cacheErr != nil {
					log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
				}

				_, err = bot.Send(photoMsg)
				if err != nil {
					log.Println("Error sending start message in /start: ", err)
				}
			}
		case "/gas":
			apiErr := testnet.GetEthSepolia(guid)
			if apiErr != nil {
				log.Println("Something went wrong GetEthSepolia in /start: ", apiErr)
			}
		case "/add":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("Error getting trade for guid: ", guid)
			}

			// Delete the add message
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)

			timer, err := rdbHoot.GetExperienceTimer(ctx, guid)
			if err != nil {
				log.Println("Error CheckExperienceTimerExpired: ", err)
				return
			}

			if timer != 0 {
				currentTime := time.Now().Unix()
				timeDifference := timer + 12*3600 - currentTime

				hoursLeft := timeDifference / 3600
				minutesLeft := (timeDifference % 3600) / 60

				if hoursLeft < 1 {
					msgNew := tgbotapi.NewMessage(update.Message.Chat.ID, strconv.FormatInt(minutesLeft, 10)+" more minutes until you can add again!")
					msgTimeLeft, tgErr := bot.Send(msgNew)
					if tgErr != nil {
						log.Println("Error sending add message in /add: ", tgErr)
						return
					}
					go utils.DeleteMessageWithTimer(bot, msgTimeLeft.Chat.ID, msgTimeLeft.MessageID, 15)
					return
				} else {
					msgNew := tgbotapi.NewMessage(update.Message.Chat.ID, strconv.FormatInt(hoursLeft, 10)+" more hours and "+strconv.FormatInt(minutesLeft, 10)+" minutes left until you can add again!")
					msgTimeLeft, tgErr := bot.Send(msgNew)
					if tgErr != nil {
						log.Println("Error sending add message in /add: ", tgErr)
						return
					}
					go utils.DeleteMessageWithTimer(bot, msgTimeLeft.Chat.ID, msgTimeLeft.MessageID, 15)
					return
				}
			}

			msgNew := tgbotapi.NewMessage(update.Message.Chat.ID, "Adding $10,000 DAI to your balance......")
			msgFirst, err := bot.Send(msgNew)
			if err != nil {
				log.Println("Error sending add message in /add: ", err)
				return
			}

			_, err = testnet.GetFreeDai(guid, trade.Chain)
			if err != nil {
				log.Println("Error with getdai", err)
				return
			} else {
				rdbErr := rdbHoot.IncrementExperience(ctx, guid, 10000)
				if rdbErr != nil {
					log.Println("Error incrementing experience inside of /add: ", rdbErr)
					return
				}

				rdbErr = rdbHoot.SetExperienceTimer(ctx, guid, time.Now().Unix())
				if rdbErr != nil {
					log.Println("Error setting experience timer inside /add: ", rdbErr)
					return
				}

				var msgChattable tgbotapi.Chattable
				msgChattable = tgbotapi.NewEditMessageText(msgFirst.Chat.ID, msgFirst.MessageID, ", you got $10,000 DAI and earned 10,000 Hoot coins!")
				// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
				msgEdit, ok := msgChattable.(tgbotapi.EditMessageTextConfig)
				if !ok {
					log.Println("Cannot convert msg to type EditMessageTextConfig")
					return
				}

				msgEdit.ParseMode = tgbotapi.ModeMarkdown
				msgChattable = msgEdit

				msgTwo, errTg := bot.Send(msgChattable)
				if errTg != nil {
					log.Println("Error sending message: ", err)
				}

				utils.DeleteMessageWithTimer(bot, msgTwo.Chat.ID, msgTwo.MessageID, 5)

				err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
				if err != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		case "/cancel":
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID-1)
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			return
		case "/secretallowance":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("Error in secretallowance")
			}
			diamond, err := native.ApproveCollateralToDiamond(client, guid, *trade)
			if err != nil {
				log.Println("Error in secretallowance: ", err)
				return
			}
			log.Println("Secret allowance successfully: ", diamond)
		case "/chart":
			b64, err := chartservice.GetChart1M("BTC/USD")
			if err != nil {
				log.Println(err)
			}
			err = chartservice.SendChart(bot, update.Message.Chat.ID, b64)
			if err != nil {
				log.Println(err)
			}
		case "/leverage":
			err := rdbHoot.SetActiveWindow(guid, 0)
			if err != nil {
				log.Println("Error setting ActiveWindow")
			}
			board := newtradebuilders.BuildNewTradeBoard(rdbHoot, guid)
			// Update the message content with the desired message
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist /leverage.")
				return
			}
			perpMenuString := stringbuilders.BuildNewTradeString(client, rdbHoot, guid, trade.PairIndex)
			// Create a new editable message with the updated text
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, perpMenuString)
			// This type assertion is needed because tgbotapi.NewEditMessageText returns a tgbotapi.MessageConfig, not a tgbotapi.EditMessageTextConfig.
			msg.ReplyMarkup = board
			msg.ParseMode = tgbotapi.ModeMarkdown

			msgToDelete := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)

			_, err = bot.Request(msgToDelete)
			if err != nil {
				errorhandling.HandleTelegramError(bot, err, update.CallbackQuery.ID)
			}

			spoofMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Loading...")
			spoofMsg.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: false,
				Selective:  true,
			}
			spoofDel, err := bot.Send(spoofMsg)
			if err != nil {
				log.Println("/leverage error: ", err)
				errorhandling.HandleTelegramError(bot, err, update.CallbackQuery.ID)
			}

			utils.DeleteMessage(bot, spoofDel.Chat.ID, spoofDel.MessageID)
			log.Println(trade.Collateral)
			approveExecuted, txID := DiamondApprovalCheck(client, rdbHoot, guid)
			log.Println(approveExecuted, txID)
			//err := tg_utils.DeleteMessage(bot, msgToDelete)
			//if err != nil {
			//	log.Println("Error deleting message: ", err)
			//}
		case "/sniper":
			log.Println("sniper")
		case "/wallet":
			isSet, err := database.CheckIfSet(client, guid)
			if err != nil {
				fmt.Println(err)
				return
			}
			if isSet {
				pubkey, err := database.GetPublicKey(client, guid)
				if err != nil {
					log.Println("Error retrieving public key in wallet", err)
					return
				}
				updatedMessage, err := stringbuilders.BuildSimpleWalletMainString(client, guid)
				if err != nil {
					fmt.Println("Error building string for users: " + strconv.Itoa(int(guid)))
					fmt.Println(err)
				}
				// Create a new editable message with the updated text
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, updatedMessage)
				msg.ParseMode = tgbotapi.ModeMarkdown
				qrFileName, createQrErr := wallet.CreateQRWithLogo(pubkey)
				if createQrErr != nil {
					log.Println("Error with CreateQRWithLogo: ", createQrErr)
				}
				// Send the QR code as a photo
				photoMedia := tg_utils.NewInputMediaPhotoFromFile(qrFileName)

				photoMsg := tgbotapi.NewPhoto(guid, photoMedia.BaseInputMedia.Media)
				_, tgErr := bot.Send(photoMsg)
				if tgErr != nil {
					log.Println("Failed to send QR code photo:", err)
					return
				}

				_, tgErr = bot.Send(msg)
				if err != nil {
					log.Println("Failed to send wallet builded:", err)
					return
				}
				return
			} else {
				pubkey, pk := wallet.GenerateWallet()
				dbErr := database.AddPublicKey(client, guid, pubkey)
				if dbErr != nil {
					log.Println(dbErr)
					return
				}
				dbErr = database.AddPrivateKey(client, guid, pk)
				if dbErr != nil {
					log.Println(dbErr)
					return
				}

				updatedMessage, err := stringbuilders.BuildSimpleWalletMainString(client, guid)
				if err != nil {
					fmt.Println("Error building string for users: " + strconv.Itoa(int(guid)))
					fmt.Println(err)
				}
				// Create a new editable message with the updated text
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, updatedMessage)
			}
		// Handle Chain switches
		case "/polygon":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetChain(guid, 137)
			cacheErr := rdbHoot.SetStartTradeOrOrder(client, guid)
			if cacheErr != nil {
				log.Println("Could not SetStartTradeOrOrder for guid: ", guid)
			}
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/practice":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetChain(guid, 421614) // Sepolia
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/arbitrum":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetChain(guid, 42161)
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/native":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetCollateralToDai(guid)
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/dai":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetCollateralToDai(guid)
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}

			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/weth":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetCollateralToWeth(guid)
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}

			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				log.Println("User does not exist.")
			}

			fmt.Println("Coll", trade.Collateral)
			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
		case "/usdc":
			// First delete the messsage received
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			// Switch the cache to a different chain
			err := rdbHoot.SetCollateralToUsdc(guid)
			if err != nil {
				log.Println("Error setting chain for user: ", guid)
			}

			err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			// REALKEYBOARD HANDLING
		}
		// Paircatcher Regex
		if regexes.CatcherPattern.MatchString(update.Message.Text) {
			match := regexes.CatcherPattern.FindStringSubmatch(update.Message.Text)
			log.Println(match)
			if len(match) >= 1 {

				pair := strings.Replace(match[0], "x", "/", -1)
				pair = strings.Trim(pair, "/")
				index := pairmaps.PairToIndex[pair]

				// Set the new pair in the cache
				err := rdbHoot.SetPairIndex(int64(index), guid)
				if err != nil {
					log.Println("Error setting PairIndex: ", err)
				}

				// Retrieve the price of the new asset from Redis
				price, err := rdbHoot.GetPrice(index)
				if err != nil {
					log.Println("Error retrieving the new price of the asset from Redis: ", err)
				}

				leverage, err := rdbHoot.GetLeverage(guid)
				if err != nil {
					log.Println("Error retrieving leverage from tradeCache: ", err)
				}

				isLong, err := rdbHoot.IsLong(guid)
				if err != nil {
					log.Println("Error retrieving isLong from tradeCache: ", err)
				}

				// Update the open-price in the cache
				err = rdbHoot.SetEntryPrice(guid, price)
				if err != nil {
					log.Println("Error setting the entry price: ", err)
				}
				// Update the TP
				tp := utils.CalculateTakeProfit(price, float64(leverage), 900)
				err = rdbHoot.SetTakeProfit(guid, tp)
				if err != nil {
					log.Println("Error setting take-profit", err)
				}

				// Update the SL
				sl := 0.00
				err = rdbHoot.SetStopLoss(guid, sl)
				if err != nil {
					log.Println("Error setting stop-loss", err)
				}

				// Update the liq
				liq := utils.CalculateLiquidationPrice(price, float64(leverage), isLong)
				err = rdbHoot.SetLiquidation(guid, liq)
				if err != nil {
					log.Println("Error setting liquidation: ", err)
				}

				err = handlers.UpdateNewTradeOrActiveTrades(bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
				if err != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades inside pair catcher: ", err)
				}
			} else {
				log.Println("Paircatcher found a match but not a matchedString: ", update.Message.Text)
			}
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			return
		}
		if msg.Text != "" {
			boardMsg, err := bot.Send(msg)
			if err != nil {
				log.Println("Error sending message Update.Message.Text: ", err)
				errorhandling.HandleTelegramError(bot, err, strconv.FormatInt(update.Message.Chat.ID, 10))
			} else {
				boardCache.Set(update.Message.From.ID, &boardMsg)
			}
		}

	} else if update.Message.ReplyToMessage != nil {
		guid := update.Message.From.ID

		boardMsg, ok := boardCache.Get(update.Message.From.ID)
		if !ok {
			log.Println("Message does not exist")
		}

		var responseText string // Declare a string to store the response message
		switch update.Message.ReplyToMessage.Text {
		case "Please reply to this message with your query:":
			query := update.Message.Text
			resultArr := search.SearchPairs(query)
			board := search.BuildSearchBoard()
			msgTxt := search.BuildSearchString(resultArr)

			finalMsg, err := utils.CreateEditMessage(board, msgTxt, boardMsg.Chat.ID, boardMsg.MessageID, "Markdown")
			if err != nil {
				log.Println("Error CreatingEditMessage message in query reply search: ", err)
			}

			utils.DeleteMessage(bot, boardMsg.Chat.ID, update.Message.ReplyToMessage.MessageID)
			utils.DeleteMessage(bot, boardMsg.Chat.ID, update.Message.MessageID)

			_, err = bot.Send(finalMsg)
			if err != nil {
				log.Println("Error with message Hoot! Hoot no comprendo")
			}

			return
		// customprice handle
		case "Please provide your order entry price, type \"market\" for market orders or /cancel":
			handlers.HandleCustomPrice(bot, client, rdbHoot, &update, boardMsg.Chat.ID, boardMsg.MessageID, guid)
		case "Please provide collateral size or /cancel":
			positionSize := update.Message.Text
			positionSizeFloat, err := strconv.ParseFloat(positionSize, 64)
			if err != nil {
				fmt.Println("positionSize Custom Error: ", err)
				responseText = "Invalid position size, please check your input! Your input was: " + update.Message.Text
			} else {
				err = rdbHoot.SetPositionSize(guid, positionSizeFloat)
				trade, exists := rdbHoot.GetCoreCache(guid)
				if !exists {
					fmt.Println("User not found in cache")
				}
				err = rdbHoot.SetStopLoss(guid, utils.CalculateStopLoss(trade.OpenPrice, float64(trade.Leverage), 10))
				if err != nil {
					fmt.Println("SetStopLoss error inside provide position size: ", err)
				}
				err = rdbHoot.SetTakeProfit(guid, utils.CalculateTakeProfit(trade.OpenPrice, float64(trade.Leverage), 900))
				if err != nil {
					fmt.Println("SetTakeProfit error inside provide position size: ", err)
				}
				if err != nil {
					fmt.Println(err)
				}
				responseText = "Position size set!"

				msgBoard := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				respMsgDel, err := bot.Send(msgBoard)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				time.Sleep(constants.UserInputDelayActiveTradesBoard)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
				updateErr := handlers.UpdateNewTradeOrActiveTrades(
					bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
				if updateErr != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		case "Please provide the leverage you want or /cancel":
			trade, exists := rdbHoot.GetCoreCache(guid)
			if !exists {
				fmt.Println("User not found in cache")
			}

			leverageFactor := update.Message.Text

			// Extract all numeric sequences from the input string
			numericSequences := regexes.LeveragePattern.FindAllString(leverageFactor, -1)

			// Concatenate all numeric sequences
			concatenatedNumber := strings.Join(numericSequences, "")

			leverageFactorInt, err := strconv.Atoi(concatenatedNumber)
			if err != nil {
				log.Println("positionSize Custom Error: ", err)
				responseText = "Invalid leverage size, please check your input! Your input was: " + update.Message.Text
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				msgError := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				msgErrDel, err := bot.Send(msgError)
				if err != nil {
					log.Println(msgError)
				}
				time.Sleep(2000 * time.Second)
				utils.DeleteMessage(bot, msgErrDel.Chat.ID, msgErrDel.MessageID)
				return
			}

			//
			minLev, apiErr := pairsstorage.GetPairMinLev(uint64(trade.PairIndex), trade.Chain)
			if apiErr != nil {
				log.Println("There is an error: ", apiErr)
				return
			}

			maxLev, apiErr := pairsstorage.GetPairMaxLev(uint64(trade.PairIndex), trade.Chain)
			if apiErr != nil {
				log.Println("There is an error: ", apiErr)
				return
			}

			if uint64(leverageFactorInt) < minLev {
				msgStr := fmt.Sprintf("Your leverage is too low, the minimum leverage for this asset is: %d", minLev)
				msgWarning := tgbotapi.NewMessage(update.Message.Chat.ID, msgStr)
				_, err = bot.Send(msgWarning)
				if err != nil {
					log.Println("Error with message Hoot! Hoot no comprendo")
				}
				return

			}

			if uint64(leverageFactorInt) > maxLev {
				msgStr := fmt.Sprintf("Your leverage is too high, the maximum leverage for this asset is: %d", maxLev)
				msgWarning := tgbotapi.NewMessage(update.Message.Chat.ID, msgStr)
				_, err = bot.Send(msgWarning)
				if err != nil {
					log.Println("Error with message Hoot! Hoot no comprendo")
				}
				return

			}

			// Update leverage
			err = rdbHoot.SetLeverage(guid, float32(leverageFactorInt))
			if err != nil {
				fmt.Println(err)
			}
			price, err := rdbHoot.GetPrice(int(trade.PairIndex))
			if err != nil {
				log.Println("Error getting price.")
			}
			switch trade.Buy {
			case true:
				// Calculate take profit
				takeProfit := utils.CalculateTakeProfit(price, float64(leverageFactorInt), 800)
				// Update take profit
				err = rdbHoot.SetTakeProfit(guid, takeProfit)
				if err != nil {
					fmt.Println(err)
				}
				// Update stop loss
				err = rdbHoot.SetStopLoss(guid, 0)
				if err != nil {
					fmt.Println(err)
				}
				// Calculate liquidation
				liquidation := utils.CalculateLiquidationPrice(price, float64(leverageFactorInt), trade.Buy)
				// Update liquidation
				err = rdbHoot.SetLiquidation(guid, liquidation)
				if err != nil {
					fmt.Println(err)
				}
			case false:
				// Calculate take profit
				takeProfit := utils.CalculateStopLoss(price, float64(leverageFactorInt), 800)
				// Update take profit
				err = rdbHoot.SetTakeProfit(guid, takeProfit)
				if err != nil {
					fmt.Println(err)
				}
				// Update stop loss
				err = rdbHoot.SetStopLoss(guid, 0)
				if err != nil {
					fmt.Println(err)
				}
				// Calculate liquidation
				liquidation := utils.CalculateLiquidationPrice(price, float64(leverageFactorInt), trade.Buy)
				// Update liquidation
				err = rdbHoot.SetLiquidation(guid, liquidation)
				if err != nil {
					fmt.Println(err)
				}
			}
			responseText = "Leverage set!"
			setMsg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
			respMsgDel, err := bot.Send(setMsg)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
			time.Sleep(1500 * time.Millisecond)
			utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
			utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
			utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
			updateErr := handlers.UpdateNewTradeOrActiveTrades(
				bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
			if updateErr != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		case "Please provide stop loss or /cancel":
			positionSize := update.Message.Text
			positionSizeInt, err := strconv.Atoi(positionSize)
			if err != nil {
				fmt.Println("stopLoss Custom Error: ", err)
				responseText = "Invalid stop loss, please check your input! Your input was: " + update.Message.Text
			} else {
				err = rdbHoot.SetStopLoss(guid, float64(positionSizeInt))
				if err != nil {
					fmt.Println(err)
				}
				responseText = "Stop loss set!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				respMsgDel, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				time.Sleep(500 * time.Millisecond)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
				updateErr := handlers.UpdateNewTradeOrActiveTrades(
					bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
				if updateErr != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		case "Please provide the new stop-loss for this open position or /cancel":
			positionSize := update.Message.Text
			positionSizeInt, err := strconv.Atoi(positionSize)
			if err != nil {
				fmt.Println("stopLoss Custom Error: ", err)
				responseText = "Invalid stop loss, please check your input! Your input was: " + update.Message.Text
			} else {
				err = rdbHoot.SetActivePerpEditSl(guid, float64(positionSizeInt))
				if err != nil {
					fmt.Println(err)
				}
				responseText = "Stop loss set!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				respMsgDel, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				time.Sleep(1 * time.Second)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
				updateErr := handlers.UpdateNewTradeOrActiveTrades(
					bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID, guid)
				if updateErr != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		case "Please provide take profit or /cancel":
			positionSize := update.Message.Text
			positionSizeInt, err := strconv.Atoi(positionSize)
			if err != nil {
				fmt.Println("stopLoss Custom Error: ", err)
				responseText = "Invalid take profit, please check your input! Your input was: " + update.Message.Text
			} else {
				err = rdbHoot.SetTakeProfit(guid, float64(positionSizeInt))
				if err != nil {
					fmt.Println(err)
				}
				responseText = "Take profit set!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				respMsgDel, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				time.Sleep(500 * time.Millisecond)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
				err = handlers.UpdateNewTradeOrActiveTrades(
					bot, client,
					rdbHoot,
					boardMsg.Chat.ID, boardMsg.MessageID,
					guid)
				if err != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		case "Please provide the new take-profit for this open position or /cancel":
			positionSize := update.Message.Text
			positionSizeInt, err := strconv.Atoi(positionSize)
			if err != nil {
				fmt.Println("stopLoss Custom Error: ", err)
				responseText = "Invalid take profit, please check your input! Your input was: " + update.Message.Text
			} else {
				err = rdbHoot.SetActivePerpEditTp(guid, float64(positionSizeInt))
				if err != nil {
					fmt.Println(err)
				}
				responseText = "Take profit set!"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
				respMsgDel, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
				time.Sleep(500 * time.Millisecond)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, update.Message.ReplyToMessage.MessageID)
				utils.DeleteMessage(bot, update.Message.Chat.ID, update.Message.MessageID)
				utils.DeleteMessage(bot, respMsgDel.Chat.ID, respMsgDel.MessageID)
				err = handlers.UpdateNewTradeOrActiveTrades(
					bot, client,
					rdbHoot,
					boardMsg.Chat.ID, boardMsg.MessageID,
					guid)
				if err != nil {
					log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
				}
				return
			}
		}
		// Check if responseText has been set, then create the message and send it
		if responseText != "" {
			err := handlers.UpdateNewTradeOrActiveTrades(
				bot, client, rdbHoot, boardMsg.Chat.ID, boardMsg.MessageID,
				guid)
			if err != nil {
				log.Println("Error in UpdateNewTradeOrActiveTrades: ", err)
			}
			return
		}
	}
}
