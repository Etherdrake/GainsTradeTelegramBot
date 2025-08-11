package types

type TradeDetails struct {
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
	TP               string `json:"tp"`
	SL               string `json:"sl"`
}

type Tier struct {
	FeeMultiplier   string `json:"feeMultiplier"`
	PointsThreshold string `json:"pointsThreshold"`
}

type Groups struct {
	GroupIndex           string `json:"groupIndex"`
	Block                string `json:"block"`
	InitialAccFeeLong    string `json:"accFeeLong"`
	InitialAccFeeShort   string `json:"accFeeShort"`
	PrevGroupAccFeeLong  string `json:"prevGroupAccFeeLong"`
	PrevGroupAccFeeShort string `json:"prevGroupAccFeeShort"`
	PairAccFeeLong       string `json:"pairAccFeeLong"`
	PairAccFeeShort      string `json:"pairAccFeeShort"`
}

type FeeTiers struct {
	Tiers       []Tier   `json:"tiers"`
	Multipliers []string `json:"multipliers"`
}
