package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func UpdateSlGns(guid int64, index uint32, newSL float64, chain uint64, collateral uint8) (string, error) {
	// Create the payload
	payload := UpdateSlPayload{
		Guid:            guid,
		Index:           index,
		NewSL:           newSL,
		Chain:           chain,
		CollateralToken: collateral,
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Get the URL from the environment variable
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return "", fmt.Errorf("HOOT_CORE_API_LOCAL environment variable not set")
	}
	url += "updatesl"

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
	defer resp.Body.Close()

	// Check for HTTP status code
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, readErr := respBody.ReadFrom(resp.Body)
		if readErr != nil {
			return "", fmt.Errorf("received non-OK response status: %v and failed to read response body: %v", resp.StatusCode, readErr)
		}
		return "", fmt.Errorf("received non-OK response status: %v, body: %v", resp.StatusCode, respBody.String())
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
