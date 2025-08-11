package eventsgains

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func unpackEvent(contractAbi abi.ABI, eventName string, log types.Log) ([]interface{}, error) {
	event, err := contractAbi.Unpack(eventName, log.Data)
	return event, err
}
