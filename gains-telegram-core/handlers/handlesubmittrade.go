package handlers

import (
	"HootCore/api"
	"HootCore/api/tradinginteractions"
	"HootCore/rdbhoot"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func SubmitTrade(client *mongo.Client, rdbHoot *rdbhoot.HootRedisClient, guid int64) (string, tradinginteractions.OpenTradePayload, error) {
	// Fetch the trade from the global tradeCache
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		return "Cache Error", tradinginteractions.OpenTradePayload{}, errors.New("trade not found in cache")
	}

	//buyOrSell := ""
	//if trade.Buy {
	//	buyOrSell = "BUY"
	//} else {
	//	buyOrSell = "SELL"
	//}

	//inProgress := "Executing " + "*" + buyOrSell + "*" + " for " + "*" + pairmaps.IndexToPair[int(trade.PairIndex)] + "*" + " @ " + "*" + strconv.FormatFloat(trade.OpenPrice, 'f', 2, 64) + "*" + "\n\n Kindly wait for confirmation."
	//progressMsg := tgbotapi.NewMessage(chatID, inProgress)
	//progressMsg.ParseMode = tgbotapi.ModeMarkdown
	//returnProgressMsg, err := bot.Send(progressMsg)
	//if err != nil {
	//	log.Println("Error sending progress message", err)
	//}

	// We need to check here if market or not:
	tradeJson := trade.CreateOpenTradePayload()

	log.Println(tradeJson)

	// Convert the trade data to JSON
	jsonData, err := json.Marshal(tradeJson)
	if err != nil {
		return "Marshalling Error", tradeJson, fmt.Errorf("failed to marshal trade data: %w", err)
	}

	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return "Env Error", tradeJson, errors.New("environment variable HOOT_CORE_API_LOCAL is not set")
	}
	url += "opentrade"

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "POST not possible", tradinginteractions.OpenTradePayload{}, fmt.Errorf("failed to make POST request: %w", err)
	}
	defer resp.Body.Close()

	time.Sleep(1 * time.Second)

	//msgEdit := tgbotapi.NewEditMessageText(chatID, returnProgressMsg.MessageID, "Kindly wait...")
	//finalMsg, editErr := bot.Send(msgEdit)
	//if editErr != nil {
	//	log.Println("Error editing message", editErr)
	//}

	// Read the response body
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "respBody Error", tradinginteractions.OpenTradePayload{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Attempt to decode as a success response
	var responseOK api.HootCoreResponseOK
	if errDecode := json.Unmarshal(respBodyBytes, &responseOK); errDecode == nil {
		txReceipt := responseOK.Ok

		// I removed a time.sleep here

		// TODO: This needs to be fixed for the new redis cache

		//if trade.OrderType == 1 || trade.OrderType == 2 {
		//	cacheErr := tradecachemanagement_legacy.SetStartTradeOrOrder(client, tradeCache, guid, trade.Chain)
		//	if cacheErr != nil {
		//		log.Println("Error setting OrderType in submit.go")
		//	}
		//	log.Println("We set a StartTradeOrOrder in submit for => ", guid)
		//}
		return txReceipt, tradeJson, nil
	}

	// Attempt to decode as an error response
	var responseErr api.HootCoreResponseErr
	if errDecode := json.Unmarshal(respBodyBytes, &responseErr); errDecode == nil {
		return "", tradeJson, errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return "", tradeJson, fmt.Errorf("unhandled response body: %s", string(respBodyBytes))
}
