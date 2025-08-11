package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PairFees struct {
	CloseFee uint64 `json:"close_fee"`
	Spread   uint64 `json:"spread"`
}

func GetPairCloseFee(indexChain PayloadIndexChainGains) (uint64, error) {
	return getPairFee("getpairclosefee", indexChain)
}

func GetPairSpread(indexChain PayloadIndexChainGains) (uint64, error) {
	return getPairFee("getpairspread", indexChain)
}

func getPairFee(route string, indexChain PayloadIndexChainGains) (uint64, error) {
	// Serialize the payload
	payloadBytes, err := json.Marshal(indexChain)
	if err != nil {
		return 0, err
	}

	// Send the payload to the Rust API
	resp, err := http.Post(fmt.Sprintf("http://localhost:3030/%s", route), "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var pairFees string
	if err := json.NewDecoder(resp.Body).Decode(&pairFees); err != nil {
		return 0, err
	}

	var fee uint64
	switch route {
	case "getpairclosefee":
		fee, _ = strconv.ParseUint(pairFees, 10, 64)
	case "getpairspread":
		fee, _ = strconv.ParseUint(pairFees, 10, 64)
	}

	return fee, nil
}

func GetMaxLeverageGains(route string, id int64, index int64, chain string) (int64, error) {
	Packet := GetMaxLeverage{
		ID:    id,
		Index: index,
		Chain: chain,
	}
	// Serialize the payload
	payloadBytes, err := json.Marshal(Packet)
	if err != nil {
		return 0, err
	}

	// Send the payload to the Rust API
	resp, err := http.Post(fmt.Sprintf("http://localhost:3030/%s", route), "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var maxLeverage string
	if err := json.NewDecoder(resp.Body).Decode(&maxLeverage); err != nil {
		return 0, err
	}
	maxLvrgInt, err := strconv.Atoi(maxLeverage)
	if err != nil {
		return -1, err
	}
	return int64(maxLvrgInt), nil
}
