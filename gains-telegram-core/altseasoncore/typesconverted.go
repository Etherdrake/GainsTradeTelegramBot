package altseasoncore

import (
	"time"
)

type GainsTradeContextConverted struct {
	LastRefreshed             time.Time                  `json:"lastRefreshed"`
	RefreshID                 int                        `json:"refreshId"`
	TradingState              int                        `json:"tradingState"`
	MaxGainP                  int                        `json:"maxGainP"`
	MarketOrdersTimeoutBlocks uint64                     `json:"marketOrdersTimeoutBlocks"`
	Pairs                     []PairConverted            `json:"pairs"`
	Groups                    []GroupConverted           `json:"groups"`
	Fees                      []FeeConverted             `json:"fees"`
	PairInfos                 PairInfosConverted         `json:"pairInfos"`
	Collaterals               []CollateralConverted      `json:"collaterals"`
	SSSTokenBalance           string                     `json:"sssTokenBalance"`
	SSSLegacyTokenBalance     string                     `json:"sssLegacyTokenBalance"`
	SSSRewardTokens           []string                   `json:"sssRewardTokens"`
	VaultClosingFeeP          string                     `json:"vaultClosingFeeP"`
	MaxNegativePnlOnOpenP     int                        `json:"maxNegativePnlOnOpenP"`
	BlockConfirmations        uint64                     `json:"blockConfirmations"`
	OIWindowsSettings         OIWindowsSettingsConverted `json:"oiWindowsSettings"`
	OIWindows                 []OIWindowConverted        `json:"oiWindows"`
	FeeTiers                  FeeTiers                   `json:"feeTiers"`
	AllTrades                 []TradeConverted           `json:"allTrades"`
	CurrentBlock              uint64                     `json:"currentBlock"`
	CurrentL1Block            uint64                     `json:"currentL1Block"`
	IsForexOpen               bool                       `json:"isForexOpen"`
	IsStocksOpen              bool                       `json:"isStocksOpen"`
	IsIndicesOpen             bool                       `json:"isIndicesOpen"`
	IsCommoditiesOpen         bool                       `json:"isCommoditiesOpen"`
}

type PairConverted struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	SpreadP    float64 `json:"spreadP"`
	GroupIndex int     `json:"groupIndex"`
	FeeIndex   int     `json:"feeIndex"`
}

type GroupConverted struct {
	Name        string  `json:"name"`
	MinLeverage float64 `json:"minLeverage"`
	MaxLeverage float64 `json:"maxLeverage"`
}

type FeeConverted struct {
	OpenFeeP           float64 `json:"openFeeP"`           // 1e10 (% of position size)
	CloseFeeP          float64 `json:"closeFeeP"`          // 1e10 (% of position size)
	OracleFeeP         float64 `json:"oracleFeeP"`         // 1e10 (% of position size)
	TriggerOrderFeeP   float64 `json:"triggerOrderFeeP"`   // 1e10 (% of position size)
	MinPositionSizeUsd float64 `json:"minPositionSizeUsd"` // 1e18 (collateral x leverage, useful for min fee)
}

type PairInfosConverted struct {
	MaxLeverages []int                `json:"maxLeverages"`
	PairDepths   []PairDepthConverted `json:"pairDepths"`
}

type PairDepthConverted struct {
	OnePercentDepthAboveUsd uint64 `json:"onePercentDepthAboveUsd"`
	OnePercentDepthBelowUsd uint64 `json:"onePercentDepthBelowUsd"`
}

type CollateralConverted struct {
	CollateralIndex  int                       `json:"collateralIndex"`
	Collateral       string                    `json:"collateral"`
	Symbol           string                    `json:"symbol"`
	IsActive         bool                      `json:"isActive"`
	Prices           CollateralPricesConverted `json:"prices"`
	CollateralConfig CollateralConfigConverted `json:"collateralConfig"`
	GToken           GTokenConverted           `json:"gToken"`
	BorrowingFees    BorrowingFeesConverted    `json:"borrowingFees"`
}

// Custom type that has been added
type CollateralPricesConverted struct {
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
	GnsPriceCollateral float64 `json:"gnsPriceCollateral"`
	GnsPriceUsd        float64 `json:"gnsPriceUsd"`
}

type CollateralConfigConverted struct {
	Precision      string `json:"precision"`
	PrecisionDelta string `json:"precisionDelta"`
	Decimals       int    `json:"decimals"`
}
type GTokenConverted struct {
	Address                  string `json:"address"`
	CurrentBalanceCollateral string `json:"currentBalanceCollateral"`
	MaxBalanceCollateral     string `json:"maxBalanceCollateral"`
	MarketCap                string `json:"marketCap"`
}

type BorrowingFeesConverted struct {
	Pairs  []PairsConverted  `json:"pairs"`
	Groups []GroupsConverted `json:"groups"`
}

type PairsConverted struct {
	Oi                  OiConverted `json:"oi"`
	FeePerBlock         float64     `json:"feePerBlock"` // 1e10 (%)
	AccFeeLong          float64     `json:"accFeeLong"`  // 1e10 (%)
	AccFeeShort         float64     `json:"accFeeShort"` // 1e10 (%)
	AccLastUpdatedBlock uint64      `json:"prevGroupAccFeeLong"`
	FeeExponent         float64     `json:"feeExponent"`
	Groups              []Groups    `json:"groups"`
}

type GroupsConverted struct {
	GroupIndex           float64 `json:"groupIndex"`
	Block                uint64  `json:"block"`
	InitialAccFeeLong    float64 `json:"accFeeLong"`           // 1e10 (%)
	InitialAccFeeShort   float64 `json:"accFeeShort"`          // 1e10 (%)
	PrevGroupAccFeeLong  float64 `json:"prevGroupAccFeeLong"`  // 1e10 (%)
	PrevGroupAccFeeShort float64 `json:"prevGroupAccFeeShort"` // 1e10 (%)
	PairAccFeeLong       float64 `json:"pairAccFeeLong"`       // 1e10 (%)
	PairAccFeeShort      float64 `json:"pairAccFeeShort"`      // 1e10 (%)
}

type OiConverted struct {
	Long  float64 `json:"long"`  // 1e10 (%)
	Short float64 `json:"short"` // 1e10 (%)
	Max   float64 `json:"max"`   // 1e10 (%)
}

type OIWindowsSettingsConverted struct {
	StartTs         int `json:"startTs"`
	WindowsDuration int `json:"windowsDuration"`
	WindowsCount    int `json:"windowsCount"`
}

type OIWindowConverted struct {
	Window map[string]OIWindowDetailsConverted `json:"window"`
}

type OIWindowDetailsConverted struct {
	OILongUsd  float64 `json:"oiLongUsd"`  // 1e10 (%)
	OIShortUsd float64 `json:"oiShortUsd"` // 1e10 (%)
}

type FeeTiersConverted struct {
	Tiers       []Tier   `json:"tiers"`
	Multipliers []string `json:"multipliers"`
}

type TierConverted struct {
	FeeMultiplier   string `json:"feeMultiplier"`
	PointsThreshold string `json:"pointsThreshold"`
}

type TradeConverted struct {
	Trade          TradeDetails   `json:"trade"`
	TradeInfo      TradeInfo      `json:"tradeInfo"`
	InitialAccFees InitialAccFees `json:"initialAccFees"`
}

type TradeDetailsConverted struct {
	User             string  `json:"user"`
	Index            uint    `json:"index"`
	PairIndex        uint    `json:"pairIndex"`
	Leverage         float64 `json:"leverage"`
	Long             bool    `json:"long"`
	IsOpen           bool    `json:"isOpen"`
	CollateralIndex  uint    `json:"collateralIndex"`
	TradeType        uint    `json:"tradeType"`
	CollateralAmount float64 `json:"collateralAmount"`
	OpenPrice        float64 `json:"openPrice"`
	TP               float64 `json:"tp"`
	SL               float64 `json:"sl"`
}

type TradeInfoConverted struct {
	CreatedBlock       uint64  `json:"createdBlock"`
	TPLastUpdatedBlock uint64  `json:"tpLastUpdatedBlock"`
	SLLastUpdatedBlock uint64  `json:"slLastUpdatedBlock"`
	MaxSlippageP       float64 `json:"maxSlippageP"`
	LastOiUpdateTs     uint64  `json:"lastOiUpdateTs"`
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
}

type InitialAccFeesConverted struct {
	AccPairFee  float64 `json:"accPairFee"`  // 1e10 (%)
	AccGroupFee float64 `json:"accGroupFee"` // 1e10 (%)
	Block       uint64  `json:"block"`
}
