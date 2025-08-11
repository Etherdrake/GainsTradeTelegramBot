package altseasoncore

import (
	"log"
	"math"
)

func GetBorrowingFee(
	posDai float64,
	pairIndex int,
	long bool,
	initialAccFees InitialAccFeesConverted,
	context GainsTradeContextConverted) float64 {

	var groups []GroupsConverted
	//var pairs []PairsConverted

	groups = context.Collaterals[0].BorrowingFees.Groups
	//pairs = context.Collaterals[0].BorrowingFees.Pairs

	var firstPairGroup GroupsConverted
	if len(groups) > 0 {
		firstPairGroup = groups[0]
	}

	fee := 0.0
	if firstPairGroup.Block > initialAccFees.Block {
		if long {
			fee = firstPairGroup.PairAccFeeLong - initialAccFees.AccPairFee
		} else {
			fee = firstPairGroup.PairAccFeeShort - initialAccFees.AccPairFee
		}
	}

	for i := len(groups) - 1; i >= 0; i-- {
		deltaGroup, deltaPair, beforeTradeOpen := getPairGroupAccFeesDeltas(i, groups, initialAccFees, pairIndex, long, context)
		fee += math.Max(deltaGroup, deltaPair)
		if beforeTradeOpen {
			break
		}
	}

	return (posDai * fee) / 100
}

//func withinMaxGroupOi(pairIndex int, long bool, positionSizeCollateral float64, context altseasonfees.GainsTradeContextConverted) bool {
//	if len(context.Groups) == 0 || len(context.Pairs) == 0 {
//		return false
//	}
//
//	g := context.Groups[getPairGroupIndex(pairIndex, context)].Oi
//	if g.Max == 0 {
//		return true
//	}
//	if long {
//		return g.Long+positionSizeCollateral <= g.Max
//	} else {
//		return g.Short+positionSizeCollateral <= g.Max
//	}
//}

//func getPairGroupIndex(pairIndex int, context altseasonfees.GainsTradeContextConverted) int {
//	pairs := context.Groups
//	if pairIndex >= PairIndex(len(pairs)) {
//		return 0
//	}
//
//	pairGroups := context.Fees
//	if len(pairGroups) == 0 {
//		return 0
//	}
//	return pairGroups[0].GroupIndex
//}

func getPairPendingAccFees(pairIndex int, currentBlock uint64, context GainsTradeContextConverted) (float64, float64, float64) {
	pair := context.Collaterals[0].BorrowingFees.Pairs[pairIndex]
	return getPendingAccFees(pair.AccFeeLong, pair.AccFeeShort, pair.Oi.Long, pair.Oi.Short, pair.FeePerBlock, currentBlock, pair.AccLastUpdatedBlock, pair.Oi.Max, pair.FeeExponent)
}

func getPairPendingAccFee(pairIndex int, currentBlock uint64, long bool, context GainsTradeContextConverted) float64 {
	accFeeLong, accFeeShort, _ := getPairPendingAccFees(pairIndex, currentBlock, context)
	if long {
		return accFeeLong
	} else {
		return accFeeShort
	}
}

// TODO: Not sure if this is correct at all
func getGroupPendingAccFees(groupIndex int, currentBlock uint64, context GainsTradeContextConverted) (float64, float64, float64) {
	group := context.Collaterals[0].BorrowingFees.Pairs[groupIndex]
	return getPendingAccFees(group.AccFeeLong, group.AccFeeShort, group.Oi.Long, group.Oi.Short, group.FeePerBlock, currentBlock, group.AccLastUpdatedBlock, group.Oi.Max, group.FeeExponent)
}

func getGroupPendingAccFee(groupIndex float64, currentBlock uint64, long bool, context GainsTradeContextConverted) float64 {
	accFeeLong, accFeeShort, _ := getGroupPendingAccFees(int(groupIndex), currentBlock, context)
	if long {
		return accFeeLong
	} else {
		return accFeeShort
	}
}

func getPairGroupAccFeesDeltas(i int, pairGroups []GroupsConverted, initialFees InitialAccFeesConverted, pairIndex int, long bool, context GainsTradeContextConverted) (float64, float64, bool) {
	log.Println()
	group := pairGroups[i]
	beforeTradeOpen := group.Block < initialFees.Block

	var deltaGroup, deltaPair float64
	if i == len(pairGroups)-1 {
		deltaGroup = getGroupPendingAccFee(group.GroupIndex, context.CurrentBlock, long, context)
		deltaPair = getPairPendingAccFee(pairIndex, context.CurrentBlock, long, context)
	} else {
		nextGroup := pairGroups[i+1]
		if beforeTradeOpen && nextGroup.Block <= initialFees.Block {
			return 0, 0, beforeTradeOpen
		}

		if long {
			deltaGroup = nextGroup.PrevGroupAccFeeLong
			deltaPair = nextGroup.PairAccFeeLong
		} else {
			deltaGroup = nextGroup.PrevGroupAccFeeShort
			deltaPair = nextGroup.PairAccFeeShort
		}
	}

	if beforeTradeOpen {
		deltaGroup -= initialFees.AccGroupFee
		deltaPair -= initialFees.AccPairFee
	} else {
		if long {
			deltaGroup -= group.InitialAccFeeLong
			deltaPair -= group.PairAccFeeLong
		} else {
			deltaGroup -= group.InitialAccFeeShort
			deltaPair -= group.PairAccFeeShort
		}
	}

	return deltaGroup, deltaPair, beforeTradeOpen
}

func getPendingAccFees(accFeeLong, accFeeShort, oiLong, oiShort, feePerBlock float64, currentBlock, accLastUpdatedBlock uint64, maxOi, feeExponent float64) (float64, float64, float64) {
	moreShorts := oiLong < oiShort
	netOi := math.Abs(oiLong - oiShort)
	delta := 0.0
	if maxOi > 0 && feeExponent > 0 {
		delta = feePerBlock * float64(currentBlock-accLastUpdatedBlock) * math.Pow(netOi/maxOi, feeExponent)
	}

	newAccFeeLong := accFeeLong
	newAccFeeShort := accFeeShort
	if moreShorts {
		newAccFeeShort += delta
	} else {
		newAccFeeLong += delta
	}

	return newAccFeeLong, newAccFeeShort, delta
}

//func getBorrowingDataActiveFeePerBlock(val interface{}) float64 {
//	var oi OpenInterest
//	var feePerBlock, feeExponent float64
//
//	switch v := val.(type) {
//	case Pair:
//		oi = v.Oi
//		feePerBlock = v.FeePerBlock
//		feeExponent = v.FeeExponent
//	case Group:
//		oi = v.Oi
//		feePerBlock = v.FeePerBlock
//		feeExponent = v.FeeExponent
//	default:
//		return 0
//	}
//
//	netOi := math.Abs(oi.Long - oi.Short)
//	return feePerBlock * math.Pow(netOi/oi.Max, feeExponent)
//}

//func getActiveFeePerBlock(pair Pair, group *Group) float64 {
//	pairFeePerBlock := getBorrowingDataActiveFeePerBlock(pair)
//
//	if group == nil {
//		return pairFeePerBlock
//	}
//
//	groupFeePerBlock := getBorrowingDataActiveFeePerBlock(*group)
//	return math.Max(pairFeePerBlock, groupFeePerBlock)
//}
