package eventsgains

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func isMarketExecutedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "MarketExecuted", log)
	return err == nil, event
}

func isLimitExecutedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "LimitExecuted", log)
	return err == nil, event
}

func isMarketOpenCanceledEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "MarketOpenCanceled", log)
	return err == nil, event
}

func isMarketCloseCanceledEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "MarketCloseCanceled", log)
	return err == nil, event
}

func isNftOrderCanceledEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "NftOrderCanceled", log)
	return err == nil, event
}

func isClosingFeeSharesPUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "ClosingFeeSharesPUpdated", log)
	return err == nil, event
}

func isPairMaxLeverageUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "PairMaxLeverageUpdated", log)
	return err == nil, event
}
