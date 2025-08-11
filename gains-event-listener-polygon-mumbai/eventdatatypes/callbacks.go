package eventdatatypes

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// MarketExecutedEvent represents the MarketExecuted event
type MarketExecuted struct {
	OrderID            string
	T                  Trade
	Open               bool
	Price              *big.Int
	PriceImpactP       *big.Int
	PositionSizeDai    *big.Int
	PercentProfit      *big.Int
	DaiSentToTrader    *big.Int
	CollateralPriceUsd *big.Int
}

// LimitExecutedEvent represents the LimitExecuted event
type LimitExecuted struct {
	OrderID    string
	LimitIndex *big.Int
	Trade      Trade
	//NftHolder       common.Address this is an index event [2]
	OrderType          uint8
	Price              *big.Int
	PriceImpactP       *big.Int
	PositionSizeDai    *big.Int
	PercentProfit      *big.Int
	DaiSentToTrader    *big.Int
	CollateralPriceUsd *big.Int
	ExactExecution     bool
}

// MarketOpenCanceledEvent represents the MarketOpenCanceled event
type MarketOpenCanceled struct {
	OrderID      uint64
	Trader       string
	PairIndex    uint64
	CancelReason uint8
}

// MarketCloseCanceledEvent represents the MarketCloseCanceled event
type MarketCloseCanceled struct {
	OrderID      uint64
	Trader       string
	PairIndex    uint64
	Index        *big.Int
	CancelReason uint8
}

// NftOrderCanceledEvent represents the NftOrderCanceled event, it is the same as MarketOpenCanceled
// but just an old type.
type NftOrderCanceled struct {
	OrderID      uint64
	NftHolder    string
	OrderType    uint8
	CancelReason uint8
}

// Trade represents the Trade struct
type Trade struct {
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
}

type PairMaxLeverageUpdated struct {
	PairIndex   *big.Int
	MaxLeverage *big.Int
}

// LimitOrder represents the type of orders
type LimitOrder int

// Enumerating limit order types
const (
	TP LimitOrder = iota
	SL
	LIQ
	OPEN
)

// CancelReason represents the reason for order cancellation
type CancelReason int

// Enumerating cancel reasons
const (
	Invalid CancelReason = iota
	UserRequested
	InsufficientFunds
)
