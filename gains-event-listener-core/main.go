package main

import (
	"GainsListenerV4/database"
	"GainsListenerV4/legacy"
	"GainsListenerV4/rdblistener"
	"GainsListenerV4/rdblistener/rdbprocessing"
	"GainsListenerV4/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	WSS_ARBITRUM = "wss://backend-arbitrum.gains.trade"
	WSS_POLYGON  = "wss://backend-polygon.gains.trade"
	WSS_SEPOLIA  = "wss://backend-sepolia.gains.trade"
)

const (
	maxRetries    = 5
	retryInterval = 2 * time.Second
)

var isPoly bool
var isArb bool
var isSepolia bool

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize the Redis Client for the pricestream
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	defer func(rdbClient *redis.Client) {
		rdbErr := rdbClient.Close()
		if rdbErr != nil {
			log.Println(rdbErr)
		}
	}(rdbClient)

	rdbHoot := rdblistener.HootRedisClient{Client: rdbClient}

	// WebSocket server address
	var addr string
	chainPtr := os.Getenv("CHAIN")

	if chainPtr == "polygon" {
		addr = WSS_POLYGON
		isPoly = true
		isArb = false
		isSepolia = false
	} else if chainPtr == "arbitrum" {
		addr = WSS_ARBITRUM
		isPoly = false
		isArb = true
		isSepolia = false
	} else if chainPtr == "sepolia" {
		addr = WSS_SEPOLIA
		isPoly = false
		isArb = false
		isSepolia = true
	} else {
		panic("invalid chain option")
	}

	// Supply manually to docker for testing like in Farcaster
	MongoURI := os.Getenv("MONGO_URI")

	// Create WebSocket dialer
	dialer := websocket.DefaultDialer

	// Connect to the WebSocket server
	conn, _, err := dialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	defer conn.Close()

	// Connect to MongoDB
	var mongoTradeOrder *mongo.Database
	var mongoHooter *mongo.Database

	if isPoly {
		mongoTradeOrder, err = database.ConnectToMongoDB(MongoURI, "polygon")
		if err != nil {
			log.Fatal("MongoDB could not be initialized")
		}
	} else {
		if isArb {
			mongoTradeOrder, err = database.ConnectToMongoDB(MongoURI, "arbitrum")
			if err != nil {
				log.Fatal("MongoDB could not be initialized")
			}
		} else {
			if isSepolia {
				mongoTradeOrder, err = database.ConnectToMongoDB(MongoURI, "sepolia")
				if err != nil {
					log.Fatal("MongoDB could not be initialized")
				}
				mongoHooter, err = database.ConnectToMongoDB(MongoURI, "hooterdb")
				if err != nil {
					log.Fatal("MongoDB could not be initialized")
				}
			}
		}
	}

	//TestDB(mongoHooter, mongoTradeOrder)
	//time.Sleep(1 * time.Minute)

	// Create a channel to receive signals for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Main loop to read messages from the WebSocket server
	for {
		select {
		case <-interrupt:
			log.Println("Interrupt signal received, closing connection...")
			return
		default:
			// Read message from WebSocket server
			_, message, readErr := conn.ReadMessage()
			if readErr != nil {
				log.Println("Read:", readErr)
				closeErr := conn.Close()
				if closeErr != nil {
					return
				}
				// Close the existing connection

				time.Sleep(retryInterval)
				conn, _, err = dialer.Dial(addr, nil) // Create a new connection
				if err != nil {
					log.Fatal("Dial:", err)
				}
				continue
			}

			var data map[string]interface{}

			// Parse JSON message
			marshalErr := json.Unmarshal(message, &data)
			if marshalErr != nil {
				log.Println("Error parsing JSON:", marshalErr)
				continue
			}

			// Extract "name" field
			name, ok := data["name"].(string)
			if !ok {
				log.Println("No 'name' field in message")
				continue
			}

			// Process the received message
			handleMessage(rdbHoot, mongoHooter, mongoTradeOrder, name, message)
		}
	}
}

func handleMessage(rdbHoot rdblistener.HootRedisClient, mongoUser, mongoDB *mongo.Database, name string, message []byte) {
	switch name {
	case "event":
		log.Println("Event received")
		var event types.Event
		err := json.Unmarshal(message, &event)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		log.Println("Event", event)
		return
	//case "accBorrowingFeesGroupUpdated":
	//	log.Println("AccBorrowingFeesGroupUpdated")
	//case "accBorrowingFeesPairUpdated":
	//	log.Println("AccBorrowingFeesPairUpdated")
	case "collateralBalances":
		//log.Println("CollateralBalances")
		//var collateralBalances types.CollateralBalances
		//err := json.Unmarshal(message, &collateralBalances)
		//if err != nil {
		//	log.Println("Error parsing JSON:", err)
		//}
		//log.Println("CollateralBalances", collateralBalances)
		//return
	case "currentBlock":
		//var currentBlock CurrentBlock
		//err := json.Unmarshal(message, &currentBlock)
		//if err != nil {
		//	log.Println("Error parsing JSON:", err)
		//}
		//log.Println("CurrentBlock", currentBlock)
		//return
	case "currentL1Block":
		//var currentL1Block CurrentL1Block
		//err := json.Unmarshal(message, &currentL1Block)
		//if err != nil {
		//	log.Println("Error parsing JSON:", err)
		//}
		//log.Println("CurrentL1Block", currentL1Block)
		//return
	case "feeTiersTraderDailyPointsIncreased":
		//log.Println("FeeTiersTraderDailyPointsIncreased")
	case "feeTiersTraderInfoUpdated":
		//log.Println("FeesTiersTradersInfoUpdated")
	//case "feeTiersTraderTrailingPointsExpired":
	//	log.Println("FeeTiersTraderTrailingPointsExpired")
	case "liveEvent":
		log.Println("LiveEvent")
		var liveEvent types.LiveEvent
		err := json.Unmarshal(message, &liveEvent)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		log.Println("LiveEvent", liveEvent)
		return
	case "new-trade-history":
		log.Println("New TradeHistory")
		var tradeHistory types.TradeHistory
		err := json.Unmarshal(message, &tradeHistory)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		log.Println("TradeHistory", tradeHistory)
		return
	case "oiWindowUpdated":
		log.Println("OIWindowUpdated")
	case "openInterestGroupUpdated":
		log.Println("OpenInterestGroupUpdated")
	case "openInterestPairUpdated":
		log.Println("openInterestPairUpdated")
	case "openTrades":
		log.Println("openTrades")
		var openTrades types.OpenTrades
		err := json.Unmarshal(message, &openTrades)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		openTradesLowered := legacy.LowercaseUsers(openTrades.Value)
		rdbprocessing.ProcessOpenTrades(rdbHoot, ctx, openTradesLowered)
		return
	case "registerTrade":
		log.Println("registerTrade")
		var registerTrade types.RegisterTrade
		err := json.Unmarshal(message, &registerTrade)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		prcsErr := rdbprocessing.ProcessRegisterTrade(rdbHoot, ctx, registerTrade)
		if prcsErr != nil {
			log.Println("Error ProccessRegisterTrade: ", err)
		}
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		return
	case "unregisterTrade":
		log.Println("unregisterTrade")
		var unregisterTrade types.UnregisterTrade
		log.Println(message)
		err := json.Unmarshal(message, &unregisterTrade)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		prcsErr := rdbprocessing.ProcessUnregisterTrade(rdbHoot, ctx, unregisterTrade.Value)
		if prcsErr != nil {
			log.Println("Error ProcessUnregisterTrade: ", err)
		}
	case "tradingVariables":
		log.Println("tradingVariables")
		//var tradingVariables types.TradingVariables
		//err := json.Unmarshal(message, &tradingVariables)
		//if err != nil {
		//	log.Println("Error parsing JSON:", err)
		//}
		//log.Println("TradingVariables", tradingVariables)
		return
	//case "unconfirmedEvent":
	//	log.Println("unconfirmedEvent")
	// TODO: UpdateTrade is not even done
	case "updateTrade":
		log.Println("updateTrade")
		var updateTrade types.RegisterTrade
		err := json.Unmarshal(message, &updateTrade)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}
		log.Println("UpdateTrade", updateTrade)
		return
	}
}
