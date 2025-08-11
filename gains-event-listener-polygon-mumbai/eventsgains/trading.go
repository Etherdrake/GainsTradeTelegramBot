package eventsgains

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func isDoneEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "Done", log)
	return err == nil, event
}

func isPausedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "Paused", log)
	return err == nil, event
}

func isNumberUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "NumberUpdated", log)
	return err == nil, event
}

func isBypassTriggerLinkUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "BypassTriggerLinkUpdated", log)
	return err == nil, event
}

func isMarketOrderInitiatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "MarketOrderInitiated", log)
	return err == nil, event
}

func isOpenLimitPlacedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "OpenLimitPlaced", log)
	return err == nil, event
}

func isOpenLimitUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "OpenLimitUpdated", log)
	return err == nil, event
}

func isTpUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "TpUpdated", log)
	return err == nil, event
}

func isSlUpdatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "SlUpdated", log)
	return err == nil, event
}

func isNftOrderInitiatedEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "NftOrderInitiated", log)
	return err == nil, event
}

func isChainlinkCallbackTimeoutEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "ChainlinkCallbackTimeoutEvent", log)
	return err == nil, event
}

func isCouldNotCloseTradeEvent(contractAbi abi.ABI, log types.Log) (bool, []interface{}) {
	event, err := unpackEvent(contractAbi, "CouldNotCloseTradeEvent", log)
	return err == nil, event
}
