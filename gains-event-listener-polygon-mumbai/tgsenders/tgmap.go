package tgsenders

var EventnameToTelegramChannel = map[string]int{
	"MarketExecuted": 44,
	"LimitExecuted":  48,

	//"OpenLimitCanceled":      327,
	"MarketOpenCanceled":     331,
	"MarketCloseCanceled":    331,
	"NftOrderCanceled":       331,
	"PairMaxLeverageUpdated": 533,
	//"NftOrderInitiated":  319,
	//"OpenLimitPlaced":        307,
	//"OpenLimitUpdated":       311,
	//"TpUpdated":          315,
	//"SlUpdated":          315,
	//"CouldNotCloseTrade": 323,

	// Trading
	"Done":                     1516,
	"Paused":                   1520,
	"NumberUpdated":            1524,
	"BypassTriggerLinkUpdated": 1528,
	"MarketOrderInitiated":     1532,
	"OpenLimitPlaced":          1536,
	"OpenLimitUpdated":         1540,
	"OpenLimitCanceled":        1544,
	"TpUpdated":                1548,
	"SlUpdated":                1552,
	"NftOrderInitiated":        1556,
	"ChainlinkCallbackTimeout": 1560,
	"CouldNotCloseTrade":       1564,

	// Aggregated Channels
	"HootCallbackErrors": 498,

	// Amalgamated and Matches
	"CacheOverview":           4853,
	"AmalgamatedConsolidated": 4860,
	"AmalgamatedCache":        5922,
}
