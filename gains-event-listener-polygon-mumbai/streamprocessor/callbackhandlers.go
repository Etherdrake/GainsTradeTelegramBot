package streamprocessor

import (
	"GainsListenerMumbai/transformers"
)

//func handleMarketCloseCanceled(event transformers.MarketCloseCanceledTransform) {
//	err := database.RemoveMarketTrade(event.OrderID)
//}
//
//func handleMarketOpenCanceled(event transformers.MarketOpenCanceledTransform) {
//
//	err := database.RemoveMarketTrade(event.OrderID)
//
//}

func handleNftOrderCanceled(event transformers.NftOrderCanceledTransform) error {
	return nil
}
