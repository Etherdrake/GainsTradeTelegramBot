package chartservice

import (
	"encoding/base64"
	"fmt"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func SendChart(bot *tgbotapi.BotAPI, chatID int64, base64String string) error {
	log.Println("Base64 string: ", base64String)
	// Decode base64 string
	baseExtract, err := ExtractBase64(base64String)
	if err != nil {
		return fmt.Errorf("error decoding extractBase64 base64: %v", err)
	}
	imageBytes, err := base64.StdEncoding.DecodeString(baseExtract)
	if err != nil {
		return fmt.Errorf("error decoding DecodeString base64: %v", err)
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

	// Send the photo
	_, err = bot.Send(photoConfig)
	if err != nil {
		return fmt.Errorf("error sending photo: %v", err)
	}

	return nil
}

func ExtractBase64(dataURL string) (string, error) {
	//log.Println("dataURL string: ", dataURL)
	// Split the string using the comma as a delimiter
	parts := strings.Split(dataURL, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid data URL format")
	}

	// Return the second part, which should be the base64 encoded string
	return parts[1], nil
}
