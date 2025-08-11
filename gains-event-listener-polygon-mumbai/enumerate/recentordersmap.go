package enumerate

import "GainsListenerMumbai/api/orders"

func FindMostRecent(orderArr []orders.OpenOrdersHuman) map[string]orders.OpenOrdersHuman {
	latestOrders := make(map[string]orders.OpenOrdersHuman)

	// Iterate over orders and update the latest order for each trader
	for _, order := range orderArr {
		if current, ok := latestOrders[order.Trader]; !ok || order.Block > current.Block {
			latestOrders[order.Trader] = order
		}
	}

	return latestOrders
}

func FindMostRecentSingle(orderArr []orders.OpenOrdersHuman) *orders.OpenOrdersHuman {
	var mostRecentOrder *orders.OpenOrdersHuman

	// Iterate over orders and update the most recent order
	for _, order := range orderArr {
		if mostRecentOrder == nil || order.Block > mostRecentOrder.Block {
			mostRecentOrder = &order
		}
	}

	return mostRecentOrder
}
