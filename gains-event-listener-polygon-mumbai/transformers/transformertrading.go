package transformers

import (
	"GainsListenerMumbai/eventdatatypes"
	"GainsListenerMumbai/utils"
	"math/big"
	"strconv"
)

// DoneEvent transform function
func TransformDoneEvent(decodedEvent eventdatatypes.DoneEvent, collToken string) *DoneEventTransform {
	return &DoneEventTransform{
		Done:            decodedEvent.Done,
		CollateralToken: collToken,
	}
}

// PausedEvent transform function
func TransformPausedEvent(decodedEvent eventdatatypes.PausedEvent, collToken string) *PausedEventTransform {
	return &PausedEventTransform{
		Paused:          decodedEvent.Paused,
		CollateralToken: collToken,
	}
}

// NumberUpdatedEvent transform function
func TransformNumberUpdatedEvent(decodedEvent eventdatatypes.NumberUpdatedEvent, collToken string) *NumberUpdatedEventTransform {
	return &NumberUpdatedEventTransform{
		Name:            decodedEvent.Name,
		Value:           decodedEvent.Value.Int64(),
		CollateralToken: collToken,
	}
}

// BypassTriggerLinkUpdatedEvent transform function
func TransformBypassTriggerLinkUpdatedEvent(decodedEvent eventdatatypes.BypassTriggerLinkUpdatedEvent, collToken string) *BypassTriggerLinkUpdatedEventTransform {
	return &BypassTriggerLinkUpdatedEventTransform{
		User:            decodedEvent.User,
		Bypass:          decodedEvent.Bypass,
		CollateralToken: collToken,
	}
}

// MarketOrderInitiatedEvent transform function
func TransformMarketOrderInitiatedEvent(decodedEvent eventdatatypes.MarketOrderInitiatedEvent, collToken string) *MarketOrderInitiatedEventTransform {

	return &MarketOrderInitiatedEventTransform{
		OrderID:         decodedEvent.OrderID,
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Open:            decodedEvent.Open,
		CollateralToken: collToken,
	}
}

// OpenLimitPlacedEvent transform function
func TransformOpenLimitPlacedEvent(decodedEvent eventdatatypes.OpenLimitPlacedEvent, collToken string) *OpenLimitPlacedEventTransform {
	return &OpenLimitPlacedEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		CollateralToken: collToken,
	}
}

// OpenLimitUpdatedEvent transform function
func TransformOpenLimitUpdatedEvent(decodedEvent eventdatatypes.OpenLimitUpdatedEvent, collToken string) *OpenLimitUpdatedEventTransform {
	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)

	newPrice := new(big.Float).SetInt(decodedEvent.NewPrice)
	newPrice.Quo(newPrice, eighteenFloat) // Replace someOtherFloat with the appropriate divisor
	newPriceFloat64, _ := newPrice.Float64()

	newTp := new(big.Float).SetInt(decodedEvent.NewTp)
	newTp.Quo(newTp, tenFloat) // Replace anotherFloat with the appropriate divisor
	newTpFloat64, _ := newTp.Float64()

	newSl := new(big.Float).SetInt(decodedEvent.NewSl)
	newSl.Quo(newSl, tenFloat) // Replace yetAnotherFloat with the appropriate divisor
	newSlFloat64, _ := newSl.Float64()

	maxSlippageP := new(big.Float).SetInt(decodedEvent.MaxSlippageP)
	maxSlippageP.Quo(maxSlippageP, tenFloat) // Replace finalFloat with the appropriate divisor
	maxSlippagePFloat64, _ := maxSlippageP.Float64()

	return &OpenLimitUpdatedEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		NewPrice:        newPriceFloat64,
		NewTp:           newTpFloat64,
		NewSl:           newSlFloat64,
		MaxSlippageP:    maxSlippagePFloat64,
		CollateralToken: collToken,
	}
}

// OpenLimitCanceledEvent transform function
func TransformOpenLimitCanceledEvent(decodedEvent eventdatatypes.OpenLimitCanceledEvent, collToken string) *OpenLimitCanceledEventTransform {
	return &OpenLimitCanceledEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		CollateralToken: collToken,
	}
}

// TpUpdatedEvent transform function
func TransformTpUpdatedEvent(decodedEvent eventdatatypes.TpUpdatedEvent, collToken string) *TpUpdatedEventTransform {
	//eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)

	newTp := new(big.Float).SetInt(decodedEvent.NewTp)
	newTp.Quo(newTp, tenFloat) // Replace anotherFloat with the appropriate divisor
	newTpFloat64, _ := newTp.Float64()

	return &TpUpdatedEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		NewTp:           newTpFloat64,
		CollateralToken: collToken,
	}
}

// SlUpdatedEvent transform function
func TransformSlUpdatedEvent(decodedEvent eventdatatypes.SlUpdatedEvent, collToken string) *SlUpdatedEventTransform {
	//eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)

	newSl := new(big.Float).SetInt(decodedEvent.NewSl)
	newSl.Quo(newSl, tenFloat) // Replace yetAnotherFloat with the appropriate divisor
	newSlFloat64, _ := newSl.Float64()

	return &SlUpdatedEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		NewSl:           newSlFloat64,
		CollateralToken: collToken,
	}
}

// NftOrderInitiatedEvent transform function
func TransformNftOrderInitiatedEvent(decodedEvent eventdatatypes.NftOrderInitiatedEvent, collToken string) *NftOrderInitiatedEventTransform {
	orderID, _ := strconv.ParseInt(decodedEvent.OrderID, 10, 64)

	return &NftOrderInitiatedEventTransform{
		OrderID:          orderID,
		Trader:           utils.LowercaseString(decodedEvent.Trader),
		PairIndex:        decodedEvent.PairIndex.Int64(),
		ByPassesLinkCost: decodedEvent.ByPassesLinkCost,
		CollateralToken:  collToken,
	}
}

//// ChainlinkCallbackTimeoutEvent transform function
//func TransformChainlinkCallbackTimeoutEvent(decodedEvent eventdatatypes.ChainlinkCallbackTimeoutEvent, collToken string) *ChainlinkCallbackTimeoutEventTransform {
//	return &ChainlinkCallbackTimeoutEventTransform{
//		OrderID: decodedEvent.OrderID.Int64(),
//		Order:   TransformPendingMarketOrder(decodedEvent.Order),
//	}
//}

// CouldNotCloseTradeEvent transform function
func TransformCouldNotCloseTradeEvent(decodedEvent eventdatatypes.CouldNotCloseTradeEvent, collToken string) *CouldNotCloseTradeEventTransform {
	return &CouldNotCloseTradeEventTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trader),
		PairIndex:       decodedEvent.PairIndex.Int64(),
		Index:           decodedEvent.Index.Int64(),
		CollateralToken: collToken,
	}
}

// PendingMarketOrder transform function
func TransformPendingMarketOrder(decodedEvent eventdatatypes.PendingMarketOrder, collToken string) *PendingMarketOrderTransform {
	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)

	initialPosToken := new(big.Float).SetInt(decodedEvent.Trade.InitialPosToken)
	initialPosToken.Quo(initialPosToken, eighteenFloat)
	initialPosTokenF64, _ := initialPosToken.Float64()

	positionSizeDai := new(big.Float).SetInt(decodedEvent.Trade.PositionSizeDai)
	positionSizeDai.Quo(positionSizeDai, eighteenFloat)
	positionSizeDaiF64, _ := positionSizeDai.Float64()

	openPrice := new(big.Float).SetInt(decodedEvent.Trade.OpenPrice)
	openPrice.Quo(openPrice, tenFloat)
	openPriceFloat64, _ := openPrice.Float64()

	tp := new(big.Float).SetInt(decodedEvent.Trade.Tp)
	tp.Quo(tp, tenFloat)
	tpFloat64, _ := tp.Float64()
	sl := new(big.Float).SetInt(decodedEvent.Trade.Sl)
	sl.Quo(sl, tenFloat)
	slFloat64, _ := sl.Float64()

	Trade := TradeTransform{
		Trader:          utils.LowercaseString(decodedEvent.Trade.Trader.String()),
		PairIndex:       decodedEvent.Trade.PairIndex.Int64(),
		Index:           decodedEvent.Trade.Index.Int64(),
		InitialPosToken: initialPosTokenF64,
		PositionSizeDai: positionSizeDaiF64,
		OpenPrice:       openPriceFloat64,
		Buy:             decodedEvent.Trade.Buy,
		Leverage:        decodedEvent.Trade.Leverage.Int64(),
		Tp:              tpFloat64,
		Sl:              slFloat64,
		CollateralToken: collToken,
	}

	wantedPrice := new(big.Float).SetInt(decodedEvent.WantedPrice)
	wantedPrice.Quo(wantedPrice, tenFloat) // Replace someFloat with the appropriate divisor
	wantedPriceFloat64, _ := wantedPrice.Float64()

	slippageP := new(big.Float).SetInt(decodedEvent.SlippageP)
	slippageP.Quo(slippageP, tenFloat) // Replace anotherFloat with the appropriate divisor
	slippagePFloat64, _ := slippageP.Float64()

	spreadReductionP := new(big.Float).SetInt(decodedEvent.SpreadReductionP)
	spreadReductionP.Quo(spreadReductionP, tenFloat) // Replace yetAnotherFloat with the appropriate divisor
	spreadReductionPFloat64, _ := spreadReductionP.Float64()

	return &PendingMarketOrderTransform{
		Trade:            Trade,
		Block:            decodedEvent.Block.Int64(),
		WantedPrice:      wantedPriceFloat64,
		SlippageP:        slippagePFloat64,
		SpreadReductionP: spreadReductionPFloat64,
		TokenID:          int(decodedEvent.TokenID.Int64()),
		CollateralToken:  collToken,
	}
}
