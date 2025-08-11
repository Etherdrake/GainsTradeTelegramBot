package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PayloadIndexChainGains struct {
	ID    int64  `json:"_id"`
	Index uint64 `json:"index"`
	Chain string `json:"chain"`
}

func GetOpenFeeForPair(indexChain PayloadIndexChainGains) (uint64, error) {
	// Serialize the payload
	payloadBytes, err := json.Marshal(indexChain)
	if err != nil {
		return 0, err
	}

	// Send the payload to the Rust API
	resp, err := http.Post("http://localhost:3030/getopenfeeforpair", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Fee uint64 `json:"fee"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Fee, nil
}
