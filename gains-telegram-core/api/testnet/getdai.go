package testnet

import (
	"HootCore/api"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GetFreeDaiPayload struct {
	Guid  int64  `json:"guid"`
	Chain uint64 `json:"chain"`
}

func GetFreeDai(guid int64, chain uint64) (string, error) {
	// Create the payload
	payload := GetFreeDaiPayload{
		Guid:  guid,
		Chain: chain,
	}

	log.Println(payload)

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Get the URL from the environment variable
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	url += "getfreedai"
	if url == "" {
		return "", fmt.Errorf("HOOT_CORE_API_LOCAL environment variable not set")
	}

	// Create the POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error getting free dai status:", resp.StatusCode, string(responseBody))
		// Attempt to decode as an error response
		var responseErr api.HootCoreResponseErr
		if errDecode := json.Unmarshal(responseBody, &responseErr); errDecode == nil {
			return "", errors.New(responseErr.Err)
		}
		return "", errors.New(responseErr.Err)
	}

	// Attempt to decode as a success response
	var responseOK api.HootCoreResponseOK
	if errDecode := json.Unmarshal(responseBody, &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr api.HootCoreResponseErr
	if errDecode := json.Unmarshal(responseBody, &responseErr); errDecode == nil {
		return "", errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return "", fmt.Errorf("unhandled response body: %s", string(responseBody))
}
