package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type UpdateOpenLimitOrderJSON struct {
	ID              int64  `json:"_id"`
	Index           string `json:"index"`
	Price           string `json:"price"`
	TP              string `json:"tp"`
	SL              string `json:"sl"`
	Chain           uint8  `json:"chain"`
	CollateralToken string `json:"collateral_token"`
	PairIndex       string `json:"pair_index"`
}

func UpdateOpenLimitOrder(guid int64, index uint32, newOpen float64, newTP float64, newSL float64, chain uint64, collateral uint8) error {
	// Create the payload
	payload := UpdateOpenOrderPayload{
		Guid:            guid,
		Index:           index,
		TriggerPrice:    newOpen,
		TP:              newTP,
		SL:              newSL,
		MaxSlippageP:    500, // TODO: Slippage should be able to be corrected
		Chain:           chain,
		CollateralToken: collateral,
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Get the URL from the environment variable
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	url += "updateopenorder"
	if url == "" {
		return fmt.Errorf("HOOT_CORE_API_LOCAL environment variable not set")
	}

	// Create the POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP status code
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, readErr := respBody.ReadFrom(resp.Body)
		if readErr != nil {
			return fmt.Errorf("received non-OK response status: %v and failed to read response body: %v", resp.StatusCode, readErr)
		}
		return fmt.Errorf("received non-OK response status: %v, body: %v", resp.StatusCode, respBody.String())
	}

	return nil
}
