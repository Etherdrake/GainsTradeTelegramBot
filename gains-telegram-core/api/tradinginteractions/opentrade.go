package tradinginteractions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func OpenTrade(guid int64, trade TradePayload, chain uint64, collateralToken uint8) (string, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL") + "/opentrade"

	openTradeJSON := OpenTradePayload{
		Guid:            guid,
		Trade:           trade,
		Chain:           chain,
		CollateralToken: collateralToken,
	}

	// Convert the struct to JSON
	payload, err := json.Marshal(openTradeJSON)
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
			return "", fmt.Errorf("received non-OK response status: %v and failed to read response body: %v", resp.StatusCode, readErr)
		}
		return "", fmt.Errorf("received non-OK response status: %v, body: %v", resp.StatusCode, respBody.String())
	}

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	log.Println("OpenTrade request successful!")
	return responseBody.String(), nil
}
