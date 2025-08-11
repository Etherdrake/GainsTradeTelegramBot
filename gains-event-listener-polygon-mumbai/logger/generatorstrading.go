package logger

import (
	"GainsListenerMumbai/eventdatatypes"
	"fmt"
)

// Function to format DoneEvent as a string
func GenerateDoneEventLog(txHash string, event *eventdatatypes.DoneEvent) string {
	return fmt.Sprintf("DoneEvent Transformed:\n"+
		"Done: %t\n",
		event.Done,
	)
}

// Function to format PausedEvent as a string
func GeneratePausedEventLog(txHash string, event *eventdatatypes.PausedEvent) string {
	return fmt.Sprintf("PausedEvent Transformed:\n"+
		"Paused: %t\n",
		event.Paused,
	)
}

// Function to format NumberUpdatedEvent as a string
func GenerateNumberUpdatedEventLog(txHash string, event *eventdatatypes.NumberUpdatedEvent) string {
	return fmt.Sprintf("NumberUpdatedEvent Transformed:\n"+
		"Name: %s\n"+
		"Value: %s\n",
		event.Name,
		event.Value.String(),
	)
}

// Function to format BypassTriggerLinkUpdatedEvent as a string
func GenerateBypassTriggerLinkUpdatedEventLog(txHash string, event *eventdatatypes.BypassTriggerLinkUpdatedEvent) string {
	return fmt.Sprintf("BypassTriggerLinkUpdatedEvent Transformed:\n"+
		"User: %s\n"+
		"Bypass: %t\n",
		event.User,
		event.Bypass,
	)
}

// Function to format MarketOrderInitiatedEvent as a string
func GenerateMarketOrderInitiatedEventLog(txHash string, event *eventdatatypes.MarketOrderInitiatedEvent) string {
	return fmt.Sprintf("MarketOrderInitiatedEvent Transformed:\n"+
		"OrderID: %s\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Open: %t\n",
		event.OrderID,
		event.Trader,
		event.PairIndex.String(),
		event.Open,
	)
}

// Function to format OpenLimitPlacedEvent as a string
func GenerateOpenLimitPlacedEventLog(txHash string, event *eventdatatypes.OpenLimitPlacedEvent) string {
	return fmt.Sprintf("OpenLimitPlacedEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
	)
}

// Function to format OpenLimitUpdatedEvent as a string
func GenerateOpenLimitUpdatedEventLog(txHash string, event *eventdatatypes.OpenLimitUpdatedEvent) string {
	return fmt.Sprintf("OpenLimitUpdatedEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n"+
		"NewPrice: %s\n"+
		"NewTp: %s\n"+
		"NewSl: %s\n"+
		"MaxSlippageP: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
		event.NewPrice.String(),
		event.NewTp.String(),
		event.NewSl.String(),
		event.MaxSlippageP.String(),
	)
}

// Function to format OpenLimitCanceledEvent as a string
func GenerateOpenLimitCanceledEventLog(txHash string, event *eventdatatypes.OpenLimitCanceledEvent) string {
	return fmt.Sprintf("OpenLimitCanceledEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
	)
}

// Function to format TpUpdatedEvent as a string
func GenerateTpUpdatedEventLog(txHash string, event *eventdatatypes.TpUpdatedEvent) string {
	return fmt.Sprintf("TpUpdatedEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n"+
		"NewTp: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
		event.NewTp.String(),
	)
}

// Function to format SlUpdatedEvent as a string
func GenerateSlUpdatedEventLog(txHash string, event *eventdatatypes.SlUpdatedEvent) string {
	return fmt.Sprintf("SlUpdatedEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n"+
		"NewSl: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
		event.NewSl.String(),
	)
}

// Function to format NftOrderInitiatedEvent as a string
func GenerateNftOrderInitiatedEventLog(txHash string, event *eventdatatypes.NftOrderInitiatedEvent) string {
	return fmt.Sprintf("NftOrderInitiatedEvent Transformed:\n"+
		"OrderID: %s\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"ByPassesLinkCost: %t\n",
		event.OrderID,
		event.Trader,
		event.PairIndex.String(),
		event.ByPassesLinkCost,
	)
}

// Function to format ChainlinkCallbackTimeoutEvent as a string
func GenerateChainlinkCallbackTimeoutEventLog(txHash string, event *eventdatatypes.ChainlinkCallbackTimeoutEvent) string {
	return fmt.Sprintf("ChainlinkCallbackTimeoutEvent Transformed:\n"+
		"OrderID: %s\n"+
		"Order: %v\n",
		event.OrderID.String(),
		event.Order,
	)
}

// Function to format CouldNotCloseTradeEvent as a string
func GenerateCouldNotCloseTradeEventLog(txHash string, event *eventdatatypes.CouldNotCloseTradeEvent) string {
	return fmt.Sprintf("CouldNotCloseTradeEvent Transformed:\n"+
		"Trader: %s\n"+
		"PairIndex: %s\n"+
		"Index: %s\n",
		event.Trader,
		event.PairIndex.String(),
		event.Index.String(),
	)
}
