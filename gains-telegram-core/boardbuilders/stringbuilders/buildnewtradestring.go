package stringbuilders

import (
	"HootCore/api/native"
	"HootCore/database"
	"HootCore/formatters"
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var ctx = context.Background()

func BuildNewTradeString(client *mongo.Client, rdbHoot *rdbhoot.HootRedisClient, guid int64, pairIndex int64) string {
	publicKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		return "GetPublicKey Error in BuildPerpMainStringV2"
	}

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User not found in cache", guid)
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

	//balFloat, err := strconv.ParseFloat(bal, 64)
	//if err != nil {
	//	fmt.Println("GetBalance Error: ", err)
	//}

	//var testTokens string
	//if balFloat <= 900 {
	//	testTokens = "/getdai"
	//} else {
	//	testTokens = ""
	//}

	payload := native.GetGasBalancePayload{
		PublicKey:  publicKey,
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	gasBalance, _, err := native.GetGasBalance(payload)
	if err != nil {
		log.Println("Get GasBalance Error: ", err)
	}

	trimGasBalance := utils.TrimToLength(gasBalance, 7)

	var gasToken string
	switch trade.Chain {
	case 137:
		gasToken = "MATIC"
	case 42161:
		gasToken = "ETH"
	case 421614:
		gasToken = "ETH"
		gasToken = "ETH"
	}
	//balanceFloat, _ := strconv.ParseFloat(bal, 64)

	balanceFloatUS, fmtErr := formatters.FormatFinancialUS(bal)
	if fmtErr != nil {
		log.Println("FormatFinancialUS Error in BuildPerpMainStringV2: ", fmtErr)
	}

	// Retrieve the currentPrice from Redis
	currentPrice, err := rdbHoot.GetPrice(int(trade.PairIndex)) // Assuming you have the GetPrice function implemented to fetch currentPrice from Redis
	if err != nil {
		log.Println("Error fetching currentPrice from Redis in Stringbuilder:", err)
	}
	var LongShortString string
	if trade.Buy {
		LongShortString = "Long"
	} else {
		LongShortString = "Short"
	}

	var orderType string
	switch trade.OrderType {
	case 0:
		orderType = "Market"
	case 1:
		orderType = "Limit"
	case 2:
		orderType = "Stop"
	default:
		log.Println("Unrecognized order type for user: ", trade.OrderType)
	}

	var chainStr string
	var _ string
	var realPractice string
	// Check on which mode / chain we are:
	switch trade.Chain {
	case 137:
		chainStr = "Polygon"
		_ = "/arbitrum"
		realPractice = "🟢 Live Mode /practice"
	case 42161:
		chainStr = "Arbitrum"
		_ = "/polygon"
		realPractice = "🟢 Live Mode /practice"
	case 421614:
		chainStr = "Practice"
		_ = "/arbitrum"
		realPractice = "🧪 Practice Mode Testnet"
	}

	directOrderString := orderType + " " + LongShortString

	var activeCollateral string
	switch trade.Collateral {
	//case 0:
	//	activeCollateral = "Active Collateral: *native* /dai /weth | /usdc"
	case 1:
		activeCollateral = "Active Collateral: *DAI*"
		//case 2:
		//	activeCollateral = "Active Collateral: *WETH* /dai | /usdc | /native"
		//case 3:
		//	activeCollateral = "Active Collateral: *USDC* /dai | /weth | /native"
	}

	positionSize := trade.PositionSizeDai * float64(trade.Leverage)

	xp, err := rdbHoot.GetExperience(ctx, guid)

	currentPriceFormatted := formatters.FormatPrice(currentPrice)
	positionSizeUS, err := formatters.FormatFinancialUS(trade.PositionSizeDai)
	takeProfitUS, err := formatters.FormatFinancialUS(trade.TP)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}
	stopLossUS, err := formatters.FormatFinancialUS(trade.SL)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}
	liquidationUS, err := formatters.FormatFinancialUS(trade.Liq)
	if err != nil {
		log.Println("Error taking financial US in buildorderortradev2", err)
	}

	//allFees, err := gnsfees.GetFeesAPI(chain, trade.Buy, trade.PositionSizeDai, trade.Leverage, trade.PairIndex, trade.Collateral, trade.PositionSizeDai)
	//if err != nil {
	//	log.Println("Error calling GetFeesAPI")
	//}
	// TODO: ☁️ Swap: /swap -> COMING SOON
	msg := fmt.Sprintf(`
-----------------------------------------
%s 
💸 Balance: *%s* /add
⛽ Gas Balance: *%s %s*
🍥 Sushi: *%s*
🪙 %s
-----------------------------------------
⭐️ *NEW TRADE*

%s: %s @ *$%s*
with *%s* collateral and *%.2fx* lever

💲 Current Price: *$%s*
💰 Total position size: *$%.2f*
🤸 Spread: *%.4f%%*
⛽ Fees:* $%d* | LIQ: *%s*
TP: *%s* SL: *%s* 
🔗 %s`,
		realPractice,
		balanceFloatUS,
		trimGasBalance,
		gasToken,
		xp,
		activeCollateral,
		directOrderString,
		utils.EscapeMarkdownV2(pairmaps.IndexToPair[int(pairIndex)]),
		currentPriceFormatted,
		positionSizeUS,
		trade.Leverage,
		currentPriceFormatted,
		positionSize, // Notice this change
		0.69,         // PLACEHOLDER
		420,          // PLACEHOLDER
		liquidationUS,
		takeProfitUS,
		stopLossUS,
		chainStr,
		//switchStr,
	)
	return msg
}
