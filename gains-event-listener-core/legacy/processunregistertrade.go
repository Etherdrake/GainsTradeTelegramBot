package legacy

import (
	"GainsListenerV4/database"
	"GainsListenerV4/helpers"
	"GainsListenerV4/types"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

func ProcessUnregisterTrade(mongoHooter, mongoTrades *mongo.Database, unregisterTrade types.UnregisteredTrade) error {
	// Convert the user to lowercase and create a unique identifier
	trader := strings.ToLower(unregisterTrade.User)
	identifier := trader + unregisterTrade.Index

	//usersCollection := mongoHooter.Collection("users")

	// Get the primary MongoDB collection
	primaryCollection, getErr := helpers.GetMongoCollection(mongoTrades, true, false)
	if getErr != nil {
		return errors.New("failed to get primary MongoDB collection")
	}

	// Check if the trade exists in the primary collection
	exists, checkErr := database.TradeExistsByID(ctx, primaryCollection, identifier)
	if checkErr != nil {
		return checkErr
	}

	switch exists {
	case true:
		// Delete the trade by its identifier from the primary collection
		delErr := database.DeleteTradeByID(ctx, primaryCollection, identifier)
		if delErr != nil {
			return delErr
		}
		log.Println("unregister trade successfully: ", identifier)
	case false:
		// Get the secondary MongoDB collection
		secondaryCollection, getErr := helpers.GetMongoCollection(mongoTrades, false, true)
		if getErr != nil {
			return errors.New("failed to get secondary MongoDB collection")
		}

		// Try to delete the trade by its identifier from the secondary collection
		delErr := database.DeleteTradeByID(ctx, secondaryCollection, identifier)
		if delErr != nil {
			return delErr
		}
		log.Println("unregister order successfully: ", identifier)
	}
	return fmt.Errorf("unregister trade unsuccessfully.\nUser: %s\nIdentifier: %s", unregisterTrade.User, identifier)
}
