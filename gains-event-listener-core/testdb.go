package main

import (
	"GainsListenerV4/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

var ctx = context.Background()

func TestDB(mongoHooter, mongoTrades *mongo.Database) {
	usersCollection := mongoHooter.Collection("users")
	guid, idErr := database.GetIdWithPublicKey(ctx, usersCollection, strings.ToLower("0xbc9b019aa5885ba50f3770b6fd7e5da981007068"))
	if idErr != nil {
		log.Println("Could not retrieve PublicKey: ", idErr)
	}
	xpAddErr := database.IncrementUserXP(ctx, usersCollection, int64(guid), 69)
	if xpAddErr != nil {
		log.Println("Failed to increment XP: ", xpAddErr)
	} else {
		log.Println("Successfully incremented XP: ", xpAddErr)
	}
}
