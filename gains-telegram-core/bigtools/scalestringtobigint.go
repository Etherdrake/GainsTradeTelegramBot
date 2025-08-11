package bigtools

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func ConvertStringToBigIntAsString(input string) (string, error) {
	// Split the input string into integer and decimal parts
	parts := strings.Split(input, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format")
	}

	intPart, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to convert integer part: %v", err)
	}

	decimalPart := parts[1]

	// Count the number of decimal places to determine the scale
	scale := len(decimalPart)

	// Convert decimal part to the desired format (multiply by 10^scale)
	decimalBigInt, success := new(big.Int).SetString(decimalPart, 10)
	if !success {
		return "", fmt.Errorf("failed to convert decimal part to big.Int")
	}

	scaleBigInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(scale)), nil)
	decimalBigInt.Mul(decimalBigInt, scaleBigInt)

	// Add the integer and decimal parts
	result := new(big.Int)
	result.SetInt64(int64(intPart))
	result.Add(result, decimalBigInt)

	// Convert the result to a string
	resultStr := result.String()

	return resultStr, nil
}
