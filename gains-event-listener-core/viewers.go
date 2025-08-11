package main

import (
	"GainsListenerV4/types"
	"fmt"
)

func ViewOpenTrades(openTrades types.OpenTrades) {
	if len(openTrades.Value) > 0 {
		firstTrade := openTrades.Value[0]

		fmt.Println("Trade details:")
		fmt.Println("User:", firstTrade.Trade.User)
		fmt.Println("Index:", firstTrade.Trade.Index)
		fmt.Println("Pair Index:", firstTrade.Trade.PairIndex)
		fmt.Println("Leverage:", firstTrade.Trade.Leverage)
		fmt.Println("Long:", firstTrade.Trade.Long)
		fmt.Println("Is Open:", firstTrade.Trade.IsOpen)
		fmt.Println("Collateral Index:", firstTrade.Trade.CollateralIndex)
		fmt.Println("Trade Type:", firstTrade.Trade.TradeType)
		fmt.Println("Collateral Amount:", firstTrade.Trade.CollateralAmount)
		fmt.Println("Open Price:", firstTrade.Trade.OpenPrice)
		fmt.Println("TP:", firstTrade.Trade.Tp)
		fmt.Println("SL:", firstTrade.Trade.Sl)

		fmt.Println("Trade Info:")
		fmt.Println("Created Block:", firstTrade.TradeInfo.CreatedBlock)
		fmt.Println("TP Last Updated Block:", firstTrade.TradeInfo.TpLastUpdatedBlock)
		fmt.Println("SL Last Updated Block:", firstTrade.TradeInfo.SlLastUpdatedBlock)
		fmt.Println("Max Slippage P:", firstTrade.TradeInfo.MaxSlippageP)
		fmt.Println("Last OI Update Ts:", firstTrade.TradeInfo.LastOiUpdateTs)
		fmt.Println("Collateral Price USD:", firstTrade.TradeInfo.CollateralPriceUsd)

		fmt.Println("Initial Acc Fees:")
		fmt.Println("Acc Pair Fee:", firstTrade.InitialAccFees.AccPairFee)
		fmt.Println("Acc Group Fee:", firstTrade.InitialAccFees.AccGroupFee)
		fmt.Println("Block:", firstTrade.InitialAccFees.Block)
	} else {
		fmt.Println("No open trades found.")
	}
}
