package sniping

var SniperSettingsGlobal struct {
	// Sniping settings
	MultiWallet   bool
	AntiRug       bool
	AntiMev       bool
	AutoSlip      bool
	AntiBlacklist bool
	PreApprove    bool
	Slippage      float64
	MaxGas        int
	ExtraGas      int
	ApproveGas    float64
	BundleTips    bool
	Tip           float64
	BlockNumber   int
}

var SniperBuySettings struct {
	// Sniping settings for buying
	AutoBuy          bool
	BuyTwice         bool
	BuyExtraGas      int64
	MaxMarketCap     float64
	MinLiquidity     float64
	MaxLiquidity     float64
	MinimumMarketCap float64
	MaxBuyTax        float64
	MaxSellTax       float64
}

var SniperSellSettings struct {
	// Sniping settings for selling
	AutoSell     bool
	ConfirmSell  bool
	SellExtraGas int64
	TProfitOne   int
	TProfitTwo   int
	AutoSellOne  int
	AutoSellTwo  int
	StoplossSell int
	MaxSellTax   int
}
