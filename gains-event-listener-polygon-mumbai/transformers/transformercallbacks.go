package transformers

import (
	"GainsListenerMumbai/eventdatatypes"
	"GainsListenerMumbai/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

func TransformMarketExecuted(eventLog types.Log, decodedEvent *eventdatatypes.MarketExecuted, collToken string) MarketExecutedTransform {
	// Retrieve the TradeID which is topic[1]
	OrderID := eventLog.Topics[1]

	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)
	sixFloat := big.NewFloat(1e6)

	var positionSizeDaiF64 float64

	// USDC
	if collToken == "USDC" {
		positionSizeDai := new(big.Float).SetInt(decodedEvent.PositionSizeDai)
		positionSizeDai.Quo(positionSizeDai, sixFloat)
		positionSizeDaiF64, _ = positionSizeDai.Float64()
	} else {
		positionSizeDai := new(big.Float).SetInt(decodedEvent.PositionSizeDai)
		positionSizeDai.Quo(positionSizeDai, eighteenFloat)
		positionSizeDaiF64, _ = positionSizeDai.Float64()
	}

	initialPosToken := new(big.Float).SetInt(decodedEvent.T.InitialPosToken)
	initialPosToken.Quo(initialPosToken, eighteenFloat)
	initialPosTokenF64, _ := initialPosToken.Float64()

	openPrice := new(big.Float).SetInt(decodedEvent.T.OpenPrice)
	openPrice.Quo(openPrice, tenFloat)
	openPriceFloat64, _ := openPrice.Float64()
	tp := new(big.Float).SetInt(decodedEvent.T.Tp)
	tp.Quo(tp, tenFloat)
	tpFloat64, _ := tp.Float64()
	sl := new(big.Float).SetInt(decodedEvent.T.Sl)
	sl.Quo(sl, tenFloat)
	slFloat64, _ := sl.Float64()

	trade := TradeTransform{
		Trader:          utils.LowercaseString(decodedEvent.T.Trader.String()),
		PairIndex:       decodedEvent.T.PairIndex.Int64(),
		Index:           decodedEvent.T.Index.Int64(),
		InitialPosToken: initialPosTokenF64,
		PositionSizeDai: positionSizeDaiF64,
		OpenPrice:       openPriceFloat64,
		Buy:             decodedEvent.T.Buy,
		Leverage:        decodedEvent.T.Leverage.Int64(),
		Tp:              tpFloat64,
		Sl:              slFloat64,
		CollateralToken: collToken,
	}

	price := new(big.Float).SetInt(decodedEvent.Price)
	price.Quo(price, tenFloat)
	priceFloat64, _ := price.Float64()

	priceImpactP := new(big.Float).SetInt(decodedEvent.PriceImpactP)
	priceImpactP.Quo(priceImpactP, tenFloat)
	priceImpactPFloat64, _ := priceImpactP.Float64()

	positionSizeDaiOuter := new(big.Float).SetInt(decodedEvent.PositionSizeDai)
	positionSizeDaiOuter.Quo(positionSizeDaiOuter, eighteenFloat)
	positionSizeDaiFloat64, _ := positionSizeDaiOuter.Float64()

	percentProfit := new(big.Float).SetInt(decodedEvent.PercentProfit)
	percentProfit.Quo(percentProfit, tenFloat)
	percentProfitFloat64, _ := percentProfit.Float64()

	daiSentToTrader := new(big.Float).SetInt(decodedEvent.DaiSentToTrader)
	daiSentToTrader.Quo(eighteenFloat, eighteenFloat)
	daiSentToTraderFloat64, _ := daiSentToTrader.Float64()

	// Convert hash to int64
	int64Value, err := HexStringToInt64(OrderID.String())
	if err != nil {
		fmt.Println("Error:", err)
	}

	transform := MarketExecutedTransform{
		OrderID:         int64Value,
		Trade:           trade,
		Open:            decodedEvent.Open,
		Price:           priceFloat64,
		PriceImpactP:    priceImpactPFloat64,
		PositionSizeDai: positionSizeDaiFloat64,
		PercentProfit:   percentProfitFloat64,
		DaiSentToTrader: daiSentToTraderFloat64,
	}
	return transform
}

func TransformLimitExecuted(eventLog types.Log, decodedEvent *eventdatatypes.LimitExecuted, collToken string) LimitExecutedTransform {
	// Retrieve the TradeID which is topic[1]
	OrderID := eventLog.Topics[1]

	//eighteenInt := big.NewInt(1e18)
	//tenInt := big.NewInt(1e10)

	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)
	sixFloat := big.NewFloat(1e6)

	var positionSizeDaiF64 float64

	fmt.Println("collToken: ", collToken)

	// USDC
	if collToken == "USDC" {
		positionSizeDai := new(big.Float).SetInt(decodedEvent.Trade.PositionSizeDai)
		positionSizeDai.Quo(positionSizeDai, sixFloat)
		positionSizeDaiF64, _ = positionSizeDai.Float64()
	} else {
		positionSizeDai := new(big.Float).SetInt(decodedEvent.Trade.PositionSizeDai)
		positionSizeDai.Quo(positionSizeDai, eighteenFloat)
		positionSizeDaiF64, _ = positionSizeDai.Float64()
	}

	initialPosToken := new(big.Float).SetInt(decodedEvent.Trade.InitialPosToken)
	initialPosToken.Quo(initialPosToken, eighteenFloat)
	initialPosTokenF64, _ := initialPosToken.Float64()

	openPrice := new(big.Float).SetInt(decodedEvent.Trade.OpenPrice)
	openPrice.Quo(openPrice, tenFloat)
	openPriceFloat64, _ := openPrice.Float64()

	tp := new(big.Float).SetInt(decodedEvent.Trade.Tp)
	tp.Quo(tp, tenFloat)
	tpFloat64, _ := tp.Float64()

	sl := new(big.Float).SetInt(decodedEvent.Trade.Sl)
	sl.Quo(sl, tenFloat)
	slFloat64, _ := sl.Float64()

	trade := TradeTransform{
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

	price := new(big.Float).SetInt(decodedEvent.Price)
	price.Quo(price, tenFloat)
	priceFloat64, _ := price.Float64()

	priceImpactP := new(big.Float).SetInt(decodedEvent.PriceImpactP)
	priceImpactP.Quo(priceImpactP, tenFloat)
	priceImpactPFloat64, _ := priceImpactP.Float64()

	positionSizeDaiOuter := new(big.Float).SetInt(decodedEvent.PositionSizeDai)
	positionSizeDaiOuter.Quo(positionSizeDaiOuter, eighteenFloat)
	positionSizeDaiFloat64, _ := positionSizeDaiOuter.Float64()

	percentProfit := new(big.Float).SetInt(decodedEvent.PercentProfit)
	percentProfit.Quo(percentProfit, tenFloat)
	percentProfitFloat64, _ := percentProfit.Float64()

	daiSentToTrader := new(big.Float).SetInt(decodedEvent.DaiSentToTrader)
	daiSentToTrader.Quo(daiSentToTrader, eighteenFloat)
	daiSentToTraderFloat64, _ := daiSentToTrader.Float64()

	// Convert hash to int64
	orderId, err := HexStringToInt64(OrderID.String())
	if err != nil {
		fmt.Println("Error:", err)
	}

	transform := LimitExecutedTransform{
		OrderID:         orderId,
		LimitIndex:      decodedEvent.LimitIndex.Int64(),
		Trade:           trade,
		NftHolder:       eventLog.Topics[2].Hex(),
		OrderType:       decodedEvent.OrderType,
		Price:           priceFloat64,
		PriceImpactP:    priceImpactPFloat64,
		PositionSizeDai: positionSizeDaiFloat64,
		PercentProfit:   percentProfitFloat64,
		DaiSentToTrader: daiSentToTraderFloat64,
		ExactExecution:  false,
		CollateralToken: collToken,
	}
	return transform
}

func TransformMarketOpenCanceled(eventLog types.Log, decodedEvent *eventdatatypes.MarketOpenCanceled, collToken string) MarketOpenCanceledTransform {
	return MarketOpenCanceledTransform{
		OrderID:         int64(decodedEvent.OrderID),
		Trader:          decodedEvent.Trader,
		PairIndex:       decodedEvent.PairIndex,
		CancelReason:    decodedEvent.CancelReason,
		CollateralToken: collToken,
	}
}

func TransformMarketCloseCanceled(eventLog types.Log, decodedEvent *eventdatatypes.MarketCloseCanceled, collToken string) MarketCloseCanceledTransform {
	return MarketCloseCanceledTransform{
		OrderID:         int64(decodedEvent.OrderID),
		Trader:          decodedEvent.Trader,
		PairIndex:       decodedEvent.PairIndex,
		Index:           decodedEvent.Index.Uint64(),
		CancelReason:    decodedEvent.CancelReason,
		CollateralToken: collToken,
	}
}

func TransformNftOrderCanceled(eventLog types.Log, decodedEvent *eventdatatypes.NftOrderCanceled, collToken string) NftOrderCanceledTransform {
	return NftOrderCanceledTransform{
		OrderID:         int64(decodedEvent.OrderID),
		NftHolder:       decodedEvent.NftHolder,
		OrderType:       decodedEvent.OrderType,
		CancelReason:    decodedEvent.CancelReason,
		CollateralToken: collToken,
	}
}
func HexStringToIntString(hexStr string) (string, error) {
	// Convert hex string to big.Int
	hexInt, success := new(big.Int).SetString(hexStr, 16)
	if !success {
		return "", fmt.Errorf("invalid hex string: %s", hexStr)
	}

	// Convert big.Int to string
	intStr := hexInt.String()

	return intStr, nil
}

func HashToInt64(hash common.Hash) (int64, error) {
	// Convert the hash to a big.Int
	hashBigInt := new(big.Int)
	hashBigInt.SetString(strings.TrimPrefix(hash.Hex(), "0x"), 16)

	// Extract the int64 value
	int64Value := hashBigInt.Int64()

	return int64Value, nil
}

func HexStringToInt64(hexStr string) (int64, error) {
	// Remove the "0x" prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Convert the hex string to a big.Int
	hexBigInt, success := new(big.Int).SetString(hexStr, 16)
	if !success {
		return 0, fmt.Errorf("invalid hex string: %s", hexStr)
	}

	// Extract the int64 value
	int64Value := hexBigInt.Int64()

	return int64Value, nil
}
