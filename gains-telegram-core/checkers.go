package main

import (
	"HootCore/api/native"
	"HootCore/database"
	"HootCore/mongolisten"
	"HootCore/rdbhoot"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)

func CheckIfUserHasOpenTradesOrOrdersGNS(client *mongo.Client, guid int64, chain uint64) (bool, error) {
	publicKey, err := database.GetPublicKey(client, guid)
	if err != nil {
		log.Println("PublicKey Error: ", err)
	}

	hasTrades, err := mongolisten.HasOpenTrades(client, publicKey, chain)
	if err != nil {
		return false, err
	}

	hasOrders, err := mongolisten.HasOpenOrders(client, publicKey, chain)
	if err != nil {
		return false, err
	}
	if hasTrades || hasOrders {
		return true, nil
	} else {
		return false, nil
	}
}

// DiamondApprovalCheck performs a check on whether the active coll has been approved to the diamond, if this is not the case it does an approval for the user
func DiamondApprovalCheck(client *mongo.Client, rdbHoot *rdbhoot.HootRedisClient, guid int64) (bool, string) {
	pubkey, err := database.GetPublicKey(client, guid)
	if err != nil {
		fmt.Println("Error is: ", err)
	}

	trade, exists := rdbHoot.GetCoreCache(guid)
	if !exists {
		log.Println("Guid does not exist: ", guid)
	}

	payload := native.GetGasBalancePayload{
		PublicKey:  pubkey,
		Collateral: trade.Collateral,
		Chain:      trade.Chain,
	}

	// First we check if the wallet even has a gastoken
	gasBalance, _, err := native.GetGasBalance(payload)
	if err != nil {
		log.Println("Error retrieving gasBalance inside checkers.go: ", err)
		return false, ""
	}
	gasBalanceFloat, err := strconv.ParseFloat(gasBalance, 64)
	if err != nil {
		return false, ""
	}
	switch trade.Chain {
	case 137:
		if gasBalanceFloat > 0.5 {
			_, approvedAmountNumber, allowanceErr := native.GetAllowanceGains(client, *trade)
			if allowanceErr != nil {
				log.Println("Error GetAllowance Gains: ", err)
				return false, "Error"
			}
			if !(approvedAmountNumber > 0) {
				txID, err := native.ApproveCollateralToDiamond(
					client,
					guid,
					*trade,
				)
				if err != nil {
					log.Println(err)
				}
				return true, txID
			} else {
				return false, ""
			}
		} else {
			return false, ""
		}
	case 42161:
		if gasBalanceFloat > 0.0025 {
			_, approvedAmountNumber, allowanceErr := native.GetAllowanceGains(client, *trade)
			if allowanceErr != nil {
				log.Println("Error GetAllowance Gains: ", err)
				return false, "Error"
			}
			if !(approvedAmountNumber > 0) {
				txID, err := native.ApproveCollateralToDiamond(
					client,
					guid,
					*trade,
				)
				if err != nil {
					log.Println(err)
				}
				return true, txID
			} else {
				return false, ""
			}
		} else {
			return false, ""
		}
	case 421614:
		if gasBalanceFloat > 0.0025 {
			_, approvedAmountNumber, err := native.GetAllowanceGains(client, *trade)
			if err != nil {
				log.Println(err)
				return false, "Error"
			}
			if !(approvedAmountNumber > 1) {
				txID, err := native.ApproveCollateralToDiamond(
					client,
					guid,
					*trade,
				)
				if err != nil {
					log.Println("Error: ", err)
				}
				return true, txID
			} else {
				return false, ""
			}
		} else {
			return false, ""
		}
	}
	return false, ""
}
