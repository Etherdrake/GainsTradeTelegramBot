package types

type TradeItemDisplay struct {
	ID             string                `json:"id"`
	Trade          TradeDisplay          `json:"trade"`
	TradeInfo      TradeInfoDisplay      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesDisplay `json:"initialAccFees"`
}

type TradeDisplay struct {
	User             string `json:"user"`
	Index            string `json:"index"`
	PairIndex        string `json:"pairIndex"`
	Leverage         string `json:"leverage"`
	Long             string `json:"long"`
	IsOpen           string `json:"isOpen"`
	CollateralIndex  string `json:"collateralIndex"`
	TradeType        string `json:"tradeType"`
	CollateralAmount string `json:"collateralAmount"`
	OpenPrice        string `json:"openPrice"`
	Tp               string `json:"tp"`
	Sl               string `json:"sl"`
}

type TradeInfoDisplay struct {
	CreatedBlock       string `json:"createdBlock"`
	TpLastUpdatedBlock string `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock string `json:"slLastUpdatedBlock"`
	MaxSlippageP       string `json:"maxSlippageP"`
	LastOiUpdateTs     string `json:"lastOiUpdateTs"`
	CollateralPriceUsd string `json:"collateralPriceUsd"`
}

type InitialAccFeesDisplay struct {
	AccPairFee  string `json:"accPairFee"`
	AccGroupFee string `json:"accGroupFee"`
	Block       string `json:"block"`
}
