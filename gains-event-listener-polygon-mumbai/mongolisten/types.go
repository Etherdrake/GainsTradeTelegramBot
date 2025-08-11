package mongolisten

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
)

type OpenLimitOrderBigString struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Trader           string             `bson:"trader" json:"trader"`
	PairIndex        int64              `bson:"pair_index" json:"pair_index"`
	Index            int64              `bson:"index" json:"index"`
	PositionSize     string             `bson:"position_size" json:"position_size"`
	SpreadReductionP uint               `bson:"spread_reduction_p" json:"spread_reduction_p"`
	Buy              bool               `bson:"buy" json:"buy"`
	Leverage         int64              `bson:"leverage" json:"leverage"`
	Tp               string             `bson:"tp" json:"tp"`
	Sl               string             `bson:"sl" json:"sl"`
	MinPrice         string             `bson:"min_price" json:"min_price"`
	MaxPrice         string             `bson:"max_price" json:"max_price"`
	Block            int64              `bson:"block" json:"block"`
	TokenID          uint               `bson:"token_id" json:"token_id"`
	CollateralToken  uint               `bson:"collateral_token" json:"collateral_token"`
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
	tpFloat64, _ := tp.Float64()

	stopLoss, success := new(big.Int).SetString(bigOrder.Sl, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	sl := new(big.Float).SetInt(stopLoss)
	sl.Quo(sl, tenFloat)
	slFloat64, _ := sl.Float64()

	minPrice, success := new(big.Int).SetString(bigOrder.MinPrice, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	mcPrce := new(big.Float).SetInt(minPrice)
	mcPrce.Quo(mcPrce, tenFloat)
	mcPrce64, _ := mcPrce.Float64()

	maxPrice, success := new(big.Int).SetString(bigOrder.MaxPrice, 10)
	if !success {
		return OpenLimitOrderConverted{}
	}
	mxPrce := new(big.Float).SetInt(maxPrice)
	mxPrce.Quo(mxPrce, tenFloat)
	maxPrice64, _ := mxPrce.Float64()

	limitOrder := OpenLimitOrderConverted{
		Trader:           bigOrder.Trader,
		PairIndex:        bigOrder.PairIndex,
		Index:            bigOrder.Index,
		PositionSize:     positionSizeDaiF64,
		SpreadReductionP: 0,
		Buy:              bigOrder.Buy,
		Leverage:         bigOrder.Leverage,
		Tp:               tpFloat64,
		Sl:               slFloat64,
		MinPrice:         mcPrce64,
		MaxPrice:         maxPrice64,
		Block:            bigOrder.Block,
		TokenID:          0,
		CollateralToken:  bigOrder.CollateralToken,
	}
	return limitOrder
}

type TradeFinal struct {
	Trader          string  `json:"trader" bson:"trader"`
	PairIndex       int64   `json:"pairIndex" bson:"pair_index"`
	Index           int64   `json:"index" bson:"index"`
	InitialPosToken float64 `json:"initialPosToken" bson:"initial_pos_token"`
	PositionSizeDai float64 `json:"positionSizeDai" bson:"position_size_dai"`
	OpenPrice       float64 `json:"openPrice" bson:"open_price"`
	Buy             bool    `json:"buy" bson:"buy"`
	Leverage        int64   `json:"leverage" bson:"leverage"`
	Tp              float64 `json:"tp" bson:"tp"`
	Sl              float64 `json:"sl" bson:"sl"`
	CollateralToken string  `json:"collateral_token" bson:"collateral_token"`
}
