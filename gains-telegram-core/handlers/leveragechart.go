package handlers

import (
	"HootCore/boardbuilders/newtradebuilders"
	"HootCore/chartservice"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/Etherdrake/telegram-bot-api/v7"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func LeverageChartSender(rdbHoot *rdbhoot.HootRedisClient, timeFrame string, bot *tgbotapi.BotAPI, guid int64, chatId int64) {
	// Update the message content with the desired message
	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist.")
	}

	var base64String string
	var board tgbotapi.InlineKeyboardMarkup
	var chartErr error
	switch timeFrame {
	case "1M":
		board = newtradebuilders.BuildNewTradeBoardCharts(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart1M(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	case "5M":
		board = newtradebuilders.BuildNewTradeBoardCharts5M(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart5M(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	case "15M":
		board = newtradebuilders.BuildNewTradeBoardCharts15M(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart15M(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	case "1H":
		board = newtradebuilders.BuildNewTradeBoardCharts1H(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart1H(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	case "4H":
		board = newtradebuilders.BuildNewTradeBoardCharts4H(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart4H(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	case "1D":
		board = newtradebuilders.BuildNewTradeBoardCharts1D(rdbHoot, guid)
		base64String, chartErr = chartservice.GetChart1D(pairmaps.IndexToPair[int(trade.PairIndex)])
		if chartErr != nil {
			log.Printf("/sendchart GetChart%s error: %s", timeFrame, chartErr)
		}
	}

	// Decode base64 string
	baseExtract, err := chartservice.ExtractBase64(base64String)
	if err != nil {
		log.Println(err)
	}
	imageBytes, err := base64.StdEncoding.DecodeString(baseExtract)
	if err != nil {
		log.Println(err)
	}

	// Generate a unique file name based on the current timestamp
	fileName := fmt.Sprintf("photo_%d.jpg", time.Now().UnixNano())

	// Create a temporary file to store the decoded image
	tmpfile, err := ioutil.TempFile("", fileName)
	if err != nil {
		log.Println()
	}
	defer func() {
		// Close the file before removing it
		tmpfile.Close()
		// Remove the temporary file
		os.Remove(tmpfile.Name())
	}()

	// Write the image data to the temporary file
	_, err = tmpfile.Write(imageBytes)
	if err != nil {
		log.Println()
	}

	// Seek to the beginning of the file for reading
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		log.Println()
	}

	// Create a FileBytes from the temporary file
	fileBytes := tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: imageBytes,
	}

	// Create a PhotoConfig to send the photo
	photoConfig := tgbotapi.NewPhoto(chatId, fileBytes)
	photoConfig.ReplyMarkup = &board
	// Assign the new inline keyboard to the message
	// Send the photo with the inline keyboard
	_, err = bot.Send(photoConfig)
	return
}
