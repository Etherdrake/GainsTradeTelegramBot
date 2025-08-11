package regexes

import "regexp"

var (
	TradePattern    = regexp.MustCompile(`/t(\d+)`)
	OrderPattern    = regexp.MustCompile(`/o(\d+)`)
	HistoryPattern  = regexp.MustCompile(`/h(\d+)`)
	LeveragePattern = regexp.MustCompile(`\d+`)

	// PairMenu
	CatcherPattern = regexp.MustCompile(`/[A-Z]+x[A-Z]+`)
)
