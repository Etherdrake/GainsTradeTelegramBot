package bigtools

import (
	"fmt"
	"math/big"
)

// ConvertBigIntString converts a BigInt string to a normal string
func ScaleBigStrToFloatString(bigIntStr string, exponent uint) (string, error) {
	// Parse the base part as a big float
	base := new(big.Float)
	_, _, err := base.Parse(bigIntStr, 10)
	if err != nil {
		return "", fmt.Errorf("failed to parse base part as big float: %s", bigIntStr)
	}

	// Calculate the divisor for the percentage (10^exponent)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil))

	// Divide base by divisor to get the actual value
	base.Quo(base, divisor)

	// Convert the result back to a string
	result := base.Text('f', 3) // Round to 3 decimal places for percentage display

	return result, nil
}

func ScaleBigStrToFloat(bigIntStr string, exponent uint) (float64, error) {
	// Parse the base part as a big float
	base := new(big.Float)
	_, _, err := base.Parse(bigIntStr, 10)
	if err != nil {
		return 0.0, fmt.Errorf("failed to parse base part as big float: %s", bigIntStr)
	}

	// Calculate the divisor for scaling (10^exponent)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(exponent)), nil))

	// Divide base by divisor to get the scaled value
	base.Quo(base, divisor)

	// Convert the result to a float64
	result, _ := base.Float64()

	return result, nil
}
