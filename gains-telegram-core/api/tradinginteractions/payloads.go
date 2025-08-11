package tradinginteractions

// TradePayload represents the payload received from HootCore
type TradePayload struct {
	User             string  `json:"user"`
	Index            uint32  `json:"index"`
	PairIndex        uint16  `json:"pair_index"`
	Leverage         float32 `json:"leverage"`
	Long             bool    `json:"long"`
	IsOpen           bool    `json:"is_open"`
	CollateralIndex  uint8   `json:"collateral_index"`
	TradeType        uint8   `json:"trade_type"`
	CollateralAmount float64 `json:"collateral_amount"`
	OpenPrice        float64 `json:"open_price"`
	TP               float64 `json:"tp"`
	SL               float64 `json:"sl"`
	Placeholder      uint    `json:"__placeholder"`
}

// OpenTradePayload represents the payload for opening a trade
type OpenTradePayload struct {
	Guid            int64        `json:"guid"`
	Trade           TradePayload `json:"trade"`
	Chain           uint64       `json:"chain"`
	CollateralToken uint8        `json:"collateral_token"`
}

// CancelOpenOrderPayload represents the payload for cancelling an open order
type CancelOpenOrderPayload struct {
	Guid            int64  `json:"guid"`
	Index           uint32 `json:"index"`
	Chain           uint64 `json:"chain"`
	CollateralToken uint8  `json:"collateral_token"`
}

// CloseTradeMarketPayload represents the payload for closing a trade at market
type CloseTradeMarketPayload struct {
	Guid            int64  `json:"guid"`
	Index           uint32 `json:"index"`
	Chain           uint64 `json:"chain"`
	CollateralToken uint8  `json:"collateral_token"`
}

// UpdateOpenOrderPayload represents the payload for updating an open order
type UpdateOpenOrderPayload struct {
	Guid            int64   `json:"guid"`
	Index           uint32  `json:"index"`
	TriggerPrice    float64 `json:"trigger_price"`
	TP              float64 `json:"tp"`
	SL              float64 `json:"sl"`
	MaxSlippageP    uint16  `json:"max_slippage_p"`
	Chain           uint64  `json:"chain"`
	CollateralToken uint8   `json:"collateral_token"`
}

// UpdateSlPayload represents the payload for updating the stop loss
type UpdateSlPayload struct {
	Guid            int64   `json:"guid"`
	Index           uint32  `json:"index"`
	NewSL           float64 `json:"new_sl"`
	Chain           uint64  `json:"chain"`
	CollateralToken uint8   `json:"collateral_token"`
}

// UpdateTpPayload represents the payload for updating the take profit
type UpdateTpPayload struct {
	Guid            int64   `json:"guid"`
	Index           uint32  `json:"index"`
	NewTP           float64 `json:"new_tp"`
	Chain           uint64  `json:"chain"`
	CollateralToken uint8   `json:"collateral_token"`
}
