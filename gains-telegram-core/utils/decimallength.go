package utils

import (
	"fmt"
	"strings"
)

func GetDecimalLength(num interface{}) (int, error) {
	switch num.(type) {
	case float32, float64:
		strNum := fmt.Sprintf("%f", num)
		decimalPart := strings.Split(strNum, ".")[1]
		return len(decimalPart), nil
	default:
		return 0, fmt.Errorf("Input is not a float")
	}
}

func GetDecimalLengthBasedOnIntegerLength(num interface{}) (int, error) {
	integerLength := 0

	switch num.(type) {
	case float32, float64:
		strNum := fmt.Sprintf("%f", num)
		integerPart := strings.Split(strNum, ".")[0]
		integerLength = len(integerPart)
	default:
		return 0, fmt.Errorf("Input is not a float")
	}

	switch integerLength {
	case 5:
		return 2, nil
	case 4:
		return 3, nil
	case 3:
		return 4, nil
	case 2:
		return 5, nil
	case 1:
		return 6, nil
	default:
		return 7, nil
	}
}
