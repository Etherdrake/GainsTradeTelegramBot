package altseasoncore

type OpenTradesMongo struct {
	Name  string           `json:"name"`
	Value []OpenTradeMongo `json:"value"`
}

type OpenTradeMongo struct {
	Trade          TradeMongo          `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}

type TradeMongo struct {
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

type TradeInfoMongo struct {
	CreatedBlock       string `json:"createdBlock"`
	TpLastUpdatedBlock string `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock string `json:"slLastUpdatedBlock"`
	MaxSlippageP       string `json:"maxSlippageP"`
	LastOiUpdateTs     string `json:"lastOiUpdateTs"`
	CollateralPriceUsd string `json:"collateralPriceUsd"`
}

type InitialAccFeesMongo struct {
	AccPairFee  string `json:"accPairFee"`
	AccGroupFee string `json:"accGroupFee"`
	Block       string `json:"block"`
}

type TradeItemMongo struct {
	ID             string              `bson:"_id,omitempty"`
	Trade          TradeMongo          `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}
