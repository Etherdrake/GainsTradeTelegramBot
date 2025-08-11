package swapbuilders

import (
	"HootCore/api/native"
	"HootCore/database"
	"HootCore/rdbhoot"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func BuildSwapStringV1(client *mongo.Client, rdbHoot *rdbhoot.HootRedisClient, guid int64) (string, error) {
	pubKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		return "", err
	}

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("User does not exist", guid)
	}

	getDaiPoly := native.BalanceOfPayload{
		PublicKey:  pubKey,
		Collateral: 1,
		Chain:      trade.Chain,
	}

	daiBalance, err := native.GetBalanceOf(getDaiPoly)
	if err != nil {
		fmt.Println("GetBalance Error: ", err)
	}

	getUsdcPoly := native.BalanceOfPayload{
		PublicKey:  pubKey,
		Collateral: 2,
		Chain:      trade.Chain,
	}

	usdcBalance, err := native.GetBalanceOf(getUsdcPoly)
	if err != nil {
		fmt.Println("GetBalance Error: ", err)
	}

	getWethPoly := native.BalanceOfPayload{
		PublicKey:  pubKey,
		Collateral: 3,
		Chain:      trade.Chain,
	}

	wethBalance, err := native.GetBalanceOf(getWethPoly)
	if err != nil {
		fmt.Println("GetBalance Error: ", err)
	}

	payloadPoly := native.GetGasBalancePayload{
		PublicKey:  pubKey,
		Collateral: 0,
		Chain:      137,
	}

	gasBalancePolygon, _, err := native.GetGasBalance(payloadPoly)
	if err != nil {
		log.Println("Get GasBalance Error: ", err)
	}

	payloadArbitrum := native.GetGasBalancePayload{
		PublicKey: pubKey,
		Chain:     42161,
	}

	gasBalanceArbitrum, _, err := native.GetGasBalance(payloadArbitrum)
	if err != nil {
		log.Println("Get GasBalance Error: ", err)
	}

	publicKey := "`" + pubKey + "`"

	// You would need to replace the 0.00 balances with actual values.
	// These could be fetched similarly to how the keys are fetched.
	//balanceMatic := "0.00 MATIC"
	//balanceArb := "0.00 ARB"
	//balanceEth := "0.00 ETH"

	return fmt.Sprintf(
		"Your Wallet: \n\n%s\n\n"+
			"*Arbitrum Holdings:* \n\n"+
			"Balance ETH: %s\n"+
			"Balance DAI: %s\n"+
			"Balance WETH: %s\n"+
			"Balance USDC: %s\n\n"+
			"*Polygon Holdings:*\n\n"+
			"Balance MATIC: %s\n"+
			"Balance DAI: %s\n"+
			"Balance WETH: %s\n"+
			"Balance USDC: %s\n",
		publicKey,
		gasBalanceArbitrum,
		daiBalance,
		wethBalance,
		usdcBalance,
		gasBalancePolygon,
		daiBalance,
		wethBalance,
		usdcBalance,
	), nil
}
