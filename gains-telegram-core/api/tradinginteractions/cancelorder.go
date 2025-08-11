package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CancelOrderGNS(id int64, index uint32, chain uint64, collateral uint8) (string, error) {
	// Create the payload
	payload := CancelOpenOrderPayload{
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

	// Create the POST request
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	url += "cancelopenorder"
	// Replace with the actual URL
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

	// Read the response
	respBody := new(bytes.Buffer)
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Check for HTTP status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK response status: %v, body: %v", resp.StatusCode, respBody.String())
	}

	return respBody.String(), nil
}
