package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func CloseTradeGNS(id int64, index uint32, chain uint64, collateral uint8) (string, error) {
	// Create the payload
	payload := CloseTradeMarketPayload{
		Guid:            id,
		Index:           index,
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
	url += "closetrademarket"
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
	respBody := new(bytes.Buffer)
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(respBody.Bytes(), &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(respBody.Bytes(), &responseErr); errDecode == nil {
		return "", errors.New(responseErr.Err)
	}

	return respBody.String(), nil
}
