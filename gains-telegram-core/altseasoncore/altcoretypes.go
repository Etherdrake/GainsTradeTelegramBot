package altseasoncore

import (
	"log"
	"math/big"
	"strconv"
)

type OpenTradesCore struct {
	Name  string           `json:"name"`
	Value []OpenTradeMongo `json:"value"`
}

type OpenTradeCore struct {
	Trade          TradeMongo          `json:"trade"`
	TradeInfo      TradeInfoMongo      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesMongo `json:"initialAccFees"`
}

type TradeCore struct {
	User             string  `json:"user"`
	Index            uint32  `json:"index"`
	PairIndex        uint16  `json:"pairIndex"`
	Leverage         float32 `json:"leverage"`
	Long             bool    `json:"long"`
	IsOpen           bool    `json:"isOpen"`
	CollateralIndex  uint8   `json:"collateralIndex"`
	TradeType        uint8   `json:"tradeType"`
	CollateralAmount float64 `json:"collateralAmount"`
	OpenPrice        float64 `json:"openPrice"`
	Tp               float64 `json:"tp"`
	Sl               float64 `json:"sl"`
}

type TradeInfoCore struct {
	CreatedBlock       uint64  `json:"createdBlock"`
	TpLastUpdatedBlock uint64  `json:"tpLastUpdatedBlock"`
	SlLastUpdatedBlock uint64  `json:"slLastUpdatedBlock"`
	MaxSlippageP       uint64  `json:"maxSlippageP"`
	LastOiUpdateTs     uint64  `json:"lastOiUpdateTs"`
	CollateralPriceUsd float64 `json:"collateralPriceUsd"`
}

type InitialAccFeesCore struct {
	AccPairFee  float64 `json:"accPairFee"`
	AccGroupFee float64 `json:"accGroupFee"`
	Block       uint64  `json:"block"`
}

type TradeItemCore struct {
	ID             string             `bson:"_id,omitempty"`
	Trade          TradeCore          `json:"trade"`
	TradeInfo      TradeInfoCore      `json:"tradeInfo"`
	InitialAccFees InitialAccFeesCore `json:"initialAccFees"`
}

func (iac InitialAccFeesCore) IntoConverted() InitialAccFeesConverted {
	return InitialAccFeesConverted{
		AccPairFee:  iac.AccPairFee,
		AccGroupFee: iac.AccGroupFee,
		Block:       iac.Block,
	}
}

func (t TradeItemCore) IntoDisplay() TradeItemDisplay {
	return TradeItemDisplay{
		ID:             t.ID,
		Trade:          t.Trade.IntoDisplay(),
		TradeInfo:      t.TradeInfo.IntoDisplay(),
		InitialAccFees: t.InitialAccFees.IntoDisplay(),
	}
}

func (t TradeCore) IntoDisplay() TradeDisplay {
	eighteenFloat := big.NewFloat(1e18)
	tenFloat := big.NewFloat(1e10)
	sixFloat := big.NewFloat(1e6)

	// Call GetDecimals which retrieves the amount of decimals for this particular pairIndex

	switch t.CollateralIndex {
	// Native
	case 0:
		log.Println("IntoDisplay not implemented for CollateralIndex 0")
	// DAI
	case 1:
		collateralAmount := big.NewFloat(t.CollateralAmount)
		collateralAmount.Quo(collateralAmount, eighteenFloat)
		collateralAmountFloat64, _ := collateralAmount.Float64()

		openPrice := big.NewFloat(t.OpenPrice)
		openPrice.Quo(openPrice, tenFloat)
		openPriceFloat64, _ := openPrice.Float64()

		tp := big.NewFloat(t.Tp)
		tp.Quo(tp, tenFloat)
		tpFloat64, _ := tp.Float64()

		sl := big.NewFloat(t.Sl)
		sl.Quo(sl, tenFloat)
		slFloat64, _ := sl.Float64()

		// Convert Leverage
		leverageFloat := float64(t.Leverage) / 1000.0

		return TradeDisplay{
			User:             t.User,
			Index:            strconv.FormatUint(uint64(t.Index), 10),
			PairIndex:        strconv.FormatUint(uint64(t.PairIndex), 10),
			Leverage:         strconv.FormatFloat(leverageFloat, 'f', -1, 64),
			Long:             strconv.FormatBool(t.Long),
			IsOpen:           strconv.FormatBool(t.IsOpen),
			CollateralIndex:  strconv.FormatUint(uint64(t.CollateralIndex), 10),
			TradeType:        strconv.FormatUint(uint64(t.TradeType), 10),
			CollateralAmount: strconv.FormatFloat(collateralAmountFloat64, 'f', -1, 64),
			OpenPrice:        strconv.FormatFloat(openPriceFloat64, 'f', -1, 64),
			Tp:               strconv.FormatFloat(tpFloat64, 'f', -1, 64),
			Sl:               strconv.FormatFloat(slFloat64, 'f', -1, 64),
		}
	// WETH
	case 2:
		collateralAmount := big.NewFloat(t.CollateralAmount)
		collateralAmount.Quo(collateralAmount, eighteenFloat)
		collateralAmountFloat64, _ := collateralAmount.Float64()

		openPrice := big.NewFloat(t.OpenPrice)
		openPrice.Quo(openPrice, tenFloat)
		openPriceFloat64, _ := openPrice.Float64()

		tp := big.NewFloat(t.Tp)
		tp.Quo(tp, tenFloat)
		tpFloat64, _ := tp.Float64()

		sl := big.NewFloat(t.Sl)
		sl.Quo(sl, tenFloat)
		slFloat64, _ := sl.Float64()

		// Convert Leverage
		leverageFloat := float64(t.Leverage) / 1000.0

		return TradeDisplay{
			User:             t.User,
			Index:            strconv.FormatUint(uint64(t.Index), 10),
			PairIndex:        strconv.FormatUint(uint64(t.PairIndex), 10),
			Leverage:         strconv.FormatFloat(leverageFloat, 'f', -1, 64),
			Long:             strconv.FormatBool(t.Long),
			IsOpen:           strconv.FormatBool(t.IsOpen),
			CollateralIndex:  strconv.FormatUint(uint64(t.CollateralIndex), 10),
			TradeType:        strconv.FormatUint(uint64(t.TradeType), 10),
			CollateralAmount: strconv.FormatFloat(collateralAmountFloat64, 'f', -1, 64),
			OpenPrice:        strconv.FormatFloat(openPriceFloat64, 'f', -1, 64),
			Tp:               strconv.FormatFloat(tpFloat64, 'f', -1, 64),
			Sl:               strconv.FormatFloat(slFloat64, 'f', -1, 64),
		}
	// USDC
	case 3:
		collateralAmount := big.NewFloat(t.CollateralAmount)
		collateralAmount.Quo(collateralAmount, sixFloat)
		collateralAmountFloat64, _ := collateralAmount.Float64()

		openPrice := big.NewFloat(t.OpenPrice)
		openPrice.Quo(openPrice, tenFloat)
		openPriceFloat64, _ := openPrice.Float64()

		tp := big.NewFloat(t.Tp)
		tp.Quo(tp, tenFloat)
		tpFloat64, _ := tp.Float64()

		sl := big.NewFloat(t.Sl)
		sl.Quo(sl, tenFloat)
		slFloat64, _ := sl.Float64()

		// Convert Leverage
		leverageFloat := float64(t.Leverage) / 1000.0

		return TradeDisplay{
			User:             t.User,
			Index:            strconv.FormatUint(uint64(t.Index), 10),
			PairIndex:        strconv.FormatUint(uint64(t.PairIndex), 10),
			Leverage:         strconv.FormatFloat(leverageFloat, 'f', -1, 64),
			Long:             strconv.FormatBool(t.Long),
			IsOpen:           strconv.FormatBool(t.IsOpen),
			CollateralIndex:  strconv.FormatUint(uint64(t.CollateralIndex), 10),
			TradeType:        strconv.FormatUint(uint64(t.TradeType), 10),
			CollateralAmount: strconv.FormatFloat(collateralAmountFloat64, 'f', -1, 64),
			OpenPrice:        strconv.FormatFloat(openPriceFloat64, 'f', -1, 64),
			Tp:               strconv.FormatFloat(tpFloat64, 'f', -1, 64),
			Sl:               strconv.FormatFloat(slFloat64, 'f', -1, 64),
		}
	}
	return TradeDisplay{}
}
func (t TradeInfoCore) IntoDisplay() TradeInfoDisplay {
	return TradeInfoDisplay{
		CreatedBlock:       strconv.FormatUint(t.CreatedBlock, 10),
		TpLastUpdatedBlock: strconv.FormatUint(t.TpLastUpdatedBlock, 10),
		SlLastUpdatedBlock: strconv.FormatUint(t.SlLastUpdatedBlock, 10),
		MaxSlippageP:       strconv.FormatUint(t.MaxSlippageP, 10),
		LastOiUpdateTs:     strconv.FormatUint(t.LastOiUpdateTs, 10),
		CollateralPriceUsd: strconv.FormatFloat(t.CollateralPriceUsd, 'f', -1, 64),
	}
}

func (t InitialAccFeesCore) IntoDisplay() InitialAccFeesDisplay {
	accPairFeeStr := strconv.FormatFloat(t.AccPairFee, 'f', -1, 64)
	accGroupFeeStr := strconv.FormatFloat(t.AccGroupFee, 'f', -1, 64)
	blockStr := strconv.FormatUint(t.Block, 10)

	initialAccFeesDisplay := InitialAccFeesDisplay{
		AccPairFee:  accPairFeeStr,
		AccGroupFee: accGroupFeeStr,
		Block:       blockStr,
	}

	return initialAccFeesDisplay
}
