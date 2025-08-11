package history

import "time"

const (
	TradeOpenedMarket TradeAction = "TradeOpenedMarket"
	TradeClosedMarket TradeAction = "TradeClosedMarket"
	TradeOpenedLimit  TradeAction = "TradeOpenedLimit"
	TradeClosedLimit  TradeAction = "TradeClosedLimit"
	TradeClosedTP     TradeAction = "TradeClosedTP"
	TradeClosedSL     TradeAction = "TradeClosedSL"
	TradeClosedLIQ    TradeAction = "TradeClosedLIQ"
)

// TradeHistory defines the structure of a trade record
type TradeHistory struct {
	Date               time.Time `json:"date"`
	Pair               string    `json:"pair"`
	Block              int64     `json:"block"`
	Address            string    `json:"address"`
	Action             string    `json:"action"`
	Price              float64   `json:"price"`
	CollateralPriceUsd float64   `json:"collateralPriceUsd"`
	Long               int       `json:"long"`
	Size               float64   `json:"size"`
	Leverage           float64   `json:"leverage"`
	Pnl                float64   `json:"pnl"`
	PnlNet             float64   `json:"pnl_net"`
	Tx                 string    `json:"tx"`
	TradeId            int       `json:"tradeId"`
	CollateralIndex    int       `json:"collateralIndex"`
}

// TradeAction represents the type of trade action
type TradeAction string

type TradeHistoryDisplay struct {
	Date               time.Time   `json:"date"`
	Pair               string      `json:"pair"`
	Block              string      `json:"block"`
	Address            string      `json:"address"`
	Action             TradeAction `json:"action"`
	Price              string      `json:"price"`
	CollateralPriceUsd string      `json:"collateralPriceUsd"`
	Long               string      `json:"long"`
	Size               string      `json:"size"`
	Leverage           string      `json:"leverage"`
	Pnl                string      `json:"pnl"`
	PnlNet             string      `json:"pnl_net"`
	Tx                 string      `json:"tx"`
	TradeId            string      `json:"tradeId"`
	CollateralIndex    string      `json:"collateralIndex"`
}
