package types

import "strings"

type CurrentBlock struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type CurrentL1Block struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Event struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type TradingVariables struct {
	LastRefreshed             string  `json:"lastRefreshed"`
	RefreshID                 int     `json:"refreshId"`
	TradingState              int     `json:"tradingState"`
	MaxGainP                  int     `json:"maxGainP"`
	MarketOrdersTimeoutBlocks int     `json:"marketOrdersTimeoutBlocks"`
	Pairs                     []Pair  `json:"pairs"`
	Groups                    []Group `json:"groups"`
	Fees                      []Fee   `json:"fees"`
}

type Pair struct {
	From       string `json:"from"`
	To         string `json:"to"`
	SpreadP    string `json:"spreadP"`
	GroupIndex string `json:"groupIndex"`
	FeeIndex   string `json:"feeIndex"`
}

type Group struct {
	Name        string `json:"name"`
	MinLeverage string `json:"minLeverage"`
	MaxLeverage string `json:"maxLeverage"`
}

type Fee struct {
	OpenFeeP string `json:"openFeeP"`
}

type UnregisterTrade struct {
	Name  string            `json:"name"`
	Value UnregisteredTrade `json:"value"`
}

type UnregisteredTrade struct {
	User  string `json:"user"`
	Index string `json:"index"`
}

type RegisterTrade struct {
	Name  string    `json:"name"`
	Value OpenTrade `json:"value"`
}

type OpenTrades struct {
	Name  string `json:"name"`
	Value []OpenTrade
}

type OpenTrade struct {
	Trade          Trade               `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}

func (t *Trade) Lowercase() {
	t.User = strings.ToLower(t.User)
	t.Index = strings.ToLower(t.Index)
	t.PairIndex = strings.ToLower(t.PairIndex)
	t.Leverage = strings.ToLower(t.Leverage)
	t.CollateralIndex = strings.ToLower(t.CollateralIndex)
	t.TradeType = strings.ToLower(t.TradeType)
	t.CollateralAmount = strings.ToLower(t.CollateralAmount)
	t.OpenPrice = strings.ToLower(t.OpenPrice)
	t.Tp = strings.ToLower(t.Tp)
	t.Sl = strings.ToLower(t.Sl)
}

type Trade struct {
	User             string `json:"user"`
	Index            string `json:"index"`
	PairIndex        string `json:"pairIndex"`
	Leverage         string `json:"leverage"`
	Long             bool   `json:"long"`
	IsOpen           bool   `json:"isOpen"`
	CollateralIndex  string `json:"collateralIndex"`
	TradeType        string `json:"tradeType"`
	CollateralAmount string `json:"collateralAmount"`
	OpenPrice        string `json:"openPrice"`
	Tp               string `json:"tp"`
	Sl               string `json:"sl"`
}

type TradeInfo struct {
	CreatedBlock       string `json:"createdBlock"`
	TpLastUpdatedBlock string `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock string `json:"slLastUpdatedBlock"`
	MaxSlippageP       string `json:"maxSlippageP"`
	LastOiUpdateTs     string `json:"lastOiUpdateTs"`
	CollateralPriceUsd string `json:"collateralPriceUsd"`
}

type InitialAccFees struct {
	AccPairFee  string `json:"accPairFee"`
	AccGroupFee string `json:"accGroupFee"`
	Block       string `json:"block"`
}

type TradeHistory struct {
	Name               string  `json:"name"`
	Date               string  `json:"date"`
	Pair               string  `json:"pair"`
	Address            string  `json:"address"`
	Action             string  `json:"action"`
	Price              float64 `json:"price"`
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
	Long               int     `json:"long"`
	Size               float64 `json:"size"`
	Leverage           float64 `json:"leverage"`
	PnlNet             float64 `json:"pnl_net"`
	Tx                 string  `json:"tx"`
	CollateralIndex    string  `json:"collateralIndex"`
}

type CollateralBalances struct {
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Balances []string `json:"balances"`
}

type LiveEvent struct {
	Name            string       `json:"name"`
	Event           string       `json:"event"`
	BlockNumber     int          `json:"blockNumber"`
	L1BlockNumber   int          `json:"l1BlockNumber"`
	TransactionHash string       `json:"transactionHash"`
	ReturnValues    ReturnValues `json:"returnValues"`
}

type ReturnValues struct {
	Index     []interface{} `json:"0"`
	User      string        `json:"1"`
	PairIndex string        `json:"2"`
	IsOpen    bool          `json:"3"`
	OrderID   []interface{} `json:"orderId"`
	Trader    string        `json:"trader"`
	Open      bool          `json:"open"`
}
