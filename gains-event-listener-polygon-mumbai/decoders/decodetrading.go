package decoders

import (
	"GainsListenerMumbai/eventdatatypes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DecodeDoneEvent decodes Done event
func DecodeDoneEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.DoneEvent, error) {
	unpacked, err := contractAbi.Unpack("Done", eventLog.Data)
	if err != nil {
		return nil, err
	}

	event := &eventdatatypes.DoneEvent{
		Done: unpacked[0].(bool),
	}
	return event, nil
}

// DecodePausedEvent decodes Paused event
func DecodePausedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.PausedEvent, error) {
	unpacked, err := contractAbi.Unpack("Paused", eventLog.Data)
	if err != nil {
		return nil, err
	}

	event := &eventdatatypes.PausedEvent{
		Paused: unpacked[0].(bool),
	}
	return event, nil
}

// DecodeNumberUpdatedEvent decodes NumberUpdated event
func DecodeNumberUpdatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.NumberUpdatedEvent, error) {
	unpacked, err := contractAbi.Unpack("NumberUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	event := &eventdatatypes.NumberUpdatedEvent{
		Name:  unpacked[0].(string),
		Value: unpacked[1].(*big.Int),
	}
	return event, nil
}

// DecodeBypassTriggerLinkUpdatedEvent decodes BypassTriggerLinkUpdated event
func DecodeBypassTriggerLinkUpdatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.BypassTriggerLinkUpdatedEvent, error) {
	unpacked, err := contractAbi.Unpack("BypassTriggerLinkUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	event := &eventdatatypes.BypassTriggerLinkUpdatedEvent{
		User:   unpacked[0].(common.Address).String(),
		Bypass: unpacked[1].(bool),
	}
	return event, nil
}

// DecodeMarketOrderInitiatedEvent decodes MarketOrderInitiated event
func DecodeMarketOrderInitiatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.MarketOrderInitiatedEvent, error) {
	unpacked, err := contractAbi.Unpack("MarketOrderInitiated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	OrderID := eventLog.Topics[1].Big()
	Trader := common.BytesToAddress(eventLog.Topics[2].Bytes()).String()
	pairIndex := eventLog.Topics[3].Big()

	event := &eventdatatypes.MarketOrderInitiatedEvent{
		OrderID:   OrderID.String(),
		Trader:    Trader,
		PairIndex: pairIndex,
		Open:      unpacked[0].(bool),
	}
	return event, nil
}

// DecodeOpenLimitPlacedEvent decodes OpenLimitPlaced event
func DecodeOpenLimitPlacedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.OpenLimitPlacedEvent, error) {
	unpacked, err := contractAbi.Unpack("OpenLimitPlaced", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.OpenLimitPlacedEvent{
		Trader:    Trader,
		PairIndex: pairIndex,
		Index:     unpacked[0].(*big.Int),
	}
	return event, nil
}

// DecodeOpenLimitUpdatedEvent decodes OpenLimitUpdated event
func DecodeOpenLimitUpdatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.OpenLimitUpdatedEvent, error) {
	unpacked, err := contractAbi.Unpack("OpenLimitUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.OpenLimitUpdatedEvent{
		Trader:       Trader,
		PairIndex:    pairIndex,
		Index:        unpacked[0].(*big.Int),
		NewPrice:     unpacked[1].(*big.Int),
		NewTp:        unpacked[2].(*big.Int),
		NewSl:        unpacked[3].(*big.Int),
		MaxSlippageP: unpacked[4].(*big.Int),
	}
	return event, nil
}

// DecodeOpenLimitCanceledEvent decodes OpenLimitCanceled event
func DecodeOpenLimitCanceledEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.OpenLimitCanceledEvent, error) {
	unpacked, err := contractAbi.Unpack("OpenLimitCanceled", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.OpenLimitCanceledEvent{
		Trader:    Trader,
		PairIndex: pairIndex,
		Index:     unpacked[0].(*big.Int),
	}
	return event, nil
}

// DecodeTpUpdatedEvent decodes TpUpdated event
func DecodeTpUpdatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.TpUpdatedEvent, error) {
	unpacked, err := contractAbi.Unpack("TpUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.TpUpdatedEvent{
		Trader:    Trader,
		PairIndex: pairIndex,
		Index:     unpacked[0].(*big.Int),
		NewTp:     unpacked[1].(*big.Int),
	}
	return event, nil
}

// DecodeSlUpdatedEvent decodes SlUpdated event
func DecodeSlUpdatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.SlUpdatedEvent, error) {
	unpacked, err := contractAbi.Unpack("SlUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.SlUpdatedEvent{
		Trader:    Trader,
		PairIndex: pairIndex,
		Index:     unpacked[0].(*big.Int),
		NewSl:     unpacked[1].(*big.Int),
	}
	return event, nil
}

// DecodeNftOrderInitiatedEvent decodes NftOrderInitiated event
func DecodeNftOrderInitiatedEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.NftOrderInitiatedEvent, error) {
	unpacked, err := contractAbi.Unpack("NftOrderInitiated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.NftOrderInitiatedEvent{
		OrderID:          unpacked[0].(*big.Int).String(),
		Trader:           Trader,
		PairIndex:        pairIndex,
		ByPassesLinkCost: unpacked[1].(bool),
	}
	return event, nil
}

// DecodeChainlinkCallbackTimeoutEvent decodes ChainlinkCallbackTimeout event
func DecodeChainlinkCallbackTimeoutEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.ChainlinkCallbackTimeoutEvent, error) {
	unpacked, err := contractAbi.Unpack("ChainlinkCallbackTimeout", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].Big()

	event := &eventdatatypes.ChainlinkCallbackTimeoutEvent{
		OrderID: orderID,
		Order:   unpacked[0].(eventdatatypes.PendingMarketOrder),
	}
	return event, nil
}

// DecodeCouldNotCloseTradeEvent decodes CouldNotCloseTrade event
func DecodeCouldNotCloseTradeEvent(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.CouldNotCloseTradeEvent, error) {
	unpacked, err := contractAbi.Unpack("CouldNotCloseTrade", eventLog.Data)
	if err != nil {
		return nil, err
	}

	Trader := common.BytesToAddress(eventLog.Topics[1].Bytes()).String()
	pairIndex := eventLog.Topics[2].Big()

	event := &eventdatatypes.CouldNotCloseTradeEvent{
		Trader:    Trader,
		PairIndex: pairIndex,
		Index:     unpacked[0].(*big.Int),
	}
	return event, nil
}
