package main

import (
	"GainsListenerMumbai/abitools"
	"GainsListenerMumbai/database"
	"GainsListenerMumbai/streamprocessor"
	"GainsListenerMumbai/transformers"
	"context"
	"fmt"
	gnsconstants "github.com/Etherdrake/gns-constants"
	"github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strings"
	"sync"
)

var ctx = context.Background()

var mongoCtx = context.Background()

func InitRedisPriceStream() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // change this to your Redis server address
		DB:   1,                // using logical database 1 for pricedata
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		return nil
	}

	fmt.Println("Connected to Redis!")
	return rdb
}

var polyGroup sync.WaitGroup
var arbiGroup sync.WaitGroup

var isArb bool
var isMumbai bool

var (
	openLimitChan         = make(chan transformers.OpenLimitPlacedEventTransform)
	limitExecutedChan     = make(chan transformers.LimitExecutedTransform)
	nftOrderCanceledChan  = make(chan transformers.NftOrderCanceledTransform)
	openLimitCanceledChan = make(chan transformers.OpenLimitCanceledEventTransform)
	nftOrderInitiatedChan = make(chan transformers.NftOrderInitiatedEventTransform)

	marketOrderInitiatedChan = make(chan transformers.MarketOrderInitiatedEventTransform)
	marketExecutedChan       = make(chan transformers.MarketExecutedTransform)
	marketOpenCanceledChan   = make(chan transformers.MarketOpenCanceledTransform)
	marketCloseCanceledChan  = make(chan transformers.MarketCloseCanceledTransform)

	TpUpdatedChan = make(chan transformers.TpUpdatedEventTransform)
	SlUpdatedChan = make(chan transformers.SlUpdatedEventTransform)
)

func main() {
	args := os.Args
	if args[1] == "mumbai" {
		isMumbai = true
		isArb = false
	} else {
		log.Fatal("ERROR: Invalid command-line argument supplied. Please add `arbitrum` or `polygon` to select the chain you want to listen on.")
	}

	var chainstackWSS string
	chainstackWSS = "wss://ws-nd-447-606-550.p2pify.com/105b8aac8c8002a9159ffabfda070cc9"

	client, err := ethclient.Dial(chainstackWSS)
	if err != nil {
		log.Fatal(err)
	}

	var mongoDB *mongo.Database
	mongoDB, err = database.ConnectToMongoDB(gnsconstants.ConnectionString, gnsconstants.DbNameMumbai)
	if err != nil {
		log.Fatal("MongoDB could not be initialized")
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Price Stream Redis client
	rdbPrice := InitRedisPriceStream()
	defer rdbPrice.Close()

	tradingBot, err := tgbotapi.NewBotAPI(gnsconstants.EventListenerTradingTelegramToken)
	if err != nil {
		log.Panic(err)
	}

	callbackBot, err := tgbotapi.NewBotAPI(gnsconstants.EventListenerCallbacksTelegramToken)
	if err != nil {
		fmt.Println("Error encountered with TG BOT: ", err)
	}

	amalgamatedBot, err := tgbotapi.NewBotAPI(gnsconstants.EventlistenerAmalgamatedToken)
	if err != nil {
		fmt.Println("Error encountered with TG BOT: ", err)
	}

	for i := 0; i < 3; i++ {
		var collat string
		if i == 0 {
			collat = "DAI"
		}
		if i == 1 {
			collat = "WETH"
		}
		if i == 2 {
			collat = "USDC"
		}
		// Launch Arbitrum listeners
		arbiGroup.Add(2)
		go callbackListener(
			client,
			rdbPrice,
			callbackBot,
			openLimitChan,
			limitExecutedChan,
			openLimitCanceledChan,
			nftOrderInitiatedChan,
			nftOrderCanceledChan,
			marketOrderInitiatedChan,
			marketExecutedChan,
			marketOpenCanceledChan,
			marketCloseCanceledChan,
			collat,
		)

		go tradingListener(
			client,
			rdbPrice,
			tradingBot,
			openLimitChan,
			limitExecutedChan,
			openLimitCanceledChan,
			nftOrderInitiatedChan,
			nftOrderCanceledChan,
			marketOrderInitiatedChan,
			marketExecutedChan,
			marketOpenCanceledChan,
			marketCloseCanceledChan,
			TpUpdatedChan,
			SlUpdatedChan,
			collat,
		)

		// Run the matcher for the provided chains
		go streamprocessor.MatchLimitOrders(amalgamatedBot, mongoDB, openLimitChan, limitExecutedChan, openLimitCanceledChan, nftOrderInitiatedChan, nftOrderCanceledChan, collat)
		go streamprocessor.MatchMarketOrders(amalgamatedBot, mongoDB, marketOrderInitiatedChan, marketExecutedChan, marketOpenCanceledChan, marketCloseCanceledChan, collat)
		go streamprocessor.MatchChanges(mongoClient, TpUpdatedChan, SlUpdatedChan, collat)
	}

	arbiGroup.Wait()
}
func tradingListener(
	client *ethclient.Client,
	rdbPrice *redis.Client,
	bot *tgbotapi.BotAPI,
	openLimitChan chan transformers.OpenLimitPlacedEventTransform,
	limitExecutedChan chan transformers.LimitExecutedTransform,
	openLimitCanceledChan chan transformers.OpenLimitCanceledEventTransform,
	nftOrderInitiatedChan chan transformers.NftOrderInitiatedEventTransform,
	nftOrderCanceledChan chan transformers.NftOrderCanceledTransform,
	marketOrderInitiatedChan chan transformers.MarketOrderInitiatedEventTransform,
	marketExecutedChan chan transformers.MarketExecutedTransform,
	marketOpenCanceledChan chan transformers.MarketOpenCanceledTransform,
	marketCloseCanceledChan chan transformers.MarketCloseCanceledTransform,
	TpUpdatedChan chan transformers.TpUpdatedEventTransform,
	SlUpdatedChan chan transformers.SlUpdatedEventTransform,
	multiCollat string,

) {

	var contractAddress common.Address
	var contractABI string
	if multiCollat == "DAI" {
		contractAddress = common.HexToAddress(gnsconstants.GainsTradingAddressDaiMumbai)
		// Read the contents of the ABI file
		contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsTradingAbiPath)
	} else {
		if multiCollat == "WETH" {
			contractAddress = common.HexToAddress(gnsconstants.GainsTradingAddressWethMumbai)
			// Read the contents of the ABI file
			contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsTradingAbiPath)
		} else {
			if multiCollat == "USDC" {
				contractAddress = common.HexToAddress(gnsconstants.GainsTradingAddressUsdcMumbai)
				// Read the contents of the ABI file
				contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsTradingAbiPath)
			}
		}
	}

	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Println("Error parsing the abi: ", err)
	}

	// Create a channel to receive events
	eventChannel := make(chan types.Log)

	// Start the event listener
	go listenForTradingEvents(
		client,
		bot,
		contractAddress,
		parsedABI,
		eventChannel,
		openLimitChan,
		limitExecutedChan,
		openLimitCanceledChan,
		nftOrderInitiatedChan,
		marketOrderInitiatedChan,
		marketExecutedChan,
		marketOpenCanceledChan,
		marketCloseCanceledChan,
		TpUpdatedChan,
		SlUpdatedChan,
		multiCollat,
	)
}

func callbackListener(
	client *ethclient.Client,
	rdbPrice *redis.Client,
	bot *tgbotapi.BotAPI,
	openLimitChan chan transformers.OpenLimitPlacedEventTransform,
	limitExecutedChan chan transformers.LimitExecutedTransform,
	openLimitCanceledChan chan transformers.OpenLimitCanceledEventTransform,
	nftOrderInitiatedChan chan transformers.NftOrderInitiatedEventTransform,
	nftOrderCanceledChan chan transformers.NftOrderCanceledTransform,
	marketOrderInitiatedChan chan transformers.MarketOrderInitiatedEventTransform,
	marketExecutedChan chan transformers.MarketExecutedTransform,
	marketOpenCanceledChan chan transformers.MarketOpenCanceledTransform,
	marketCloseCanceledChan chan transformers.MarketCloseCanceledTransform,
	multiCollat string,
) {
	var contractAddress common.Address
	var contractABI string
	var err error
	if multiCollat == "DAI" {
		contractAddress = common.HexToAddress(gnsconstants.GainsCallbackAddressDaiMumbai)
		// Read the contents of the ABI file
		contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsCallbacksAbiPath)
	} else {
		if multiCollat == "WETH" {
			contractAddress = common.HexToAddress(gnsconstants.GainsCallbackAddressWethMumbai)
			// Read the contents of the ABI file
			contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsCallbacksAbiPath)
		} else {
			if multiCollat == "USDC" {
				contractAddress = common.HexToAddress(gnsconstants.GainsCallbackAddressUsdcMumbai)
				// Read the contents of the ABI file
				contractABI, _ = abitools.ReadABIFile(gnsconstants.GainsCallbacksAbiPath)
			}
		}
	}

	// Parse the contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Println("Error parsing the ABI: ", err)
	}

	// Create a channel to receive events
	eventChannel := make(chan types.Log)

	// Start the event listener
	go listenForCallbackEvents(
		client,
		bot,
		rdbPrice,
		contractAddress,
		parsedABI,
		eventChannel,
		openLimitChan,
		limitExecutedChan,
		openLimitCanceledChan,
		nftOrderInitiatedChan,
		marketOrderInitiatedChan,
		marketExecutedChan,
		marketOpenCanceledChan,
		marketCloseCanceledChan,
		multiCollat,
	)

	//// Wait for events
	//for {
	//	select {
	//	case event := <-eventChannel:
	//		fmt.Println("Received TxHash Callbacks:", event.TxHash)
	//	}
	//}
}
