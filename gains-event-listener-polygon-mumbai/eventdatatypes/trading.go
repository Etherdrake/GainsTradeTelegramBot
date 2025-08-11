package eventdatatypes

import (
	"math/big"
)

// DoneEvent represents the Done event
type DoneEvent struct {
	Done bool
}

// PausedEvent represents the Paused event
type PausedEvent struct {
	Paused bool
}

// NumberUpdatedEvent represents the NumberUpdated event
type NumberUpdatedEvent struct {
	Name  string
	Value *big.Int
}

// BypassTriggerLinkUpdatedEvent represents the BypassTriggerLinkUpdated event
type BypassTriggerLinkUpdatedEvent struct {
	User   string
	Bypass bool
}

// MarketOrderInitiatedEvent represents the MarketOrderInitiated event
type MarketOrderInitiatedEvent struct {
	OrderID   string
	Trader    string
	PairIndex *big.Int
	Open      bool
}

// OpenLimitPlacedEvent represents the OpenLimitPlaced event
type OpenLimitPlacedEvent struct {
	Trader    string
	PairIndex *big.Int
	Index     *big.Int
}

// OpenLimitUpdatedEvent represents the OpenLimitUpdated event
type OpenLimitUpdatedEvent struct {
	Trader       string
	PairIndex    *big.Int
	Index        *big.Int
	NewPrice     *big.Int
	NewTp        *big.Int
	NewSl        *big.Int
	MaxSlippageP *big.Int
}

// OpenLimitCanceledEvent represents the OpenLimitCanceled event
type OpenLimitCanceledEvent struct {
	Trader    string
	PairIndex *big.Int
	Index     *big.Int
}

// TpUpdatedEvent represents the TpUpdated event
type TpUpdatedEvent struct {
	Trader    string
	PairIndex *big.Int
	Index     *big.Int
	NewTp     *big.Int
}

// SlUpdatedEvent represents the SlUpdated event
type SlUpdatedEvent struct {
	Trader    string
	PairIndex *big.Int
	Index     *big.Int
	NewSl     *big.Int
}

// NftOrderInitiatedEvent represents the NftOrderInitiated event
type NftOrderInitiatedEvent struct {
	OrderID          string
	Trader           string
	PairIndex        *big.Int
	ByPassesLinkCost bool
}

// ChainlinkCallbackTimeoutEvent represents the ChainlinkCallbackTimeout event
type ChainlinkCallbackTimeoutEvent struct {
	OrderID *big.Int
	Order   PendingMarketOrder
}

// CouldNotCloseTradeEvent represents the CouldNotCloseTrade event
type CouldNotCloseTradeEvent struct {
	Trader    string
	PairIndex *big.Int
	Index     *big.Int
}

// PendingMarketOrder represents a pending market order
type PendingMarketOrder struct {
	Trade
	Block            *big.Int
	WantedPrice      *big.Int // PRECISION
	SlippageP        *big.Int // PRECISION (%)
	SpreadReductionP *big.Int
	TokenID          *big.Int // index in supportedTokens
}
