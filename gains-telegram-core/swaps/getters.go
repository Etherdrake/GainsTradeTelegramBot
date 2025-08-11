package swaps

import (
	"HootCore/rdbhoot"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.zcx.com/trade/v1/info/token"
)

func GetSwapConfig(ctx context.Context, rdbHoot *rdbhoot.HootRedisClient, guid string) (*SwapInfos, error) {
	log.Println("Retrieving swap configuration...")

	// Get SwapInfos JSON from Redis using RedisJSON
	val, err := rdbHoot.JSONGet(ctx, guid, ".").Result()
	if err != nil {
		return nil, fmt.Errorf("error retrieving SwapInfos JSON from Redis: %v", err)
	}

	// Unmarshal JSON to SwapInfos struct
	var swapInfos SwapInfos
	err = json.Unmarshal([]byte(val), &swapInfos)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling SwapInfos JSON: %v", err)
	}

	return &swapInfos, nil
}

func GetTokenInfo(chainID int, contractAddress string) (*TokenInfo, error) {
	url := fmt.Sprintf("%s/%d/%s", baseURL, chainID, contractAddress)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error with NewRequest: ", err)
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("UNIZEN_AUTH_KEY_TRIAL"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to Unizen: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var tokenInfo UnizenTokenInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		log.Println("Error decoding JSON response: ", err)
		return nil, err
	}
	return &tokenInfo.TokenInfo, nil
}
