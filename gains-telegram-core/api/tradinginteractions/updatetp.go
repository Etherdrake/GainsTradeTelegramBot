package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func UpdateTpGns(guid int64, index uint32, newTP float64, chain uint64, collateral uint8) (string, error) {
	// Get the URL from the environment variable
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return "", fmt.Errorf("HOOT_CORE_API_LOCAL environment variable not set")
	}
	url += "/updatetpgns"

	// Create the payload
	updateTpJSON := UpdateTpPayload{
		Guid:            guid,
		Index:           index,
		NewTP:           newTP,
		Chain:           chain,
		CollateralToken: collateral,
	}

	// Convert the struct to JSON
	payload, err := json.Marshal(updateTpJSON)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("POST request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, readErr := respBody.ReadFrom(resp.Body)
		if readErr != nil {
			return "", fmt.Errorf("unexpected response status: %v and failed to read response body: %v", resp.Status, readErr)
		}
		return "", fmt.Errorf("unexpected response status: %v, body: %v", resp.Status, respBody.String())
	}

	// Read the response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(responseBody.Bytes(), &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(responseBody.Bytes(), &responseErr); errDecode == nil {
		return "", errors.New(responseErr.Err)
	}

	return "", fmt.Errorf("failed to decode response: %v", responseBody.String())
}
