package mongolisten

import (
	"HootCore/altseasoncore"
	"fmt"
	"sort"
)

func SortOpenOrdersByBlock(orders []altseasoncore.TradeItemMongo) ([]altseasoncore.TradeItemMongo, error) {
	if orders == nil {
		return nil, fmt.Errorf("orders slice is nil")
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].TradeInfo.CreatedBlock > orders[j].TradeInfo.CreatedBlock
	})

	return orders, nil
}
