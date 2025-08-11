package altseasoncore

import (
	"fmt"
	"strconv"
)

// ToFeeConverted Method to convert Fee struct with string fields to FeeConverted struct
func (f Fee) ToFeeConverted() (FeeConverted, error) {
	openFeeP, err := parseAndDivide(f.OpenFeeP, 1e10)
	if err != nil {
		return FeeConverted{}, err
	}

	closeFeeP, err := parseAndDivide(f.CloseFeeP, 1e10)
	if err != nil {
		return FeeConverted{}, err
	}

	oracleFeeP, err := parseAndDivide(f.OracleFeeP, 1e10)
	if err != nil {
		return FeeConverted{}, err
	}

	triggerOrderFeeP, err := parseAndDivide(f.TriggerOrderFeeP, 1e10)
	if err != nil {
		return FeeConverted{}, err
	}

	minPositionSizeUsd, err := parseAndDivide(f.MinPositionSizeUsd, 1e18)
	if err != nil {
		return FeeConverted{}, err
	}

	return FeeConverted{
		OpenFeeP:           openFeeP,
		CloseFeeP:          closeFeeP,
		OracleFeeP:         oracleFeeP,
		TriggerOrderFeeP:   triggerOrderFeeP,
		MinPositionSizeUsd: minPositionSizeUsd,
	}, nil
}

// ToPairsFloat Method to convert Pairs struct to PairsConverted struct
func (p Pairs) ToPairsFloat() (PairsConverted, error) {
	oiFloat, err := p.Oi.ToOiFloat()
	if err != nil {
		return PairsConverted{}, err
	}

	feePerBlock, err := strconv.ParseFloat(p.FeePerBlock, 64)
	if err != nil {
		return PairsConverted{}, err
	}

	accFeeLong, err := strconv.ParseFloat(p.AccFeeLong, 64)
	if err != nil {
		return PairsConverted{}, err
	}

	accFeeShort, err := strconv.ParseFloat(p.AccFeeShort, 64)
	if err != nil {
		return PairsConverted{}, err
	}

	accLastUpdatedBlock, err := strconv.ParseUint(p.AccLastUpdatedBlock, 10, 64)
	if err != nil {
		return PairsConverted{}, err
	}

	feeExponent, err := strconv.ParseFloat(p.FeeExponent, 64)
	if err != nil {
		return PairsConverted{}, err
	}

	return PairsConverted{
		Oi:                  oiFloat,
		FeePerBlock:         feePerBlock,
		AccFeeLong:          accFeeLong,
		AccFeeShort:         accFeeShort,
		AccLastUpdatedBlock: accLastUpdatedBlock,
		FeeExponent:         feeExponent,
		Groups:              p.Groups,
	}, nil
}

func (g Groups) ToGroupsFloat() (GroupsConverted, error) {
	groupIndex, err := parseAndDivide(g.GroupIndex, 1)
	if err != nil {
		return GroupsConverted{}, err
	}

	block, err := strconv.ParseUint(g.Block, 10, 64)
	if err != nil {
		return GroupsConverted{}, err
	}

	initialAccFeeLong, err := parseAndDivide(g.InitialAccFeeLong, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	initialAccFeeShort, err := parseAndDivide(g.InitialAccFeeShort, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	prevGroupAccFeeLong, err := parseAndDivide(g.PrevGroupAccFeeLong, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	prevGroupAccFeeShort, err := parseAndDivide(g.PrevGroupAccFeeShort, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	pairAccFeeLong, err := parseAndDivide(g.PairAccFeeLong, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	pairAccFeeShort, err := parseAndDivide(g.PairAccFeeShort, 1e10)
	if err != nil {
		return GroupsConverted{}, err
	}

	return GroupsConverted{
		GroupIndex:           groupIndex,
		Block:                block,
		InitialAccFeeLong:    initialAccFeeLong,
		InitialAccFeeShort:   initialAccFeeShort,
		PrevGroupAccFeeLong:  prevGroupAccFeeLong,
		PrevGroupAccFeeShort: prevGroupAccFeeShort,
		PairAccFeeLong:       pairAccFeeLong,
		PairAccFeeShort:      pairAccFeeShort,
	}, nil
}

// Convert GainTradeContext to GainTradeContextConverted
func (gtc GainsTradeContext) ToConverted() GainsTradeContextConverted {
	pairs := make([]PairConverted, len(gtc.Pairs))
	for i, pair := range gtc.Pairs {
		pairs[i] = pair.ToConverted()
	}

	groups := make([]GroupConverted, len(gtc.Groups))
	for i, group := range gtc.Groups {
		groups[i] = group.ToConverted()
	}

	fees := make([]FeeConverted, len(gtc.Fees))
	for i, fee := range gtc.Fees {
		fees[i] = fee.ToConverted()
	}

	collaterals := make([]CollateralConverted, len(gtc.Collaterals))
	for i, collateral := range gtc.Collaterals {
		collaterals[i] = collateral.ToConverted()
	}

	oiWindows := make([]OIWindowConverted, len(gtc.OIWindows))
	for i, window := range gtc.OIWindows {
		oiWindows[i] = window.ToConverted()
	}

	//allTrades := make([]TradeConverted, len(gtc.AllTrades))
	//for i, trade := range gtc.AllTrades {
	//	allTrades[i] = trade.ToConverted()
	//}

	return GainsTradeContextConverted{
		LastRefreshed:             gtc.LastRefreshed,
		RefreshID:                 gtc.RefreshID,
		TradingState:              gtc.TradingState,
		MaxGainP:                  gtc.MaxGainP,
		MarketOrdersTimeoutBlocks: uint64(gtc.MarketOrdersTimeoutBlocks),
		Pairs:                     pairs,
		Groups:                    groups,
		Fees:                      fees,
		PairInfos:                 gtc.PairInfos.ToConverted(),
		Collaterals:               collaterals,
		SSSTokenBalance:           gtc.SSSTokenBalance,
		SSSLegacyTokenBalance:     gtc.SSSLegacyTokenBalance,
		SSSRewardTokens:           gtc.SSSRewardTokens,
		VaultClosingFeeP:          gtc.VaultClosingFeeP,
		MaxNegativePnlOnOpenP:     gtc.MaxNegativePnlOnOpenP,
		BlockConfirmations:        uint64(gtc.BlockConfirmations),
		OIWindowsSettings:         gtc.OIWindowsSettings.ToConverted(),
		OIWindows:                 oiWindows,
		FeeTiers:                  gtc.FeeTiers,       // TODO: This one is missing
		AllTrades:                 []TradeConverted{}, // TODO: We are ignoring AllTrades to save space
		CurrentBlock:              uint64(gtc.CurrentBlock),
		CurrentL1Block:            uint64(gtc.CurrentL1Block),
		IsForexOpen:               gtc.IsForexOpen,
		IsStocksOpen:              gtc.IsStocksOpen,
		IsIndicesOpen:             gtc.IsIndicesOpen,
		IsCommoditiesOpen:         gtc.IsCommoditiesOpen,
	}
}

// Convert Pair to PairConverted
func (p Pair) ToConverted() PairConverted {
	// Handle spreadP conversion with parseAndDivide
	spreadP, err := parseAndDivide(p.SpreadP, 12) // example factor
	if err != nil {
		spreadP = 0 // or handle the error as needed
	}

	groupIndexConv, _ := strconv.Atoi(p.GroupIndex)
	feeIndexConv, _ := strconv.Atoi(p.FeeIndex)

	return PairConverted{
		From:       p.From,
		To:         p.To,
		SpreadP:    spreadP,
		GroupIndex: groupIndexConv,
		FeeIndex:   feeIndexConv,
	}
}

// Convert Group to GroupConverted
func (g Group) ToConverted() GroupConverted {
	minLeverage, _ := strconv.ParseFloat(g.MinLeverage, 64)
	maxLeverage, _ := strconv.ParseFloat(g.MaxLeverage, 64)

	return GroupConverted{
		Name:        g.Name,
		MinLeverage: minLeverage,
		MaxLeverage: maxLeverage,
	}
}

// Convert OIWindow to OIWindowConverted
func (ow OIWindow) ToConverted() OIWindowConverted {
	convertedWindow := make(map[string]OIWindowDetailsConverted)
	for k, v := range ow.Window {
		convertedWindow[k], _ = v.ToConverted()
	}

	return OIWindowConverted{
		Window: convertedWindow,
	}
}

// Convert OIWindowDetails to OIWindowDetailsConverted
func (owd OIWindowDetails) ToConverted() (OIWindowDetailsConverted, error) {
	longUsd, err := parseAndDivide(owd.OILongUsd, 1e18)
	if err != nil {
		return OIWindowDetailsConverted{}, fmt.Errorf("failed to parse OILongUsd: %v", err)
	}

	shortUsd, err := parseAndDivide(owd.OIShortUsd, 1e18)
	if err != nil {
		return OIWindowDetailsConverted{}, fmt.Errorf("failed to parse OIShortUsd: %v", err)
	}

	return OIWindowDetailsConverted{
		OILongUsd:  longUsd,
		OIShortUsd: shortUsd,
	}, nil
}

// Convert Fee to FeeConverted
func (f Fee) ToConverted() FeeConverted {
	openFeeP, _ := parseAndDivide(f.OpenFeeP, 1e10)
	closeFeeP, _ := parseAndDivide(f.CloseFeeP, 1e10)
	oracleFeeP, _ := parseAndDivide(f.OracleFeeP, 1e10)
	triggerOrderFeeP, _ := parseAndDivide(f.TriggerOrderFeeP, 1e10)
	minPositionSizeUsd, _ := strconv.ParseFloat(f.MinPositionSizeUsd, 64)

	return FeeConverted{
		OpenFeeP:           openFeeP,
		CloseFeeP:          closeFeeP,
		OracleFeeP:         oracleFeeP,
		TriggerOrderFeeP:   triggerOrderFeeP,
		MinPositionSizeUsd: minPositionSizeUsd,
	}
}

// Convert PairInfos to PairInfosConverted
func (pi PairInfos) ToConverted() PairInfosConverted {
	pairDepths := make([]PairDepthConverted, len(pi.PairDepths))
	for i, depth := range pi.PairDepths {
		pairDepths[i] = depth.ToConverted()
	}

	return PairInfosConverted{
		MaxLeverages: pi.MaxLeverages,
		PairDepths:   pairDepths,
	}
}

// Convert PairDepth to PairDepthConverted
func (pd PairDepth) ToConverted() PairDepthConverted {
	onePercentDepthAboveUsd, _ := strconv.ParseUint(pd.OnePercentDepthAboveUsd, 10, 64)
	onePercentDepthBelowUsd, _ := strconv.ParseUint(pd.OnePercentDepthBelowUsd, 10, 64)

	return PairDepthConverted{
		OnePercentDepthAboveUsd: onePercentDepthAboveUsd,
		OnePercentDepthBelowUsd: onePercentDepthBelowUsd,
	}
}

// Convert Collateral to CollateralConverted
func (c Collateral) ToConverted() CollateralConverted {
	return CollateralConverted{
		CollateralIndex:  c.CollateralIndex,
		Collateral:       c.Collateral,
		Symbol:           c.Symbol,
		IsActive:         c.IsActive,
		Prices:           c.Prices.ToConverted(),
		CollateralConfig: c.CollateralConfig.ToConverted(),
		GToken:           c.GToken.ToConverted(),
		BorrowingFees:    c.BorrowingFees.ToConverted(),
	}
}

// Convert CollateralPrices to CollateralPricesConverted
func (cp CollateralPrices) ToConverted() CollateralPricesConverted {
	return CollateralPricesConverted{
		CollateralPriceUsd: cp.CollateralPriceUsd,
		GnsPriceCollateral: cp.GnsPriceCollateral,
		GnsPriceUsd:        cp.GnsPriceUsd,
	}
}

// Convert CollateralConfig to CollateralConfigConverted
func (cc CollateralConfig) ToConverted() CollateralConfigConverted {
	return CollateralConfigConverted{
		Precision:      cc.Precision,
		PrecisionDelta: cc.PrecisionDelta,
		Decimals:       cc.Decimals,
	}
}

// Convert GToken to GTokenConverted
func (gt GToken) ToConverted() GTokenConverted {
	return GTokenConverted{
		Address:                  gt.Address,
		CurrentBalanceCollateral: gt.CurrentBalanceCollateral,
		MaxBalanceCollateral:     gt.MaxBalanceCollateral,
		MarketCap:                gt.MarketCap,
	}
}

// Convert BorrowingFees to BorrowingFeesConverted
func (bf BorrowingFees) ToConverted() BorrowingFeesConverted {
	pairs := make([]PairsConverted, len(bf.Pairs))
	for i, pair := range bf.Pairs {
		pairs[i] = pair.ToConverted()
	}

	groups := make([]GroupsConverted, len(bf.Groups))
	for i, group := range bf.Groups {
		groups[i] = group.ToConverted()
	}

	return BorrowingFeesConverted{
		Pairs:  pairs,
		Groups: groups,
	}
}

// Convert Pairs to PairsConverted
func (p Pairs) ToConverted() PairsConverted {
	groups := make([]Groups, len(p.Groups))
	for i, group := range p.Groups {
		groups[i] = group
	}

	feePerBlock, _ := parseAndDivide(p.FeePerBlock, 10)
	accFeeLong, _ := parseAndDivide(p.AccFeeLong, 10)
	accFeeShort, _ := parseAndDivide(p.AccFeeShort, 10)
	accLastUpdatedBlock, _ := strconv.ParseUint(p.AccLastUpdatedBlock, 10, 64)
	feeExponent, _ := strconv.ParseFloat(p.FeeExponent, 64)

	return PairsConverted{
		Oi:                  p.Oi.ToConverted(),
		FeePerBlock:         feePerBlock,
		AccFeeLong:          accFeeLong,
		AccFeeShort:         accFeeShort,
		AccLastUpdatedBlock: accLastUpdatedBlock,
		FeeExponent:         feeExponent,
		Groups:              groups,
	}
}

// Convert Oi to OiConverted
func (o Oi) ToConverted() OiConverted {
	long, _ := parseAndDivide(o.Long, 1)
	short, _ := parseAndDivide(o.Short, 1)
	max, _ := parseAndDivide(o.Max, 1)

	return OiConverted{
		Long:  long,
		Short: short,
		Max:   max,
	}
}

// Convert Groups to GroupsConverted
func (g Groups) ToConverted() GroupsConverted {
	groupIndex, _ := strconv.ParseFloat(g.GroupIndex, 64)
	block, _ := strconv.ParseUint(g.Block, 10, 64)
	initialAccFeeLong, _ := parseAndDivide(g.InitialAccFeeLong, 10)
	initialAccFeeShort, _ := parseAndDivide(g.InitialAccFeeShort, 10)
	prevGroupAccFeeLong, _ := parseAndDivide(g.PrevGroupAccFeeLong, 10)
	prevGroupAccFeeShort, _ := parseAndDivide(g.PrevGroupAccFeeShort, 10)
	pairAccFeeLong, _ := parseAndDivide(g.PairAccFeeLong, 10)
	pairAccFeeShort, _ := parseAndDivide(g.PairAccFeeShort, 10)

	return GroupsConverted{
		GroupIndex:           groupIndex,
		Block:                block,
		InitialAccFeeLong:    initialAccFeeLong,
		InitialAccFeeShort:   initialAccFeeShort,
		PrevGroupAccFeeLong:  prevGroupAccFeeLong,
		PrevGroupAccFeeShort: prevGroupAccFeeShort,
		PairAccFeeLong:       pairAccFeeLong,
		PairAccFeeShort:      pairAccFeeShort,
	}
}

// Convert OIWindowsSettings to OIWindowsSettingsConverted
func (ows OIWindowsSettings) ToConverted() OIWindowsSettingsConverted {
	return OIWindowsSettingsConverted{
		StartTs:         ows.StartTs,
		WindowsDuration: ows.WindowsDuration,
		WindowsCount:    ows.WindowsCount,
	}
}
