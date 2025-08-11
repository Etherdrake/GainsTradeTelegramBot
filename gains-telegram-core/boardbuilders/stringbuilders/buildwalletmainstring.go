package stringbuilders

import (
	"HootCore/api/native"
	"HootCore/database"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func BuildWalletMainString(client *mongo.Client, guid int64) (string, error) {
	keys, err := database.GetKeys(client, guid)
	if err != nil {
		return "", err
	}

	// TODO: Make sure the private key is encrypted with a user generated unique identifier like a PIN or fingerprint
	publicKey := "`" + keys.PublicKey + "`"
	privateKey := "`" + keys.PrivateKey + "`"

	polygonPayload := native.GetGasBalancePayload{
		PublicKey: keys.PublicKey,
		Chain:     137,
	}

	arbitrumPayload := native.GetGasBalancePayload{
		PublicKey: keys.PublicKey,
		Chain:     42161,
	}

	// TODO: Ethereum is not yet added
	//ethereumPayload := native.GetGasBalancePayload{
	//	PublicKey: keys.PublicKey,
	//	Chain:     1,
	//}

	sepoliaPayload := native.GetGasBalancePayload{
		PublicKey: keys.PublicKey,
		Chain:     421614,
	}

	// TODO: This is wrong
	polygonBal, _, err := native.GetGasBalance(polygonPayload)
	if err != nil {
		log.Println("Error getting balance gas token for chain: ", polygonPayload.Chain)

		return "", err
	}

	arbitrumBal, _, err := native.GetGasBalance(arbitrumPayload)
	if err != nil {
		log.Println("Error getting balance gas token for chain: ", arbitrumPayload.Chain)
		return "", err
	}

	//ethereumBal, err := native.GetGasBalance(ethereumPayload)
	//if err != nil {
	//	log.Println("Error getting balance gas token for chain: ", ethereumPayload.Chain)
	//	return "", err
	//}

	sepoliaBal, _, err := native.GetGasBalance(sepoliaPayload)
	if err != nil {
		log.Println("Error getting balance gas token for chain: ", sepoliaPayload.Chain)
		return "", err
	}

	// We fetch the gastoken three times here
	balanceMatic := polygonBal + " MATIC"
	balanceArb := arbitrumBal + " ARB"
	//z
	balanceSep := sepoliaBal + "TESTNET ETH"

	return fmt.Sprintf("Publickey: %s\n\nPrivatekey: %s\n\nBalance MATIC: %s\nBalance ARB: %s\nBalance ETH: TODO \nBalance Testnet: %s",
		publicKey, privateKey, balanceMatic, balanceArb, balanceSep), nil
}

func BuildSimpleWalletMainString(client *mongo.Client, guid int64) (string, error) {
	keys, err := database.GetKeys(client, guid)
	if err != nil {
		return "", err
	}

	// TODO: Make sure the private key is encrypted with a user generated unique identifier like a PIN or fingerprint
	publicKey := "`" + keys.PublicKey + "`"
	privateKey := "`" + keys.PrivateKey + "`"

	return fmt.Sprintf("Publickey: %s\n\nPrivatekey: %s\n\n", publicKey, privateKey), nil
}
