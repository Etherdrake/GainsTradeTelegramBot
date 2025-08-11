package streamprocessor

import (
	"GainsListenerMumbai/transformers"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type UniversalOrder struct {
}

var _ sync.Mutex

// getOrderID is a helper function to retrieve OrderID from different event types using reflection
func getOrderID(event interface{}) string {
	val := reflect.ValueOf(event)
	field := val.FieldByName("OrderID")
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String()
	}
	return ""
}

func printMaps(rngToOpenLimitPlaced map[string][]transformers.OpenLimitPlacedEventTransform,
	rngToOpenLimitCanceled map[string][]transformers.OpenLimitCanceledEventTransform,
	rngToMarketOrderInitiated map[string][]transformers.MarketOrderInitiatedEventTransform,
	rngToNftOrderInitiated map[int64]transformers.NftOrderInitiatedEventTransform) (string, string, string, string) {
	var openLimitPlacedStr, openLimitCanceledStr, marketOrderInitiatedStr, nftOrderInitiatedStr strings.Builder
	// mu.RLock() // Uncomment if you decide to use locks

	openLimitPlacedStr.WriteString("rngToOpenLimitPlaced:\n")
	for trader, events := range rngToOpenLimitPlaced {
		openLimitPlacedStr.WriteString(fmt.Sprintf("%s: %v\n", trader, events))
	}

	openLimitCanceledStr.WriteString("rngToOpenLimitCanceled:\n")
	for trader, events := range rngToOpenLimitCanceled {
		openLimitCanceledStr.WriteString(fmt.Sprintf("%s: %v\n", trader, events))
	}

	marketOrderInitiatedStr.WriteString("rngToMarketOrderInitiated:\n")
	for trader, events := range rngToMarketOrderInitiated {
		marketOrderInitiatedStr.WriteString(fmt.Sprintf("%s: %v\n", trader, events))
	}

	nftOrderInitiatedStr.WriteString("rngToNftOrderInitiated:\n")
	for trader, events := range rngToNftOrderInitiated {
		nftOrderInitiatedStr.WriteString(fmt.Sprintf("%s: %v\n", trader, events))
	}

	// mu.RUnlock() // Uncomment if you decide to use locks

	return openLimitPlacedStr.String(), openLimitCanceledStr.String(), marketOrderInitiatedStr.String(), nftOrderInitiatedStr.String()
}
