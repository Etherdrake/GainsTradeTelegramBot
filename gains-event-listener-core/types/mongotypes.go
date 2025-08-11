package types

import (
	"log"
	"strconv"
)

type OpenTradesMongo struct {
	Name  string           `json:"name"`
	Value []OpenTradeMongo `json:"value"`
}

type OpenTradeMongo struct {
	Trade          TradeMongo          `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}

type TradeItemMongo struct {
	ID             string              `bson:"_id,omitempty"`
	Trade          TradeMongo          `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}

type TradeMongo struct {
	User             string `json:"user"`
	Index            string `json:"index"`
	PairIndex        string `json:"pairIndex"`
	Leverage         string `json:"leverage"`
	Long             bool   `json:"long"`
	IsOpen           bool   `json:"isOpen"`
	CollateralIndex  string `json:"collateralIndex"`
	TradeType        string `json:"tradeType"`
	CollateralAmount string `json:"collateralAmount"`
	OpenPrice        string `json:"openPrice"`
	Tp               string `json:"tp"`
	Sl               string `json:"sl"`
}

type TradeInfoMongo struct {
	CreatedBlock       string `json:"createdBlock"`
	TpLastUpdatedBlock string `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock string `json:"slLastUpdatedBlock"`
	MaxSlippageP       string `json:"maxSlippageP"`
	LastOiUpdateTs     string `json:"lastOiUpdateTs"`
	CollateralPriceUsd string `json:"collateralPriceUsd"`
}

type InitialAccFeesMongo struct {
	AccPairFee  string `json:"accPairFee"`
	AccGroupFee string `json:"accGroupFee"`
	Block       string `json:"block"`
}

// ToCore Method to convert TradeItemMongo to TradeItemCore
func (mongo *TradeItemMongo) ToCore() TradeItemCore {
	var core TradeItemCore

	index, err := strconv.ParseUint(mongo.Trade.Index, 10, 32)
	if err != nil {
		log.Printf("failed to parse Index: %v", err)
		index = 0
	}

	pairIndex, err := strconv.ParseUint(mongo.Trade.PairIndex, 10, 16)
	if err != nil {
		log.Printf("failed to parse PairIndex: %v", err)
		pairIndex = 0
	}

	leverage, err := strconv.ParseFloat(mongo.Trade.Leverage, 32)
	if err != nil {
		log.Printf("failed to parse Leverage: %v", err)
		leverage = 0
	}

	collateralIndex, err := strconv.ParseUint(mongo.Trade.CollateralIndex, 10, 8)
	if err != nil {
		log.Printf("failed to parse CollateralIndex: %v", err)
		collateralIndex = 0
	}

	collateralAmount, err := strconv.ParseFloat(mongo.Trade.CollateralAmount, 64)
	if err != nil {
		log.Printf("failed to parse CollateralAmount: %v", err)
		collateralAmount = 0.0
	}

	openPrice, err := strconv.ParseFloat(mongo.Trade.OpenPrice, 64)
	if err != nil {
		log.Printf("failed to parse OpenPrice: %v", err)
		openPrice = 0.0
	}

	tp, err := strconv.ParseFloat(mongo.Trade.Tp, 64)
	if err != nil {
		log.Printf("failed to parse Tp: %v", err)
		tp = 0.0
	}

	sl, err := strconv.ParseFloat(mongo.Trade.Sl, 64)
	if err != nil {
		log.Printf("failed to parse Sl: %v", err)
		sl = 0.0
	}

	createdBlock, err := strconv.ParseUint(mongo.TradeInfo.CreatedBlock, 10, 64)
	if err != nil {
		log.Printf("failed to parse CreatedBlock: %v", err)
		createdBlock = 0
	}

	tpLastUpdatedBlock, err := strconv.ParseUint(mongo.TradeInfo.TpLastUpdatedBlock, 10, 64)
	if err != nil {
		log.Printf("failed to parse TpLastUpdatedBlock: %v", err)
		tpLastUpdatedBlock = 0
	}

	slLastUpdatedBlock, err := strconv.ParseUint(mongo.TradeInfo.SlLastUpdatedBlock, 10, 64)
	if err != nil {
		log.Printf("failed to parse SlLastUpdatedBlock: %v", err)
		slLastUpdatedBlock = 0
	}

	maxSlippageP, err := strconv.ParseUint(mongo.TradeInfo.MaxSlippageP, 10, 64)
	if err != nil {
		log.Printf("failed to parse MaxSlippageP: %v", err)
		maxSlippageP = 0
	}

	lastOiUpdateTs, err := strconv.ParseUint(mongo.TradeInfo.LastOiUpdateTs, 10, 64)
	if err != nil {
		log.Printf("failed to parse LastOiUpdateTs: %v", err)
		lastOiUpdateTs = 0
	}

	collateralPriceUsd, err := strconv.ParseFloat(mongo.TradeInfo.CollateralPriceUsd, 64)
	if err != nil {
		log.Printf("failed to parse CollateralPriceUsd: %v", err)
		collateralPriceUsd = 0.0
	}

	accPairFee, err := strconv.ParseFloat(mongo.InitialAccFees.AccPairFee, 64)
	if err != nil {
		log.Printf("failed to parse AccPairFee: %v", err)
		accPairFee = 0.0
	}

	accGroupFee, err := strconv.ParseFloat(mongo.InitialAccFees.AccGroupFee, 64)
	if err != nil {
		log.Printf("failed to parse AccGroupFee: %v", err)
		accGroupFee = 0.0
	}

	block, err := strconv.ParseUint(mongo.InitialAccFees.Block, 10, 64)
	if err != nil {
		log.Printf("failed to parse Block: %v", err)
		block = 0
	}

	tradeType, err := strconv.ParseUint(mongo.Trade.TradeType, 10, 64)
	if err != nil {
		log.Printf("failed to parse TradeType: %v", err)
		block = 0
	}

	core = TradeItemCore{
		ID: mongo.ID,
		Trade: TradeCore{
			User:             mongo.Trade.User,
			Index:            uint32(index),
			PairIndex:        uint16(pairIndex),
			Leverage:         float32(leverage),
			Long:             mongo.Trade.Long,
			IsOpen:           mongo.Trade.IsOpen,
			CollateralIndex:  uint8(collateralIndex),
			TradeType:        uint8(tradeType),
			CollateralAmount: collateralAmount,
			OpenPrice:        openPrice,
			Tp:               tp,
			Sl:               sl,
		},
		TradeInfo: TradeInfoCore{
			CreatedBlock:       createdBlock,
			TpLastUpdatedBlock: tpLastUpdatedBlock,
			SlLastUpdatedBlock: slLastUpdatedBlock,
			MaxSlippageP:       maxSlippageP,
			LastOiUpdateTs:     lastOiUpdateTs,
			CollateralPriceUsd: collateralPriceUsd,
		},
		InitialAccFees: InitialAccFeesCore{
			AccPairFee:  accPairFee,
			AccGroupFee: accGroupFee,
			Block:       block,
		},
	}
	return core
}
