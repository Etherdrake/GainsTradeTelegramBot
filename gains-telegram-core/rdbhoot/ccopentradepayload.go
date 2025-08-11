package rdbhoot

import "HootCore/api/tradinginteractions"

func (cpCache HootCoreCache) CreateOpenTradePayload() (payload tradinginteractions.OpenTradePayload) {
	tradePayload := tradinginteractions.TradePayload{
		User:             cpCache.Trader,
		Index:            uint32(cpCache.Index),
		PairIndex:        uint16(cpCache.PairIndex),
		Leverage:         cpCache.Leverage,
		Long:             cpCache.Buy,
		IsOpen:           cpCache.Buy,
		CollateralIndex:  cpCache.Collateral,
		TradeType:        cpCache.OrderType,
		CollateralAmount: cpCache.PositionSizeDai,
		OpenPrice:        cpCache.OpenPrice,
		TP:               cpCache.TP,
		SL:               cpCache.SL,
		Placeholder:      0,
	}

	openTradePayload := tradinginteractions.OpenTradePayload{
		Guid:            cpCache.ID,
		Trade:           tradePayload,
		Chain:           cpCache.Chain,
		CollateralToken: cpCache.Collateral,
	}

	return openTradePayload
}
