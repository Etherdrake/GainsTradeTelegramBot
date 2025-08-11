package bigtools

import "math/big"

func ConvertWeiToFloat(amountWei string, decimals int) (float64, bool) {
	power := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	decimalBig := new(big.Float).SetInt(power)
	amountWeiBig, success := new(big.Float).SetString(amountWei)
	if !success {
		return 0, false
	}
	amountWeiBig.Quo(amountWeiBig, decimalBig)
	amountWeiBigFloat64, _ := amountWeiBig.Float64()
	return amountWeiBigFloat64, true
}
