package chartservice

import (
	"HootCore/boardbuilders/newtradebuilders"
	"HootCore/rdbhoot"
	"encoding/base64"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"io/ioutil"
	"os"
	"time"
)

func SendChartBoardTimeframe(bot *tgbotapi.BotAPI, chatID int64, base64String string, rdbHoot *rdbhoot.HootRedisClient, guid int64, timeframe int) error {
	// Decode base64 string
	baseExtract, err := ExtractBase64(base64String)
	if err != nil {
		return fmt.Errorf("error decoding base64: %v", err)
	}
	imageBytes, err := base64.StdEncoding.DecodeString(baseExtract)
	if err != nil {
		return fmt.Errorf("error decoding base64: %v", err)
	}

	// Generate a unique file name based on the current timestamp
	fileName := fmt.Sprintf("photo_%d.jpg", time.Now().UnixNano())

	// Create a temporary file to store the decoded image
	tmpfile, err := ioutil.TempFile("", fileName)
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
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
		return fmt.Errorf("error writing to temporary file: %v", err)
	}

	// Seek to the beginning of the file for reading
	_, err = tmpfile.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("error seeking to the beginning of the file: %v", err)
	}

	// Create a FileBytes from the temporary file
	fileBytes := tgbotapi.FileBytes{
		Name:  fileName,
		Bytes: imageBytes,
	}

	// Create a PhotoConfig to send the photo
	photoConfig := tgbotapi.NewPhoto(chatID, fileBytes)

	var chartKeyboard tgbotapi.InlineKeyboardMarkup
	// Create an inline keyboard
	if timeframe == 1 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart1M(rdbHoot, guid)
	}
	if timeframe == 5 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart5M(rdbHoot, guid)
	}
	if timeframe == 15 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart15M(rdbHoot, guid)
	}
	if timeframe == 60 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart1H(rdbHoot, guid)
	}
	if timeframe == 240 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart4H(rdbHoot, guid)
	}
	if timeframe == 960 {
		chartKeyboard = newtradebuilders.BuildNewTradeBoardChart1D(rdbHoot, guid)
	}

	// Set the inline keyboard in the message
	photoConfig.ReplyMarkup = &chartKeyboard

	// Send the photo with the inline keyboard
	_, err = bot.Send(photoConfig)
	if err != nil {
		return fmt.Errorf("error sending photo: %v", err)
	}

	return nil
}
