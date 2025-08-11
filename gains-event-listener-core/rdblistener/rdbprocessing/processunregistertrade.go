package rdbprocessing

import (
	"GainsListenerV4/rdblistener"
	"GainsListenerV4/types"
	"context"
	"fmt"
	"log"
	"strings"
)

func ProcessUnregisterTrade(rdbHoot rdblistener.HootRedisClient, ctx context.Context, unregisterTrade types.UnregisteredTrade) error {
	// Convert the user to lowercase and create a unique identifier
	trader := strings.ToLower(unregisterTrade.User)
	identifier := trader + unregisterTrade.Index

	//usersCollection := mongoHooter.Collection("users")

	// Check if the trade exists
	exists, err := rdbHoot.CheckIfTradeExists(ctx, trader, identifier)
	if err != nil {
		log.Println(err)
	}

	switch exists {
	case true:
		// Delete the trade by its identifier from the primary collection
		delErr := rdbHoot.DeleteTradeByID(ctx, trader, identifier)
		if delErr != nil {
			return delErr
		}
		log.Println("unregister trade successfully: ", identifier)
	case false:
		// Check if the order exists
		orderExists, rdbErr := rdbHoot.CheckIfOrderExists(ctx, trader, identifier)
		if err != nil {
			log.Println(rdbErr)
		}

		switch orderExists {
		case true:
			delErr := rdbHoot.DeleteOrderByID(ctx, trader, identifier)
			if delErr != nil {
				return delErr
			}
		}

		log.Println("unregister order successful: ", identifier)
	}
	return fmt.Errorf("unregister trade unsuccessfully.\nUser: %s\nIdentifier: %s", unregisterTrade.User, identifier)
}
