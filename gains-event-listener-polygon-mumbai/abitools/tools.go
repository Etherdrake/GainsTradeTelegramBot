package abitools

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"strings"
)

const (
	GainsCallbackAddressPoly = "0x82e59334da8c667797009bbe82473b55c7a6b311"
	GainsCallbackAddressArb  = "0x42069"
	GainsTradingAbiPath      = "./gains_abi/GNSTradingV6_4_1_abi.json"
	GainsCallbackAbiPath     = "./gains_abi/GNSTradingCallbacksV6_4_1_abi.json"
)

func GetEventHash(abi, eventName string) (string, error) {
	// Concatenate event signature
	eventSignature := fmt.Sprintf("%s(%s)", eventName, abi)

	// Hash the event signature using go-ethereum's Keccak256 function
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	// Convert hash to a hex string
	hashString := hash.Hex()

	return hashString, nil
}

func GetEventId(contractName, eventName string) (string, error) {
	var path string
	if contractName == "trading" {
		path = GainsTradingAbiPath
	} else {
		path = GainsCallbackAbiPath
	}

	// Read the contents of the ABI file
	contractABI, err := ReadABIFile(path)
	if err != nil {
		log.Println(err)
	}

	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Println(err)
	}

	// Find the event in the parsed ABI
	//var event abi.Event
	for _, e := range parsedABI.Events {
		if e.Name == eventName {
			//event = e
			return e.ID.String(), nil
		}
	}
	return "", err
}

func GetEventHashGains(contractName, eventName string) (string, error) {
	var path string
	if contractName == "trading" {
		path = GainsTradingAbiPath
	} else {
		path = GainsCallbackAbiPath
	}

	// Read the contents of the ABI file
	contractABI, err := ReadABIFile(path)
	if err != nil {
		log.Println(err)
	}

	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Println(err)
	}

	// Find the event in the parsed ABI
	var event abi.Event
	for _, e := range parsedABI.Events {
		if e.Name == eventName {
			event = e
			break
		}
	}

	// Check if the event was found
	if event.Name == "" {
		return "", fmt.Errorf("Event not found: %s", eventName)
	}

	// Hash the event signature using go-ethereum's Keccak256 function
	hash := crypto.Keccak256Hash([]byte(event.ID.String()))

	// Convert hash to a hex string
	hashString := hash.Hex()

	return hashString, nil
}

func ReadABIFile(filePath string) (string, error) {
	// Read the contents of the file
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
