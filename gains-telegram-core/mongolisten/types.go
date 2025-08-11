package mongolisten

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/big"
	"strconv"
)

type OpenLimitOrderBigString struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Trader           string             `bson:"trader" json:"trader"`
	PairIndex        string             `bson:"pair_index" json:"pair_index"`
	Index            string             `bson:"index" json:"index"`
	PositionSize     string             `bson:"position_size" json:"position_size"`
	SpreadReductionP string             `bson:"spread_reduction_p" json:"spread_reduction_p"`
	Buy              bool               `bson:"buy" json:"buy"`
	Leverage         string             `bson:"leverage" json:"leverage"`
	Tp               string             `bson:"tp" json:"tp"`
	Sl               string             `bson:"sl" json:"sl"`
	MinPrice         string             `bson:"min_price" json:"min_price"`
	MaxPrice         string             `bson:"max_price" json:"max_price"`
	Block            string             `bson:"block" json:"block"`
	TokenID          string             `bson:"token_id" json:"token_id"`
	CollateralToken  string             `bson:"collateral_token" json:"collateral_token"`
}

type OpenLimitOrderConverted struct {
	Trader           string  `json:"trader"`
	PairIndex        int64   `json:"pair_index"`
	Index            int64   `json:"index"`
	PositionSize     float64 `json:"position_size"`
	SpreadReductionP uint    `json:"spread_reduction_p"`
	Buy              bool    `json:"buy"`
	Leverage         int64   `json:"leverage"`
	Tp               float64 `json:"tp"`
	Sl               float64 `json:"sl"`
	MinPrice         float64 `json:"min_price"`
	MaxPrice         float64 `json:"max_price"`
	Block            int64   `json:"block"`
	TokenID          uint    `json:"token_id"`
	CollateralToken  uint    `json:"collateral_token"`
}

func ConvertBigLimitToLimit(bigOrder OpenLimitOrderBigString) OpenLimitOrderConverted {
	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)

	positionSize, success := new(big.Int).SetString(bigOrder.PositionSize, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	positionSizeDai := new(big.Float).SetInt(positionSize)
	positionSizeDai.Quo(positionSizeDai, eighteenFloat)
	positionSizeDaiF64, _ := positionSizeDai.Float64()

	takeProfit, success := new(big.Int).SetString(bigOrder.Tp, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	tp := new(big.Float).SetInt(takeProfit)
	tp.Quo(tp, tenFloat)
	tp.SetPrec(10)
	tpFloat64, _ := tp.Float64()

	stopLoss, success := new(big.Int).SetString(bigOrder.Sl, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	sl := new(big.Float).SetInt(stopLoss)
	sl.Quo(sl, tenFloat)
	sl.SetPrec(10)
	slFloat64, _ := sl.Float64()

	minPrice, success := new(big.Int).SetString(bigOrder.MinPrice, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	mcPrce := new(big.Float).SetInt(minPrice)
	mcPrce.Quo(mcPrce, tenFloat)
	mcPrce.SetPrec(10)
	mcPrce64, _ := mcPrce.Float64()

	maxPrice, success := new(big.Int).SetString(bigOrder.MaxPrice, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	mxPrce := new(big.Float).SetInt(maxPrice)
	mxPrce.Quo(mxPrce, tenFloat)
	mxPrce.SetPrec(10)
	maxPrice64, _ := mxPrce.Float64()

	// Convert string fields to int64
	pairIndexInt64, err := strconv.ParseInt(bigOrder.PairIndex, 10, 64)
	if err != nil {
		log.Println("Error ConvertBigLimitToLimit: ", err)
	}
	indexInt64, err := strconv.ParseInt(bigOrder.Index, 10, 64)
	if err != nil {
		log.Println("Error ConvertBigLimitToLimit: ", err)
	}
	leverageInt64, err := strconv.ParseInt(bigOrder.Leverage, 10, 64)
	if err != nil {
		log.Println("Error ConvertBigLimitToLimit: ", err)
	}
	blockInt64, err := strconv.ParseInt(bigOrder.Block, 10, 64)
	if err != nil {
		log.Println("Error ConvertBigLimitToLimit: ", err)
	}

	collateralInt, err := strconv.ParseUint(bigOrder.CollateralToken, 10, 64)
	if err != nil {
		log.Println("Error ConvertBigLimitToLimit: ", err)
	}

	limitOrder := OpenLimitOrderConverted{
		Trader:           bigOrder.Trader,
		PairIndex:        pairIndexInt64,
		Index:            indexInt64,
		PositionSize:     positionSizeDaiF64,
		SpreadReductionP: 0,
		Buy:              bigOrder.Buy,
		Leverage:         leverageInt64,
		Tp:               tpFloat64,
		Sl:               slFloat64,
		MinPrice:         mcPrce64,
		MaxPrice:         maxPrice64,
		Block:            blockInt64,
		TokenID:          0,
		CollateralToken:  uint(collateralInt),
	}
	return limitOrder
}
