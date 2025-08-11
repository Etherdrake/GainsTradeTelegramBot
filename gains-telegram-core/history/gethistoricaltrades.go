package history

import (
	"HootCore/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetHistoricalTrades retrieves historical trades for a given address
func GetHistoricalTrades(address string) ([]TradeHistory, error) {
	address = strings.ToLower(address)
	checkSumAddress, _ := utils.ToChecksumAddress(address) // Convert address to lowercase
	url := fmt.Sprintf("https://backend-sepolia.gains.trade/personal-trading-history-table/%s", checkSumAddress)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers to mimic a web browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var trades []TradeHistory
	err = json.Unmarshal(body, &trades)
	if err != nil {
		log.Printf("Unmarshalling error: %v", err)
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	//log.Printf("Parsed Trades: %v", trades)

	return trades, nil
}
