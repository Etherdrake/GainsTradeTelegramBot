package altseasoncore

import (
	"time"
)

type GainsTradeContext struct {
	LastRefreshed             time.Time         `json:"lastRefreshed"`
	RefreshID                 int               `json:"refreshId"`
	TradingState              int               `json:"tradingState"`
	MaxGainP                  int               `json:"maxGainP"`
	MarketOrdersTimeoutBlocks int               `json:"marketOrdersTimeoutBlocks"`
	Pairs                     []Pair            `json:"pairs"`
	Groups                    []Group           `json:"groups"`
	Fees                      []Fee             `json:"fees"`
	PairInfos                 PairInfos         `json:"pairInfos"`
	Collaterals               []Collateral      `json:"collaterals"`
	SSSTokenBalance           string            `json:"sssTokenBalance"`
	SSSLegacyTokenBalance     string            `json:"sssLegacyTokenBalance"`
	SSSRewardTokens           []string          `json:"sssRewardTokens"`
	VaultClosingFeeP          string            `json:"vaultClosingFeeP"`
	MaxNegativePnlOnOpenP     int               `json:"maxNegativePnlOnOpenP"`
	BlockConfirmations        int               `json:"blockConfirmations"`
	OIWindowsSettings         OIWindowsSettings `json:"oiWindowsSettings"`
	OIWindows                 []OIWindow        `json:"oiWindows"`
	FeeTiers                  FeeTiers          `json:"feeTiers"`
	AllTrades                 []Trade           `json:"allTrades"`
	CurrentBlock              int               `json:"currentBlock"`
	CurrentL1Block            int               `json:"currentL1Block"`
	IsForexOpen               bool              `json:"isForexOpen"`
	IsStocksOpen              bool              `json:"isStocksOpen"`
	IsIndicesOpen             bool              `json:"isIndicesOpen"`
	IsCommoditiesOpen         bool              `json:"isCommoditiesOpen"`
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
	OpenFeeP           string `json:"openFeeP"`
	CloseFeeP          string `json:"closeFeeP"`
	OracleFeeP         string `json:"oracleFeeP"`
	TriggerOrderFeeP   string `json:"triggerOrderFeeP"`
	MinPositionSizeUsd string `json:"minPositionSizeUsd"`
}

type PairInfos struct {
	MaxLeverages []int       `json:"maxLeverages"`
	PairDepths   []PairDepth `json:"pairDepths"`
}

type PairDepth struct {
	OnePercentDepthAboveUsd string `json:"onePercentDepthAboveUsd"`
	OnePercentDepthBelowUsd string `json:"onePercentDepthBelowUsd"`
}

type Collateral struct {
	CollateralIndex  int              `json:"collateralIndex"`
	Collateral       string           `json:"collateral"`
	Symbol           string           `json:"symbol"`
	IsActive         bool             `json:"isActive"`
	Prices           CollateralPrices `json:"prices"`
	CollateralConfig CollateralConfig `json:"collateralConfig"`
	GToken           GToken           `json:"gToken"`
	BorrowingFees    BorrowingFees    `json:"borrowingFees"`
}

// Custom type that has been added
type CollateralPrices struct {
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
	GnsPriceCollateral float64 `json:"gnsPriceCollateral"`
	GnsPriceUsd        float64 `json:"gnsPriceUsd"`
}

type CollateralConfig struct {
	Precision      string `json:"precision"`
	PrecisionDelta string `json:"precisionDelta"`
	Decimals       int    `json:"decimals"`
}
type GToken struct {
	Address                  string `json:"address"`
	CurrentBalanceCollateral string `json:"currentBalanceCollateral"`
	MaxBalanceCollateral     string `json:"maxBalanceCollateral"`
	MarketCap                string `json:"marketCap"`
}

type BorrowingFees struct {
	Pairs  []Pairs  `json:"pairs"`
	Groups []Groups `json:"groups"`
}

type Pairs struct {
	Oi                  Oi       `json:"oi"`
	FeePerBlock         string   `json:"feePerBlock"`
	AccFeeLong          string   `json:"accFeeLong"`
	AccFeeShort         string   `json:"accFeeShort"`
	AccLastUpdatedBlock string   `json:"prevGroupAccFeeLong"`
	FeeExponent         string   `json:"feeExponent"`
	Groups              []Groups `json:"groups"`
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

type Oi struct {
	Long  string `json:"long"`
	Short string `json:"short"`
	Max   string `json:"max"`
}

type OIWindowsSettings struct {
	StartTs         int `json:"startTs"`
	WindowsDuration int `json:"windowsDuration"`
	WindowsCount    int `json:"windowsCount"`
}

type OIWindow struct {
	Window map[string]OIWindowDetails `json:"window"`
}

type OIWindowDetails struct {
	OILongUsd  string `json:"oiLongUsd"`
	OIShortUsd string `json:"oiShortUsd"`
}

type FeeTiers struct {
	Tiers       []Tier   `json:"tiers"`
	Multipliers []string `json:"multipliers"`
}

type Tier struct {
	FeeMultiplier   string `json:"feeMultiplier"`
	PointsThreshold string `json:"pointsThreshold"`
}

type Trade struct {
	Trade          TradeDetails   `json:"trade"`
	TradeInfo      TradeInfo      `json:"tradeInfo"`
	InitialAccFees InitialAccFees `json:"initialAccFees"`
}

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

type TradeInfo struct {
	CreatedBlock       string `json:"createdBlock"`
	TPLastUpdatedBlock string `json:"tpLastUpdatedBlock"`
	SLLastUpdatedBlock string `json:"slLastUpdatedBlock"`
	MaxSlippageP       string `json:"maxSlippageP"`
	LastOiUpdateTs     int    `json:"lastOiUpdateTs"`
	CollateralPriceUsd string `json:"collateralPriceUsd"`
}

type InitialAccFees struct {
	AccPairFee  string `json:"accPairFee"`
	AccGroupFee string `json:"accGroupFee"`
	Block       string `json:"block"`
}
