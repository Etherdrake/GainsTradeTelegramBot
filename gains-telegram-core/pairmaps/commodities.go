package pairmaps

var CommodityToIndex = map[string]int{
	"XAU/USD": 90,
	"XAG/USD": 91,
	"WTI/USD": 187,
	"XPT/USD": 188,
	"XPD/USD": 189,
	// Add more commodities if needed.
}

var IndexToCommodity = map[int]string{
	90:  "XAU/USD",
	91:  "XAG/USD",
	187: "WTI/USD",
	188: "XPT/USD",
	189: "XPD/USD",
	// Add more commodities if needed.
}
