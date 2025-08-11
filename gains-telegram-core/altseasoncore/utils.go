package altseasoncore

import (
	"fmt"
	"math/big"
)

// Helper function to convert a string to float64 by dividing by a factor
func parseAndDivide(input string, factor float64) (float64, error) {
	bigNum, ok := new(big.Float).SetString(input)
	if !ok {
		return 0, fmt.Errorf("invalid string: %s", input)
	}

	factorBig := big.NewFloat(factor)
	result := new(big.Float).Quo(bigNum, factorBig)

	resultFloat64, _ := result.Float64()
	return resultFloat64, nil
}
