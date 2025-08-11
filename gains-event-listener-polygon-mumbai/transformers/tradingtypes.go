package transformers

// DoneEventTransform
type DoneEventTransform struct {
	Done            bool   `json:"done" bson:"done"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// PausedEventTransform represents the Paused event
type PausedEventTransform struct {
	Paused          bool   `json:"paused" bson:"paused"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// NumberUpdatedEventTransform represents the NumberUpdated event
type NumberUpdatedEventTransform struct {
	Name            string `json:"name" bson:"name"`
	Value           int64  `json:"value" bson:"value"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// BypassTriggerLinkUpdatedEventTransform represents the BypassTriggerLinkUpdated event
type BypassTriggerLinkUpdatedEventTransform struct {
	User            string `json:"user" bson:"user"`
	Bypass          bool   `json:"bypass" bson:"bypass"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// MarketOrderInitiatedEventTransform
type MarketOrderInitiatedEventTransform struct {
	OrderID         string `json:"order_id"  bson:"_id"`
	Trader          string `json:"trader" bson:"trader"`
	PairIndex       int64  `json:"pair_index" bson:"pair_index"`
	Open            bool   `json:"open" bson:"open"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// OpenLimitPlacedEventTransform trader has created an open limit. It's final, no waiting no oracle checks.
// This is about an ORDER that is being placed and has gone live so we do NOT look for this event on the listener.
type OpenLimitPlacedEventTransform struct {
	Trader          string `json:"trader"  bson:"_id"`
	PairIndex       int64  `json:"pair_index" bson:"pair_index"`
	Index           int64  `json:"index" bson:"index"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// OpenLimitUpdatedEventTransform is an updated limit order that is not open yet. A limit orders stops being a limit order as soon
// as it is executed. We do NOT track this event because it is an order and NOT a trade.
type OpenLimitUpdatedEventTransform struct {
	Trader          string  `json:"trader"  bson:"_id"`
	PairIndex       int64   `json:"pair_index" bson:"pair_index"`
	Index           int64   `json:"index" bson:"index"`
	NewPrice        float64 `json:"new_price" bson:"new_price"`
	NewTp           float64 `json:"new_tp" bson:"new_tp"`
	NewSl           float64 `json:"new_sl" bson:"new_sl"`
	MaxSlippageP    float64 `json:"max_slippage_p" bson:"max_slippage_p"`
	CollateralToken string  `json:"collateral_token" bson:"collateral_token"`
}

// OpenLimitCanceledEventTransform user cancels the limit order that is open. Cannot fail.
// We do not track this because it about an ORDER and not a TRADE.
type OpenLimitCanceledEventTransform struct {
	Trader          string `json:"trader"  bson:"_id"`
	PairIndex       int64  `json:"pair_index" bson:"pair_index"`
	Index           int64  `json:"index" bson:"index"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// TpUpdatedEventTransform is for normal trades / market trades AND open trades that were opened using a limit order.
// We do track this event because it is about trades
type TpUpdatedEventTransform struct {
	Trader          string  `json:"trader"  bson:"_id"`
	PairIndex       int64   `json:"pair_index" bson:"pair_index"`
	Index           int64   `json:"index" bson:"index"`
	NewTp           float64 `json:"new_tp" bson:"new_tp"`
	CollateralToken string  `json:"collateral_token" bson:"collateral_token"`
}

// SlUpdatedEventTransform represents the SlUpdated event
// We do track this event because it is about trades
type SlUpdatedEventTransform struct {
	Trader          string  `json:"trader"  bson:"_id"`
	PairIndex       int64   `json:"pair_index" bson:"pair_index"`
	Index           int64   `json:"index" bson:"index"`
	NewSl           float64 `json:"new_sl" bson:"new_sl"`
	CollateralToken string  `json:"collateral_token" bson:"collateral_token"`
}

// NftOrderInitiatedEventTransform a limit order, could be an actual limit order / tp / sl. Nft just means that the triggering
// was done by a NFT bot. We do not track this because we track the actual execution of the trade
type NftOrderInitiatedEventTransform struct {
	OrderID          int64  `json:"order_id"  bson:"_id"`
	Trader           string `json:"trader" bson:"trader"`
	PairIndex        int64  `json:"pair_index" bson:"pair_index"`
	ByPassesLinkCost bool   `json:"bypasses_link_cost" bson:"bypasses_link_cost"`
	CollateralToken  string `json:"collateral_token" bson:"collateral_token"`
}

// ChainlinkCallbackTimeoutEventTransform represents the ChainlinkCallbackTimeout event
type ChainlinkCallbackTimeoutEventTransform struct {
	OrderID         int64                        `json:"order_id"  bson:"_id"`
	Order           *PendingMarketOrderTransform `json:"order" bson:"order"`
	CollateralToken string                       `json:"collateral_token" bson:"collateral_token"`
}

// CouldNotCloseTradeEventTransform it's when you have a pending market order, but the oracles don't respond and the
// trade is not able to be closed. The oracle times-out. User has to click MarketClose again. It's when a trader presses
// close trade market, and it times out. Check the trading contract. When you call the time-out it tries it again.
// User requests close -> emits request -> oracle does not respond -> times out -> user can try and close it again.
type CouldNotCloseTradeEventTransform struct {
	Trader          string `json:"trader"  bson:"_id"`
	PairIndex       int64  `json:"pair_index" bson:"pair_index"`
	Index           int64  `json:"index" bson:"index"`
	CollateralToken string `json:"collateral_token" bson:"collateral_token"`
}

// PendingMarketOrderTransform represents a pending market order.
type PendingMarketOrderTransform struct {
	Trade            TradeTransform `json:"trade"  bson:"_id"`
	Block            int64          `json:"block" bson:"block"`
	WantedPrice      float64        `json:"wanted_price" bson:"wanted_price"` // PRECISION
	SlippageP        float64        `json:"slippage_p" bson:"slippage_p"`     // PRECISION (%)
	SpreadReductionP float64        `json:"spread_reduction_p" bson:"spread_reduction_p"`
	TokenID          int            `json:"token_id" bson:"token_id"` // index in supportedTokens
	CollateralToken  string         `json:"collateral_token" bson:"collateral_token"`
}
