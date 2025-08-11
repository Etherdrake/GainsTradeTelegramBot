package native

import (
	"HootCore/database"
	"HootCore/gns_tools"
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
)

// srcToken is the token you want to approve TO the dstContract. Ie. DAI to GNSTrading
func ApproveCollateralToDiamond(client *mongo.Client, guid int64, trade rdbhoot.HootCoreCache) (string, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		log.Panic("environment variable HOOT_CORE_API_LOCAL is not set")
	}

	finalUrl := url + "approve"

	pubKey, _ := database.GetPublicKey(client, guid)

	token := gns_tools.GetCollateralTokenContractAddress(trade.Collateral, trade.Chain)

	spender := gns_tools.GetDiamondContractAddress(trade.Chain)

	payload := ApprovePayload{
		Guid:       trade.ID,
		Owner:      pubKey,
		Token:      token,
		Spender:    spender,
		Amount:     "100000000000000000000000000000000000000000000000000000000",
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	// Serialize the payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return "", err
	}

	// Send the payload to the Rust API
	resp, err := http.Post(finalUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("HTTP Post Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", err
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(responseBody, &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(responseBody, &responseErr); errDecode == nil {
		return "", errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return "", fmt.Errorf("unhandled response body: %s", string(responseBody))
}
