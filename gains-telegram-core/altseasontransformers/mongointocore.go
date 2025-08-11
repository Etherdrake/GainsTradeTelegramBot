package altseasontransformers

import (
	"HootCore/altseasoncore"
	"log"
	"strconv"
)

func MongoIntoCore(mongo altseasoncore.TradeItemMongo) altseasoncore.TradeItemCore {
	var core altseasoncore.TradeItemCore

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

	var lastOiUpdateTs uint64
	if len(mongo.TradeInfo.LastOiUpdateTs) > 0 {
		lastOiUpdateTs, err = strconv.ParseUint(mongo.TradeInfo.LastOiUpdateTs, 10, 64)
		if err != nil {
			log.Printf("failed to parse LastOiUpdateTs: %v", err)
			lastOiUpdateTs = 0
		}
	} else {
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

	core = altseasoncore.TradeItemCore{
		ID: mongo.ID,
		Trade: altseasoncore.TradeCore{
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
		TradeInfo: altseasoncore.TradeInfoCore{
			CreatedBlock:       createdBlock,
			TpLastUpdatedBlock: tpLastUpdatedBlock,
			SlLastUpdatedBlock: slLastUpdatedBlock,
			MaxSlippageP:       maxSlippageP,
			LastOiUpdateTs:     lastOiUpdateTs,
			CollateralPriceUsd: collateralPriceUsd,
		},
		InitialAccFees: altseasoncore.InitialAccFeesCore{
			AccPairFee:  accPairFee,
			AccGroupFee: accGroupFee,
			Block:       block,
		},
	}

	return core
}

func MongoArrIntoCoreArr(mongoArr []altseasoncore.TradeItemMongo) []altseasoncore.TradeItemCore {
	var coreArr []altseasoncore.TradeItemCore

	for _, mongoItem := range mongoArr {
		coreArr = append(coreArr, MongoIntoCore(mongoItem))
	}

	return coreArr
}
