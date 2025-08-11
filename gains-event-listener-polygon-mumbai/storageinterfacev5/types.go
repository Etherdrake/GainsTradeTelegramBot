package storageinterfacev5

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// LimitOrder represents the type of orders
type LimitOrder int

// Enumerating limit order types
const (
	TP LimitOrder = iota
	SL
	LIQ
	OPEN
)

// Trade represents a trade
type Trade struct {
	Trader          string
	PairIndex       uint256.Int
	Index           uint256.Int
	InitialPosToken uint256.Int // 1e18
	PositionSizeDai uint256.Int // 1e18
	OpenPrice       uint256.Int // PRECISION
	Buy             bool
	Leverage        uint256.Int
	Tp              uint256.Int // PRECISION
	Sl              uint256.Int // PRECISION
}

// TradeInfo represents information about a trade
type TradeInfo struct {
	TokenID           uint256.Int
	TokenPriceDai     uint256.Int // PRECISION
	OpenInterestDai   uint256.Int // 1e18
	TpLastUpdated     uint256.Int
	SlLastUpdated     uint256.Int
	BeingMarketClosed bool
}

// OpenLimitOrder represents an open limit order
type OpenLimitOrder struct {
	Trader           string
	PairIndex        uint256.Int
	Index            uint256.Int
	PositionSize     uint256.Int // 1e18 (DAI or GFARM2)
	SpreadReductionP uint256.Int
	Buy              uint256.Int
	Leverage         uint256.Int
	Tp               uint256.Int // PRECISION (%)
	Sl               uint256.Int // PRECISION (%)
	MinPrice         uint256.Int // PRECISION
	MaxPrice         uint256.Int // PRECISION
	Block            uint256.Int
	TokenID          uint256.Int // index in supportedTokens
}

// PendingMarketOrder represents a pending market order
type PendingMarketOrder struct {
	Trade
	Block            uint256.Int
	WantedPrice      uint256.Int // PRECISION
	SlippageP        uint256.Int // PRECISION (%)
	SpreadReductionP uint256.Int
	TokenID          uint256.Int // index in supportedTokens
}

// PendingNftOrder represents a pending NFT order
type PendingNftOrder struct {
	NftHolder common.Address
	NftID     uint256.Int
	Trader    common.Address
	PairIndex uint256.Int
	Index     uint256.Int
	OrderType LimitOrder
}
