package types

import (
	"log"
	"math/big"
	"strconv"
)

type TradeItemInternal struct {
	ID             string                 `json:"id"`
	Trade          TradeInternal          `json:"trade"`
	TradeInfo      TradeInfoInternal      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesInternal `json:"initialAccFees"`
}

type TradeInternal struct {
	User             string  `json:"user"`
	Index            uint32  `json:"index"`
	PairIndex        uint16  `json:"pairIndex"`
	Leverage         float64 `json:"leverage"`
	Long             bool    `json:"long"`
	IsOpen           bool    `json:"isOpen"`
	CollateralIndex  uint8   `json:"collateralIndex"`
	TradeType        uint8   `json:"tradeType"`
	CollateralAmount float64 `json:"collateralAmount"`
	OpenPrice        float64 `json:"openPrice"`
	Tp               float64 `json:"tp"`
	Sl               float64 `json:"sl"`
}

type TradeInfoInternal struct {
	CreatedBlock       uint64  `json:"createdBlock"`
	TpLastUpdatedBlock uint64  `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock uint64  `json:"slLastUpdatedBlock"`
	MaxSlippageP       float64 `json:"maxSlippageP"`
	LastOiUpdateTs     string  `json:"lastOiUpdateTs"`
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
}

type InitialAccFeesInternal struct {
	AccPairFee  float64 `json:"accPairFee"`
	AccGroupFee float64 `json:"accGroupFee"`
	Block       uint64  `json:"block"`
}

func (t TradeCore) IntoInternal() TradeInternal {
	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)
	sixFloat := big.NewFloat(1e6)

	var collateralAmount, openPrice, tp, sl *big.Float
	var collateralAmountFloat64, openPriceFloat64, tpFloat64, slFloat64 float64

	switch t.CollateralIndex {
	case 0:
		log.Println("IntoInternal not implemented for CollateralIndex 0")
	case 1:
		collateralAmount = big.NewFloat(t.CollateralAmount).Quo(big.NewFloat(t.CollateralAmount), eighteenFloat)
		collateralAmountFloat64, _ = collateralAmount.Float64()

		openPrice = big.NewFloat(t.OpenPrice).Quo(big.NewFloat(t.OpenPrice), tenFloat)
		openPriceFloat64, _ = openPrice.Float64()

		tp = big.NewFloat(t.Tp).Quo(big.NewFloat(t.Tp), tenFloat)
		tpFloat64, _ = tp.Float64()

		sl = big.NewFloat(t.Sl).Quo(big.NewFloat(t.Sl), tenFloat)
		slFloat64, _ = sl.Float64()
	case 2:
		collateralAmount = big.NewFloat(t.CollateralAmount).Quo(big.NewFloat(t.CollateralAmount), eighteenFloat)
		collateralAmountFloat64, _ = collateralAmount.Float64()

		openPrice = big.NewFloat(t.OpenPrice).Quo(big.NewFloat(t.OpenPrice), tenFloat)
		openPriceFloat64, _ = openPrice.Float64()

		tp = big.NewFloat(t.Tp).Quo(big.NewFloat(t.Tp), tenFloat)
		tpFloat64, _ = tp.Float64()

		sl = big.NewFloat(t.Sl).Quo(big.NewFloat(t.Sl), tenFloat)
		slFloat64, _ = sl.Float64()
	case 3:
		collateralAmount = big.NewFloat(t.CollateralAmount).Quo(big.NewFloat(t.CollateralAmount), sixFloat)
		collateralAmountFloat64, _ = collateralAmount.Float64()

		openPrice = big.NewFloat(t.OpenPrice).Quo(big.NewFloat(t.OpenPrice), tenFloat)
		openPriceFloat64, _ = openPrice.Float64()

		tp = big.NewFloat(t.Tp).Quo(big.NewFloat(t.Tp), tenFloat)
		tpFloat64, _ = tp.Float64()

		sl = big.NewFloat(t.Sl).Quo(big.NewFloat(t.Sl), tenFloat)
		slFloat64, _ = sl.Float64()
	}

	leverageFloat := float64(t.Leverage) / 1000.0

	return TradeInternal{
		User:             t.User,
		Index:            uint32(t.Index),
		PairIndex:        uint16(t.PairIndex),
		Leverage:         leverageFloat,
		Long:             t.Long,
		IsOpen:           t.IsOpen,
		CollateralIndex:  t.CollateralIndex,
		TradeType:        t.TradeType,
		CollateralAmount: collateralAmountFloat64,
		OpenPrice:        openPriceFloat64,
		Tp:               tpFloat64,
		Sl:               slFloat64,
	}
}

func (t TradeInfoCore) IntoInternal() TradeInfoInternal {
	return TradeInfoInternal{
		CreatedBlock:       t.CreatedBlock,
		TpLastUpdatedBlock: t.TpLastUpdatedBlock,
		SlLastUpdatedBlock: t.SlLastUpdatedBlock,
		MaxSlippageP:       float64(t.MaxSlippageP),
		LastOiUpdateTs:     strconv.FormatUint(t.LastOiUpdateTs, 10),
		CollateralPriceUsd: t.CollateralPriceUsd,
	}
}

func (t InitialAccFeesCore) IntoInternal() InitialAccFeesInternal {
	return InitialAccFeesInternal{
		AccPairFee:  t.AccPairFee,
		AccGroupFee: t.AccGroupFee,
		Block:       t.Block,
	}
}

func (t TradeItemCore) IntoInternal() TradeItemInternal {
	return TradeItemInternal{
		ID:             t.ID,
		Trade:          t.Trade.IntoInternal(),
		TradeInfo:      t.TradeInfo.IntoInternal(),
		InitialAccFees: t.InitialAccFees.IntoInternal(),
	}
}

func (t TradeInternal) IntoDisplay() TradeDisplay {
	return TradeDisplay{
		User:             t.User,
		Index:            strconv.FormatUint(uint64(t.Index), 10),
		PairIndex:        strconv.FormatUint(uint64(t.PairIndex), 10),
		Leverage:         strconv.FormatFloat(t.Leverage, 'f', -1, 64),
		Long:             strconv.FormatBool(t.Long),
		IsOpen:           strconv.FormatBool(t.IsOpen),
		CollateralIndex:  strconv.FormatUint(uint64(t.CollateralIndex), 10),
		TradeType:        strconv.FormatUint(uint64(t.TradeType), 10),
		CollateralAmount: strconv.FormatFloat(t.CollateralAmount, 'f', -1, 64),
		OpenPrice:        strconv.FormatFloat(t.OpenPrice, 'f', -1, 64),
		Tp:               strconv.FormatFloat(t.Tp, 'f', -1, 64),
		Sl:               strconv.FormatFloat(t.Sl, 'f', -1, 64),
	}
}

func (t TradeInfoInternal) IntoDisplay() TradeInfoDisplay {
	return TradeInfoDisplay{
		CreatedBlock:       strconv.FormatUint(t.CreatedBlock, 10),
		TpLastUpdatedBlock: strconv.FormatUint(t.TpLastUpdatedBlock, 10),
		SlLastUpdatedBlock: strconv.FormatUint(t.SlLastUpdatedBlock, 10),
		MaxSlippageP:       strconv.FormatFloat(t.MaxSlippageP, 'f', -1, 64),
		LastOiUpdateTs:     t.LastOiUpdateTs,
		CollateralPriceUsd: strconv.FormatFloat(t.CollateralPriceUsd, 'f', -1, 64),
	}
}

func (t InitialAccFeesInternal) IntoDisplay() InitialAccFeesDisplay {
	return InitialAccFeesDisplay{
		AccPairFee:  strconv.FormatFloat(t.AccPairFee, 'f', -1, 64),
		AccGroupFee: strconv.FormatFloat(t.AccGroupFee, 'f', -1, 64),
		Block:       strconv.FormatUint(t.Block, 10),
	}
}

func (t TradeItemInternal) IntoDisplay() TradeItemDisplay {
	return TradeItemDisplay{
		ID:             t.ID,
		Trade:          t.Trade.IntoDisplay(),
		TradeInfo:      t.TradeInfo.IntoDisplay(),
		InitialAccFees: t.InitialAccFees.IntoDisplay(),
	}
}
