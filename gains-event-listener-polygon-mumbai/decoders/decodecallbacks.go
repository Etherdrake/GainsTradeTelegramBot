package decoders

import (
	"GainsListenerMumbai/eventdatatypes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DecodeMarketExecuted(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.MarketExecuted, error) {
	unpacked, err := contractAbi.Unpack("MarketExecuted", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].String()

	// Define the struct that matches the fields in the data
	trade := unpacked[0].(struct {
		Trader          common.Address `json:"trader"`
		PairIndex       *big.Int       `json:"pairIndex"`
		Index           *big.Int       `json:"index"`
		InitialPosToken *big.Int       `json:"initialPosToken"`
		PositionSizeDai *big.Int       `json:"positionSizeDai"`
		OpenPrice       *big.Int       `json:"openPrice"`
		Buy             bool           `json:"buy"`
		Leverage        *big.Int       `json:"leverage"`
		Tp              *big.Int       `json:"tp"`
		Sl              *big.Int       `json:"sl"`
	})

	event := &eventdatatypes.MarketExecuted{
		OrderID:            orderID,
		T:                  trade,
		Open:               unpacked[1].(bool),
		Price:              unpacked[2].(*big.Int),
		PriceImpactP:       unpacked[3].(*big.Int),
		PositionSizeDai:    unpacked[4].(*big.Int),
		PercentProfit:      unpacked[5].(*big.Int),
		DaiSentToTrader:    unpacked[6].(*big.Int),
		CollateralPriceUsd: unpacked[7].(*big.Int),
	}
	return event, nil
}

func DecodeLimitExecuted(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.LimitExecuted, error) {
	unpacked, err := contractAbi.Unpack("LimitExecuted", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].String()

	// Define the struct that matches the fields in the data
	trade := unpacked[1].(struct {
		Trader          common.Address `json:"trader"`
		PairIndex       *big.Int       `json:"pairIndex"`
		Index           *big.Int       `json:"index"`
		InitialPosToken *big.Int       `json:"initialPosToken"`
		PositionSizeDai *big.Int       `json:"positionSizeDai"`
		OpenPrice       *big.Int       `json:"openPrice"`
		Buy             bool           `json:"buy"`
		Leverage        *big.Int       `json:"leverage"`
		Tp              *big.Int       `json:"tp"`
		Sl              *big.Int       `json:"sl"`
	})

	event := &eventdatatypes.LimitExecuted{
		OrderID:            orderID,
		LimitIndex:         unpacked[0].(*big.Int),
		Trade:              trade,
		OrderType:          unpacked[2].(uint8),
		Price:              unpacked[3].(*big.Int),
		PriceImpactP:       unpacked[4].(*big.Int),
		PositionSizeDai:    unpacked[5].(*big.Int),
		PercentProfit:      unpacked[6].(*big.Int),
		DaiSentToTrader:    unpacked[7].(*big.Int),
		CollateralPriceUsd: unpacked[8].(*big.Int),
		ExactExecution:     unpacked[9].(bool),
	}
	return event, nil
}

func DecodeMarketOpenCanceled(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.MarketOpenCanceled, error) {
	unpacked, err := contractAbi.Unpack("MarketOpenCanceled", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].Big().Uint64()
	trader := eventLog.Topics[2].Hex()
	pairIndex := eventLog.Topics[3].Big().Uint64()

	event := &eventdatatypes.MarketOpenCanceled{
		OrderID:      orderID,
		Trader:       trader,
		PairIndex:    pairIndex,
		CancelReason: unpacked[0].(uint8),
	}

	return event, nil
}

func DecodeMarketCloseCanceled(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.MarketCloseCanceled, error) {
	unpacked, err := contractAbi.Unpack("MarketCloseCanceled", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].Big().Uint64()
	trader := eventLog.Topics[2].Hex()
	pairIndex := eventLog.Topics[3].Big().Uint64()

	event := &eventdatatypes.MarketCloseCanceled{
		OrderID:      orderID,
		Trader:       trader,
		PairIndex:    pairIndex,
		Index:        unpacked[0].(*big.Int),
		CancelReason: unpacked[1].(uint8),
	}

	return event, nil
}

func DecodePairMaxLeverageUpdated(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.PairMaxLeverageUpdated, error) {
	unpacked, err := contractAbi.Unpack("PairMaxLeverageUpdated", eventLog.Data)
	if err != nil {
		return nil, err
	}

	pairIndex := eventLog.Topics[1].Big()

	event := &eventdatatypes.PairMaxLeverageUpdated{
		PairIndex:   pairIndex,
		MaxLeverage: unpacked[0].(*big.Int),
	}

	return event, nil
}

// DecodeNftOrderCanceled
//
//		event NftOrderCanceled(
//	    uint indexed orderId,
//	    address indexed nftHolder,
//	    StorageInterfaceV5.LimitOrder orderType,
//	    CancelReason cancelReason
func DecodeNftOrderCanceled(contractAbi *abi.ABI, eventLog *types.Log) (*eventdatatypes.NftOrderCanceled, error) {
	unpacked, err := contractAbi.Unpack("NftOrderCanceled", eventLog.Data)
	if err != nil {
		return nil, err
	}

	orderID := eventLog.Topics[1].Big().Uint64()
	nftHolder := eventLog.Topics[2].Hex()

	event := &eventdatatypes.NftOrderCanceled{
		OrderID:      orderID,
		NftHolder:    nftHolder,
		OrderType:    unpacked[0].(uint8),
		CancelReason: unpacked[0].(uint8),
	}
	return event, nil
}
