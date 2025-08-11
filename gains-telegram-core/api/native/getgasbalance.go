package native

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// GetGasBalance returns the gas balance in ether as a string, float, and error
func GetGasBalance(payload GetGasBalancePayload) (string, float64, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return "", 0, errors.New("environment variable HOOT_CORE_API_LOCAL is not set")
	}
	url += "getgasbalance"

	// Serialize the payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return "", 0, err
	}

	// Send the payload to the Rust API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("HTTP Post Error:", err)
		return "", 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return "", 0, err
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(responseBody, &responseOK); errDecode == nil {
		etherBalanceStr := responseOK.Ok
		etherBalance, err := strconv.ParseFloat(etherBalanceStr, 64)
		if err != nil {
			return "", 0, err
		}
		return etherBalanceStr, etherBalance, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(responseBody, &responseErr); errDecode == nil {
		return "", 0, errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return "", 0, fmt.Errorf("unhandled response body: %s", string(responseBody))
}
