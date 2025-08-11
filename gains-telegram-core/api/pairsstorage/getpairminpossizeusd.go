package pairsstorage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MinPosSizeUsdResponse struct {
	MinPosSizeUsd int64  `json:"min_pos_size_usd"`
	PairIndex     int64  `json:"pair_index"`
	Chain         string `json:"chain"`
}

func GetPairMinPosSizeUsd(pairIndex uint64, chain uint64) (float64, error) {
	url := os.Getenv("HOOT_CORE_API_LOCAL")
	if url == "" {
		return 0, errors.New("environment variable HOOT_CORE_API_LOCAL is not set")
	}
	url += "getpairminpossizeusd"

	requestData := GetPairMinPosSizeUsdPayload{
		Index: pairIndex,
		Chain: chain,
	}

	// Convert the requestData to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return 0, err
	}

	// Make the POST request
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	// Attempt to decode as a success response
	var responseOK HootCoreResponseFloatOK
	if errDecode := json.Unmarshal(body, &responseOK); errDecode == nil {
		return responseOK.Ok, nil
	}

	// Attempt to decode as an error response
	var responseErr HootCoreResponseErr
	if errDecode := json.Unmarshal(body, &responseErr); errDecode == nil {
		return 0, errors.New(responseErr.Err)
	}

	// If decoding fails for both, return an error
	return 0, fmt.Errorf("unhandled response body: %s", string(body))
}
