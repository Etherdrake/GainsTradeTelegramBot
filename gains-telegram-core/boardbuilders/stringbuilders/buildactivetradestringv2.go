package stringbuilders

import (
	"HootCore/altseasoncore"
	"HootCore/altseasontools"
	"HootCore/api/native"
	"HootCore/database"
	"HootCore/formatters"
	"HootCore/mongolisten"
	"HootCore/rdbhoot"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func BuildActiveGainsTradeStringV2(
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	guid int64, chain uint64) string {

	publicKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		log.Println("Error getting pubkey")
	}

	_, openTrade, openOrder, err := mongolisten.HasOpenTradeOrOrder(client, guid, chain)

	// 0 is nothing set, 1 trade, 2 is order for selectedTradeOrOrder
	doesTradeOrOrderExist, isTradeOrOrder, selectedTradeOrOrder, err := rdbHoot.GetSelectedTradeOrOrderOrNone(guid)
	if err != nil {
		if !doesTradeOrOrderExist {
			log.Println("GetSelectedTradeOrOrderOrNone failed in BuildOrderOrTradeV2")
		}
		return ""
	}

	if openTrade || openOrder {
		var selectedIsTrade bool
		switch isTradeOrOrder {
		// There is no trade
		case 0:
			log.Println("IsTradeOrOrder is 0")
		// It's a trade
		case 1:
			selectedIsTrade = true
		// It's an order
		case 2:
			selectedIsTrade = false
		}

		// TODO: What the fuck do you even have here you motherfucker
		openTrades, err := mongolisten.GetAllOpenTrades(client, publicKey, chain)
		if err != nil {
			log.Println("GetAllTraders error: ", err)
		}

		openOrders, err := mongolisten.GetOpenOrders(client, publicKey, chain)
		if err != nil {
			log.Println("GetOpenOrders error: ", err)
		}

		var openTradesMsg string
		var openOrdersMsg string
		var totalTrades int
		var totalOrders int
		var activeTradeOrOrder string

		var sortedTrades []altseasoncore.TradeItemCore

		// We generate the string here, and we skip the selectedTrade if it's selected
		if openTrade {
			sortedTrades = altseasontools.SortTrades(openTrades)
			for z := len(sortedTrades) - 1; z >= 0; z-- {
				// TODO: Active Trade -> This should not compare CreatedBlock but the new UID
				if selectedIsTrade {
					if selectedTradeOrOrder.ID == sortedTrades[z].ID {
						totalTrades += 1
						continue
					}
				}
				openTradeStr, err := BuildSingleTradeStringGNS(rdbHoot, sortedTrades[z].IntoInternal(), z+1, true)
				if err != nil {
					log.Println("Error generating position string in BuildActiveTradeStringGNS")
				}
				if z != len(sortedTrades) {
					openTradesMsg += openTradeStr + "\n"
				} else {
					openTradesMsg += openTradeStr
				}
				totalTrades += 1
			}
			if selectedIsTrade {
				// We retrieve the activeTradeOrOrder
				activeTradeOrOrder = BuildOrderOrTradeStringV2(
					client,
					rdbHoot,
					guid,
					chain,
					selectedIsTrade,
					openTrades)
			}

		} else {
			openTradesMsg = "You don't have any open trades."
		}

		sortedOrders := altseasontools.SortTrades(openOrders)

		// We generate the string here, and we skip the selectedOrder if it's selected
		if openOrder {
			for x := len(sortedTrades) - 1; x >= 0; x-- {
				if !selectedIsTrade {
					// We increment one extra if the selected order is found
					if selectedTradeOrOrder.ID == sortedOrders[x].ID {
						totalOrders += 1
						continue
					}
				}
				openOrderStr, err := BuildSingleOrderStringGNS(rdbHoot, sortedOrders[x].IntoInternal(), x+1, sortedOrders[x].ID)
				if err != nil {
					log.Println("Error getting generating order string in BuildPositionsAndOrderString")
				}
				openOrdersMsg += openOrderStr + "\n"
				totalOrders += 1
			}
			if !selectedIsTrade {
				// We retrieve the activeTradeOrOrder
				activeTradeOrOrder = BuildOrderOrTradeStringV2(
					client,
					rdbHoot,
					guid,
					chain,
					selectedIsTrade,
					openOrders)
			}
		} else {
			openOrdersMsg = "You don't have any open orders."
		}

		totalActive := totalTrades

		trade, exists := rdbHoot.GetCoreCache(guid)
		if !exists {
			log.Println("Trade not found in cache for guid: ", guid)
		}

		getDaiPOST := native.BalanceOfPayload{
			PublicKey:  publicKey,
			Collateral: trade.Collateral,
			Chain:      trade.Chain,
		}

		bal, err := native.GetBalanceOf(getDaiPOST)
		if err != nil {
			fmt.Println("GetBalance Error: ", err)
		}

		balanceFloatUS, fmtErr := formatters.FormatFinancialUS(bal)
		if fmtErr != nil {
			log.Println("FormatFinancialUS Error in BuildPerpMainStringV2: ", fmtErr)
		}

		xp, rdbErr := rdbHoot.GetExperience(ctx, guid)
		if rdbErr != nil {
			log.Println("GetXP Error in BuildPerpMainStringV2: ", err)
		}

		topString := fmt.Sprintf("-----------------------------------------\n💸 Balance: *%s* /add\n🎟 Hoot Coins: *%s*\n⭐️ Total Active Trades: *%d*\n-----------------------------------------", balanceFloatUS, xp, totalActive)

		var msg string
		// We have no trades but an order
		if len(openTrades) == 0 && len(openOrders) > 0 {
			msg = fmt.Sprintf(`
%s
%s

%s`,
				topString,
				activeTradeOrOrder,
				openOrdersMsg,
			)
		} else {
			// We have trades and an order
			if len(openTrades) > 0 && len(openOrders) > 0 {
				msg = fmt.Sprintf(`
%s
%s

%s
%s`,
					topString,
					activeTradeOrOrder,
					openTradesMsg,
					openOrdersMsg,
				)
			} else {
				// We have trades and no order
				if len(openTrades) > 0 && len(openOrders) == 0 {
					msg = fmt.Sprintf(`
%s
%s

%s`,
						topString,
						activeTradeOrOrder,
						openTradesMsg,
					)
				}
			}
		}
		return msg
	} else {
		return ""
	}
}
