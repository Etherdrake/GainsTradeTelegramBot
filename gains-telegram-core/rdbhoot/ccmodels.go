package rdbhoot

import "HootCore/altseasoncore"

type HootCoreCache struct {
	ID              int64           `json:"_id"`
	Trader          string          `json:"trader"`
	PairIndex       int64           `json:"pair_index"`
	Index           int64           `json:"activeOrderIndex"`
	InitialPosToken int64           `json:"initial_pos_token"`
	PositionSizeDai float64         `json:"position_size_dai"` // TODO: Needs to be edited so it is different than what we use
	OpenPrice       float64         `json:"open_price"`
	ClosePrice      float64         `json:"close_price"`
	Buy             bool            `json:"buy"`
	Leverage        float32         `json:"leverage"`
	TP              float64         `json:"tp"`
	SL              float64         `json:"sl"`
	Liq             float64         `json:"liq"`
	PercentageTP    float64         `json:"percentage_tp"`
	PercentageSL    float64         `json:"percentage_sl"`
	Collateral      uint8           `json:"collateral"`
	OrderType       uint8           `json:"order_type"`
	OrderStatus     uint8           `json:"order_status"`
	PairPage        uint8           `json:"pair_page"`
	PnL             float64         `json:"pnl"`
	Chain           uint64          `json:"chain"`
	ActiveCharts    bool            `json:"active_charts"`
	ActiveWindow    uint8           `json:"active_window"` // 0 is newtrade, 1 is active position window.
	ActiveOpenTrade int             `json:"active_open_trade"`
	ActivePerp      ActivePerpetual `json:"active_perp"`
}

type ActivePerpetual struct {
	IsTradeOrOrder     uint // 0 is nothing set, 1 trade, 2 is order
	ActiveTradeID      string
	ActiveTradeOrOrder altseasoncore.TradeItemInternal
}

var ordertypes = map[int]string{
	0: "MARKET",
	1: "LIMIT",
	2: "STOP",
}

var orderStatus = map[int]string{
	0: "NONE",
	1: "IN_PROGRESS",
	2: "OPEN_ORDER",
	3: "OPEN_TRADE",
	4: "CLOSED",
}
