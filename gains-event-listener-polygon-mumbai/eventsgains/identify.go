package eventsgains

import (
	"GainsListenerMumbai/topicmaps"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func IdentifyCallbackEvent(contractAbi abi.ABI, log types.Log) (string, []interface{}, bool) {
	// Define the list of event identification functions
	eventIdentificationFunctions := map[string]func(abi.ABI, types.Log) (bool, []interface{}){
		"MarketExecuted":           isMarketExecutedEvent,
		"LimitExecuted":            isLimitExecutedEvent,
		"MarketOpenCanceled":       isMarketOpenCanceledEvent,
		"MarketCloseCanceled":      isMarketCloseCanceledEvent,
		"NftOrderCanceled":         isNftOrderCanceledEvent,
		"ClosingFeeSharesPUpdated": isClosingFeeSharesPUpdatedEvent,
		"PairMaxLeverageUpdated":   isPairMaxLeverageUpdatedEvent,
	}

	// Iterate over the event identification functions
	for eventName, identificationFunc := range eventIdentificationFunctions {
		isEvent, eventData := identificationFunc(contractAbi, log)
		if isEvent {
			return eventName, eventData, true
		}
	}

	// If no event is identified
	return "", nil, false
}

func IdentifyTradingEvent(contractAbi abi.ABI, log types.Log) (string, bool) {
	// Define the list of event identification functions
	if eventName, ok := topicmaps.TopicHashToEventnameTrading[log.Topics[0].String()]; ok {
		return eventName, true
	}

	// If no event is identified
	return "", false
}

func IdentifyTradingEventUnpack(contractAbi abi.ABI, log types.Log) (string, []interface{}, bool) {
	// Define the list of event identification functions
	eventIdentificationFunctions := map[string]func(abi.ABI, types.Log) (bool, []interface{}){
		"Done":                          isDoneEvent,
		"Paused":                        isPausedEvent,
		"NumberUpdated":                 isNumberUpdatedEvent,
		"BypassTriggerLinkUpdated":      isBypassTriggerLinkUpdatedEvent,
		"MarketOrderInitiated":          isMarketOrderInitiatedEvent,
		"OpenLimitPlaced":               isOpenLimitPlacedEvent,
		"OpenLimitUpdated":              isOpenLimitUpdatedEvent,
		"TpUpdated":                     isTpUpdatedEvent,
		"SlUpdated":                     isSlUpdatedEvent,
		"NftOrderInitiated":             isNftOrderInitiatedEvent,
		"ChainlinkCallbackTimeoutEvent": isChainlinkCallbackTimeoutEvent,
		"CouldNotCloseTradeEvent":       isCouldNotCloseTradeEvent,
	}

	// Iterate over the event identification functions
	for eventName, identificationFunc := range eventIdentificationFunctions {
		isEvent, eventData := identificationFunc(contractAbi, log)
		if isEvent {
			return eventName, eventData, true
		}
	}

	// If no event is identified
	return "", nil, false
}
