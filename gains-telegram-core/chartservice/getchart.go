package chartservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetChart1M(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 60,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func GetChart5M(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 300,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func GetChart15M(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 900,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func GetChart1H(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 3600,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func GetChart4H(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 14400,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func GetChart1D(pair string) (string, error) {
	chartUrl := os.Getenv("CHARTS_URL")

	// Create payload with default values
	payload := Payload{
		Pair:                pair,
		CandlesWidthSeconds: 86400,
		NumCandles:          100,
	}

	// Perform API request with payload
	apiResponse, err := PerformApiRequest(chartUrl+"", payload)
	if err != nil {
		return "", err
	}

	// Return only the chart string
	return apiResponse.Chart, nil
}

func PerformApiRequest(uri string, payload Payload) (ApiResponse, error) {
	log.Println(uri)

	// Convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling payload to JSON in PerformApiRequest")
		return ApiResponse{}, err
	}

	// Make HTTP POST request
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Println("Error making HTTP POST request in PerformApiRequest")
		return ApiResponse{}, err
	}
	defer resp.Body.Close()

	// Check if the status code is 404
	if resp.StatusCode == http.StatusNotFound {
		log.Println("Received 404 status code in PerformApiRequest")
		return ApiResponse{}, fmt.Errorf("received 404 status code from %s", uri)
	}

	// Decode the JSON response
	var apiResponse ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		log.Println("Error decoding JSON response in PerformApiRequest")
		return ApiResponse{}, err
	}
	return apiResponse, nil
}
