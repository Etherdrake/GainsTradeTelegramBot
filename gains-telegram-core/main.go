package main

import (
	"HootCore/concurrentmaps"
	"HootCore/rdbhoot"
	redislocker "HootCore/rdbhoot/rdblocker"
	"context"
	"github.com/Etherdrake/telegram-bot-api/v7"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("We're so back baby!")

	// Initialize the Redis Client for the pricestream
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	defer rdbClient.Close()

	rdbHoot := &rdbhoot.HootRedisClient{Client: rdbClient}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Println("Error creating MongoDB client: ", err)
		return
	}
	ctxMongo, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctxMongo)
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Disconnect(ctxMongo)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_REAGAN_BETA"))
	if err != nil {
		log.Println(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Existing logic for bot updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updatesChan := bot.GetUpdatesChan(u)

	boardCache := concurrentmaps.New()

	errWipe := redislocker.WipeCacheSelectively(ctx, rdbHoot, "active_users")
	if errWipe != nil {
		log.Println("Fatal encountered during RedisLocker Flush: ", errWipe)
		return
	}

	errStartWipe := redislocker.WipeCacheSelectively(ctx, rdbHoot, "start_users")
	if errStartWipe != nil {
		log.Println("Fatal encountered during RedisLocker Flush: ", errStartWipe)
		return
	}

	go func() {
		for update := range updatesChan {
			// Check if we are sending this update ourselves
			if !update.SentFrom().IsBot {
				go func() {
					userID := update.SentFrom().ID
					exists, errCheck := redislocker.CheckUserID(ctx, rdbHoot, userID)
					if errCheck != nil {
						log.Println("Error checking UserID in RdbLocker: ", userID)
						return
					}
					if exists {
						log.Println(""+
							"RdbLocker User exists: ", userID)
						return
					}
					// Add userID to Redis
					errAdd := redislocker.AddUserID(ctx, rdbHoot, userID)
					if errAdd != nil {
						log.Printf("Error add UserID %d to Redis: %v", userID, errAdd)
					}
					// Distribute the update to worker goroutines
					go func() {
						handleUpdates(bot, client, update, boardCache, rdbHoot)
						// Remove UserID from Redis after handleUpdates is done
						if errRemove := redislocker.RemoveUserID(ctx, rdbHoot, userID); errRemove != nil {
							log.Printf("Error removing UserID %d from Redis: %v", userID, errRemove)
						}
					}()
				}()
			}
		}
	}()

	// Block the main function from exiting
	select {}
}
