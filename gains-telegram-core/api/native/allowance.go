package native

import (
	"HootCore/database"
	gns_tools "HootCore/gns_tools"
	"HootCore/rdbhoot"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Message struct {
	Ok  string `json:"Ok"`
	Err string `json:"Err"`
}

func GetAllowanceGains(client *mongo.Client, trade rdbhoot.HootCoreCache) (string, float64, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL")

	publicKey, _ := database.GetPublicKey(client, trade.ID)

	// Define the URL
	params := "getallowance"
	finalUrl := url + params

	diamondAddr := gns_tools.GetDiamondContractAddress(trade.Chain)
	tokenAddr := gns_tools.GetCollateralTokenContractAddress(trade.Collateral, trade.Chain)

	payload := GetAllowancePayload{
		Owner:      publicKey,
		Token:      tokenAddr,
		Spender:    diamondAddr,
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", 0, fmt.Errorf("error marshaling payload: %v", err)
	}

	// Create a new HTTP request with the payload
	req, err := http.NewRequest("POST", finalUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", 0, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and send the request
	clientHyper := &http.Client{}
	resp, err := clientHyper.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Read and parse the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("error reading response body: %v", err)
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(responseBody, &responseOK); errDecode == nil {
		respConv, err := strconv.ParseFloat(responseOK.Ok, 64)
		if err != nil {
			log.Printf("error converting OK to uint64: %v", err)
		}
		return responseOK.Ok, respConv, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(responseBody, &responseErr); errDecode == nil {
		return "", 0, errors.New(responseErr.Err)
	}
	return "", 0, fmt.Errorf("error decoding response body: %v", err)
}
