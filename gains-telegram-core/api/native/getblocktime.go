package native

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetBlockTime(blocknumber, chain uint64) (string, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return "", errors.New("environment variable HOOT_CORE_API_LOCAL is not set")
	}
	url += "getblocktime"

	requestData := GetBlockTimePayload{
		BlockNumber: blocknumber,
		Chain:       chain,
	}

	// Convert the requestData to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// Make the POST request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "GET error: CHECK IF THE API IS RUNNING", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseOK
	if errDecode := json.Unmarshal(body, &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(body, &responseErr); errDecode == nil {
		return "", errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return "", fmt.Errorf("unhandled response body: %s", string(body))
}
