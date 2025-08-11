package logger

import (
	"GainsListenerMumbai/eventdatatypes"
	"GainsListenerMumbai/pairmaps"
	"GainsListenerMumbai/priceserver"
	"GainsListenerMumbai/transformers"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

// Function to format MarketExecuted event as a string
func GenerateMarketExecutedLog(txHash string, event *eventdatatypes.MarketExecuted) string {
	pairInt := int(event.T.PairIndex.Int64())
	return fmt.Sprintf("MarketExecuted Decoded:\n"+
		"TxHash: %s\n"+
		"Trader: %s\n"+
		"PairIndex: %s"+" PAIR: "+pairmaps.IndexToPair[pairInt]+"\n"+
		"Index: %s\n"+
		"InitialPosToken: %s\n"+
		"PositionSizeDai: %s\n"+
		"OpenPrice: %s\n"+
		"Buy: %v\n"+
		"Leverage: %s\n"+
		"Tp: %s\n"+
		"Sl: %s\n"+
		"Open: %v\n"+
		"Price: %s\n"+
		"PriceImpactP: %s\n"+
		"PositionSizeDai: %s\n"+
		"PercentProfit: %s\n"+
		"DaiSentToTrader: %s",
		txHash,
		event.T.Trader.Hex(),
		event.T.PairIndex.String(),
		event.T.Index.String(),
		event.T.InitialPosToken.String(),
		event.T.PositionSizeDai.String(),
		event.T.OpenPrice.String(),
		event.T.Buy,
		event.T.Leverage.String(),
		event.T.Tp.String(),
		event.T.Sl.String(),
		event.Open,
		event.Price.String(),
		event.PriceImpactP.String(),
		event.PositionSizeDai.String(),
		event.PercentProfit.String(),
		event.DaiSentToTrader.String(),
	)
}

// Function to format MarketExecuted event as a string
func GenerateTransformedMarketExecutedLog(rdbPrice *redis.Client, txHash string, event transformers.MarketExecutedTransform) string {
	price, err := priceserver.GetPrice(rdbPrice, int(event.Trade.PairIndex))
	if err != nil {
		log.Println("Error getting price.")
	}

	pairInt := int(event.Trade.PairIndex)
	return fmt.Sprintf("MarketExecuted Transformed:\n"+
		"TxHash: %s\n"+
		"Trader: %s\n"+
		"PairIndex: %d"+" PAIR: "+pairmaps.IndexToPair[pairInt]+"\n"+
		"Index: %d\n"+
		"InitialPosToken: %.2f\n"+
		"PositionSizeDai: %.2f\n"+
		"OpenPrice: %f"+" CURRENT PRICE: "+"%f"+"\n"+
		"Buy: %v\n"+
		"Leverage: %d\n"+
		"Tp: %f\n"+
		"Sl: %f\n"+
		"Open: %v\n"+
		"Price: %f\n"+
		"PriceImpactP: %f\n"+
		"PositionSizeDai: %.2f\n"+
		"PercentProfit: %.2f\n"+
		"DaiSentToTrader: %.2f"+
		txHash,
		event.Trade.Trader,
		event.Trade.PairIndex,
		event.Trade.Index,
		event.Trade.InitialPosToken,
		event.Trade.PositionSizeDai,
		event.Trade.OpenPrice,
		price,
		event.Trade.Buy,
		event.Trade.Leverage,
		event.Trade.Tp,
		event.Trade.Sl,
		event.Open,
		event.Price,
		event.PriceImpactP,
		event.PositionSizeDai,
		event.PercentProfit,
		event.DaiSentToTrader,
	)
}

// Function to format LimitExecuted event as a string
func GenerateLimitExecutedLog(txHash string, event *eventdatatypes.LimitExecuted) string {
	pairInt := int(event.Trade.PairIndex.Int64())
	return fmt.Sprintf("LimitExecuted Decoded:\n"+
		"TxHash: %s\n"+
		"OrderID: %s\n"+
		"LimitIndex: %s\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+" PAIR: "+pairmaps.IndexToPair[pairInt]+"\n"+
		"Index: %s\n"+
		"InitialPosToken: %s\n"+
		"PositionSizeDai: %s\n"+
		"OpenPrice: %s\n"+
		"Buy: %v\n"+
		"Leverage: %s\n"+
		"Tp: %s\n"+
		"Sl: %s\n"+
		"Price: %s\n"+
		"PriceImpactP: %s\n"+
		"PositionSizeDai: %s\n"+
		"PercentProfit: %s\n"+
		"DaiSentToTrader: %s\n"+
		"ExactExecution: %v",
		txHash,
		event.OrderID,
		event.LimitIndex.String(),
		event.Trade.Trader.Hex(),
		event.Trade.PairIndex.String(),
		event.Trade.Index.String(),
		event.Trade.InitialPosToken.String(),
		event.Trade.PositionSizeDai.String(),
		event.Trade.OpenPrice.String(),
		event.Trade.Buy,
		event.Trade.Leverage.String(),
		event.Trade.Tp.String(),
		event.Trade.Sl.String(),
		event.Price.String(),
		event.PriceImpactP.String(),
		event.PositionSizeDai.String(),
		event.PercentProfit.String(),
		event.DaiSentToTrader.String(),
		event.ExactExecution,
	)
}

// Function to format LimitExecutedTransform event as a string
func GenerateLimitExecutedTransformLog(rdbPrice *redis.Client, txHash string, event *transformers.LimitExecutedTransform) string {
	price, err := priceserver.GetPrice(rdbPrice, int(event.Trade.PairIndex))
	if err != nil {
		log.Println("Error getting price.")
	}

	pairInt := int(event.Trade.PairIndex)

	return fmt.Sprintf("LimitExecutedTransform Decoded:\n"+
		"TxHash: %s\n"+
		"OrderID: %d\n"+
		"LimitIndex: %d\n"+
		"Trader: %s\n"+
		"PairIndex: %d"+" PAIR: "+pairmaps.IndexToPair[pairInt]+"\n"+
		"Index: %d\n"+
		"InitialPosToken: %f\n"+
		"PositionSizeDai: %f\n"+
		"OpenPrice: %f"+" CURRENT PRICE: "+"%f"+"\n"+
		"Buy: %v\n"+
		"Leverage: %d\n"+
		"Tp: %f\n"+
		"Sl: %f\n"+
		"Price: %f\n"+
		"PriceImpactP: %f\n"+
		"PositionSizeDai: %f\n"+
		"PercentProfit: %f\n"+
		"DaiSentToTrader: %f\n"+
		"ExactExecution: %v",
		txHash,
		event.OrderID,
		event.LimitIndex,
		event.Trade.Trader,
		event.Trade.PairIndex,
		event.Trade.Index,
		event.Trade.InitialPosToken,
		event.Trade.PositionSizeDai,
		event.Trade.OpenPrice,
		price,
		event.Trade.Buy,
		event.Trade.Leverage,
		event.Trade.Tp,
		event.Trade.Sl,
		event.Price,
		event.PriceImpactP,
		event.PositionSizeDai,
		event.PercentProfit,
		event.DaiSentToTrader,
		event.ExactExecution,
	)
}

// Function to format MarketOpenCanceled event as a string
func GenerateMarketOpenCanceledLog(txHash string, event *eventdatatypes.MarketOpenCanceled) string {
	return fmt.Sprintf("MarketOpenCanceled Transformed:\n"+
		"OrderID: %d\n"+
		"Trader: %s\n"+
		"PairIndex: %d\n"+
		"CancelReason: %d\n",
		event.OrderID,
		event.Trader,
		event.PairIndex,
		event.CancelReason,
	)
}

// Function to format MarketCloseCanceled event as a string
func GenerateMarketCloseCanceledLog(txHash string, event *eventdatatypes.MarketCloseCanceled) string {
	return fmt.Sprintf("MarketCloseCanceled Transformed:\n"+
		"OrderID: %d\n"+
		"Trader: %d\n"+
		"PairIndex: %d\n"+
		"Index: %s\n"+
		"CancelReason: %d\n",
		event.OrderID,
		event.Trader,
		event.PairIndex,
		event.Index.String(),
		event.CancelReason,
	)
}

// Function to format NftOrderCanceled event as a string
func GenerateNftOrderCanceledLog(txHash string, event *eventdatatypes.NftOrderCanceled) string {
	return fmt.Sprintf("NftOrderCanceled Transformed:\n"+
		"OrderID: %d\n"+
		"NftHolder: %s\n"+
		"OrderType: %d\n"+
		"CancelReason: %d\n",
		event.OrderID,
		event.NftHolder,
		event.OrderType,
		event.CancelReason,
	)
}

func GeneratePairMaxLeverageUpdatedLog(txHash string, event *eventdatatypes.PairMaxLeverageUpdated) string {
	return fmt.Sprintf("PairMaxLeverageUpdated Transformed:\n"+
		"MaxLeverageL: %s\n",
		event.MaxLeverage.String(),
	)
}
