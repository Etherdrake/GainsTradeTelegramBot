package api

// Changes to OpenTradeJSON must also be made to the
// corresponding  Rust structs
type OpenTradeJSON struct {
	ID              int64   `json:"_id"`
	TradeID         string  `json:"trade_id"`
	Trader          string  `json:"trader"`
	Paper           bool    `json:"paper"`
	PairIndex       string  `json:"pair_index"`
	Index           string  `json:"index"`
	InitialPosToken string  `json:"initial_pos_token"`
	PositionSizeDai string  `json:"position_size_dai"`
	OpenPrice       string  `json:"open_price"`
	ClosePrice      string  `json:"close_price"`
	Buy             bool    `json:"buy"`
	Leverage        string  `json:"leverage"`
	TP              string  `json:"tp"`
	SL              string  `json:"sl"`
	Liq             string  `json:"liq"`
	PercentageTP    string  `json:"percentage_tp"`
	PercentageSL    string  `json:"percentage_sl"`
	OrderType       uint8   `json:"order_type"`
	OrderStatus     uint8   `json:"order_status"`
	PnL             float64 `json:"pnl"`
	Chain           uint64  `json:"chain"`
	// Added values - these are values we add that did
	// not exist on this struct when it was in the cache
	Timestamp       uint64 `json:"timestamp"`
	CollateralToken string `json:"collateral_token"`
}

type GetTokenBalancePayload struct {
	ID         int64  `json:"_id"`
	PublicKey  string `json:"public_key"`
	Collateral int    `json:"collateral"`
	Chain      uint64 `json:"chain"`
}

type GetWETH struct {
	ID        int64  `json:"_id"`
	PublicKey string `json:"public_key"`
	Chain     uint64 `json:"chain"`
}

type GetUSDC struct {
	ID        int64  `json:"_id"`
	PublicKey string `json:"public_key"`
	Chain     uint64 `json:"chain"`
}

type GetMaxLeverage struct {
	ID              int64  `json:"_id"`
	Index           int64  `json:"index"`
	Chain           uint64 `json:"chain"`
	CollateralToken string `json:"collateral_token"`
}

type UpdateStopLossJSON struct {
	ID              int64  `json:"_id"`
	PairIndex       string `json:"pair_index"`
	Index           string `json:"index"`
	NewSL           string `json:"new_sl"`
	Chain           uint64 `json:"chain"`
	CollateralToken string `json:"collateral_token"`
}

type UpdateTakeProfitJSON struct {
	ID              int64  `json:"_id"`
	PairIndex       string `json:"pair_index"`
	Index           string `json:"index"`
	NewTP           string `json:"new_tp"`
	Chain           uint64 `json:"chain"`
	CollateralToken string `json:"collateral_token"`
}
