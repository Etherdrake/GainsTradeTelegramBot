package swapbuilders

import (
	"HootCore/pairmaps"
	"fmt"
)

func BuildSwapPairString(chain uint64) string {
	header := "📌 **All Crypto Pairs**\n\nSelect Asset To Trade (24hr)\n\n"
	var pairsList string

	pageArray := pairmaps.StarterSwapMap // -> THIS IS NOT UPDATED in COMMIT 100644
	switch chain {
	case 3:
		pairsList += fmt.Sprintf("/%s: %s \n", "ETH", "0x2170ed0880ac9a755fd29b2688956bd959f933f8")

	case 0:
		pairsList += fmt.Sprintf("/%s: %s \n", "MATIC", "0x0000000000000000000000000000000000001010")

	case 1:
		pairsList += fmt.Sprintf("/%s: %s \n", "WETH", "0x82af49447d8a07e3bd95bd0d56f35241523fbab1")

	default:
		return ""
	}

	for ticker, pair := range pageArray {
		switch chain {
		case 3:
			pairsList += fmt.Sprintf("/%s: %s \n", ticker, pair.EthereumAddress)
			continue
		case 0:
			pairsList += fmt.Sprintf("/%s: %s \n", ticker, pair.PolygonAddress)
			continue
		case 1:
			pairsList += fmt.Sprintf("/%s: %s \n", ticker, pair.ArbitrumAddress)
			continue
		default:
			return ""
		}
	}
	return header + pairsList
}
