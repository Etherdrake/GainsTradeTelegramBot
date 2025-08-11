package stringbuilders

import (
	"HootCore/altseasoncore"
	"HootCore/api/gnsfees"
	"HootCore/formatters"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

func BuildOrderOrTradeStringV2(
	client *mongo.Client,
	rdbHoot *rdbhoot.HootRedisClient,
	guid int64,
	chain uint64,
	selectedIsTrade bool,
	openTradesOrOrdersArr []altseasoncore.TradeItemCore) string {

	// Retrieve activeOrderOrTrade
	var currentPrice float64
	var closeFee int
	var pairSpreadStr string
	var liquidation float64
	var collToken string
	var err error

	var chainStr string
	// Check on which mode / chain we are:
	switch chain {
	case 137:
		chainStr = "Polygon"
	case 42161:
		chainStr = "Arbitrum"
	case 421614:
		chainStr = "Sepolia"
	}

	// isTradeOrOrder is the uint that
	doesTradeOrOrderExist, _, selectedTradeOrOrder, err := rdbHoot.GetSelectedTradeOrOrderOrNone(guid) // 0 is nothing set, 1 trade, 2 is order
	if err != nil {
		if !doesTradeOrOrderExist {
			log.Println("GetSelectedTradeOrOrderOrNone failed in BuildOrderOrTradeV2")
		}
		return ""
	}

	selectedTradeId, err := rdbHoot.GetSelectedTradeID(guid)
	if err != nil {
		log.Println("Error retrieving GetSelectedTradeID in buildorderortradev2")
	}

	var selectedTradeOrOrderIndexInArr = -1 // Initialize to -1 outside the loop

	// Find the index by sifting through the ordered array that was given
	for i := 0; i < len(openTradesOrOrdersArr); i++ {
		if selectedTradeId == openTradesOrOrdersArr[i].ID {
			selectedTradeOrOrderIndexInArr = i
			break
		}
	}

	if selectedTradeOrOrderIndexInArr == -1 {
		selectedTradeOrOrderIndexInArr = 0
	}

	var isTrade bool
	if selectedIsTrade {
		isTrade = true
	} else {
		isTrade = false
	}

	var activeTradeOrOrderString string
	if isTrade {
		// We increment by one because otherwise stuff is not human-readable
		activeTradeOrOrderString, err = BuildSingleTradeStringGNS(rdbHoot, selectedTradeOrOrder, selectedTradeOrOrderIndexInArr+1, isTrade)
		if err != nil {
			log.Println("Error in BuildSingleTradeStringGNS inside buildorderortradev2: ", err)
		}
	} else {
		// We increment by one because otherwise stuff is not human-readable
		activeTradeOrOrderString, err = BuildSingleOrderStringGNS(rdbHoot, selectedTradeOrOrder, selectedTradeOrOrderIndexInArr+1, selectedTradeOrOrder.ID)
		if err != nil {
			log.Println("Error in BuildSingleTradeStringGNS inside buildorderortradev2: ", err)
		}
	}

	// Retrieve the price from Redis
	currentPrice, err = rdbHoot.GetPrice(int(selectedTradeOrOrder.Trade.PairIndex)) // Assuming you have the GetPrice function implemented to fetch price from Redis
	if err != nil {
		fmt.Println("Error fetching price from Redis in BuildActiveTradeStringV2:", err)
	}

	// TODO: Rework the fees
	allFees := gnsfees.FeesResult{
		Borrowing: struct {
			FeePerBlock float64 `json:"feePerBlock"`
			FeePerHour  float64 `json:"feePerHour"`
		}{
			FeePerBlock: 0.0,
			FeePerHour:  0.0,
		},
		Fees:      0.0,
		AvgSpread: 0.0,
	}

	pairSpreadStr = fmt.Sprintf("%.2f%%", allFees.AvgSpread)

	dSelectedTradeOrOrder := selectedTradeOrOrder.Trade.IntoDisplay()

	//collAmtF64, err := strconv.ParseFloat(dSelectedTradeOrOrder.CollateralAmount, 64)

	openPriceF64, err := strconv.ParseFloat(dSelectedTradeOrOrder.OpenPrice, 64)
	leverageF32, err := strconv.ParseFloat(dSelectedTradeOrOrder.Leverage, 32)

	//totalPosSize := collAmtF64 * leverageF32

	liquidation = utils.CalculateLiquidationPrice(openPriceF64, float64(leverageF32), selectedTradeOrOrder.Trade.Long)
	collToken = strconv.Itoa(int(selectedTradeOrOrder.Trade.CollateralIndex))

	entryPriceFormattedLegacy := formatters.FormatPrice(selectedTradeOrOrder.Trade.OpenPrice)
	if err != nil {
		log.Println("Error in BuildSingleTradeStringGNS inside formatters.FormatFinancialUS(dSelectedTradeOrOrder.OpenPrice): ", err)
	}
	//openUK, err := formatters.FormatFinancialUS(openPriceF64)
	currentPriceFormattedLegacy := formatters.FormatPrice(currentPrice)
	//positionSizeUS, err := formatters.FormatFinancialUS(totalPosSize)
	takeProfitUS, err := formatters.FormatFinancialUS(dSelectedTradeOrOrder.Tp)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}
	stopLossUS, err := formatters.FormatFinancialUS(dSelectedTradeOrOrder.Sl)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}
	liquidationUS, err := formatters.FormatFinancialUS(liquidation)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}

	var longShortString string
	if selectedTradeOrOrder.Trade.Long {
		longShortString = "LONG"
	} else {
		longShortString = "SHORT"
	}

	switch selectedTradeOrOrder.Trade.CollateralIndex {
	case 1:
		collToken = "DAI"
	case 2:
		collToken = "WETH"
	case 3:
		collToken = "USDC"
	default:
		// Handle other cases or add a default value if needed
		collToken = "UNKNOWN"
	}

	msg := fmt.Sprintf(`
%s

🦋 Entry Price: *$%s*
💲 Current Price: *$%s*
💰 Total collateral: *%s DAI*
⚖️ Leverage: *%sx %s*
🤸 Spread: %s | Token: *%s*
TP: *%s* | SL: *%s*
LIQ: *%s* | 🎯 Gains Network
⛓️ *%s* | ⛽ Fees: %d
-----------------------------------------`,
		activeTradeOrOrderString,
		entryPriceFormattedLegacy,
		currentPriceFormattedLegacy,
		dSelectedTradeOrOrder.CollateralAmount, // Notice this change
		dSelectedTradeOrOrder.Leverage,
		longShortString,
		pairSpreadStr,
		collToken,
		takeProfitUS,
		stopLossUS,
		liquidationUS,
		chainStr,
		closeFee,
	)
	return msg
}
