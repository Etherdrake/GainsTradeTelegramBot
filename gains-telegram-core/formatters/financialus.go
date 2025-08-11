package formatters

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// FormatFinancialUS graciously formats a given float number or string as a financial string in US format with the currency symbol ($)
func FormatFinancialUS(amount interface{}) (string, error) {
	var floatAmount float64
	var err error

	// Politely ascertain the type of the input and convert it to float64 if necessary
	switch v := amount.(type) {
	case float32:
		floatAmount = float64(v)
	case float64:
		floatAmount = v
	case string:
		// Attempt to convert string to float64
		floatAmount, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return "", fmt.Errorf("regrettably, the provided string cannot be converted to a float: %s", v)
		}
	default:
		return "", fmt.Errorf("regrettably, the provided type is unsupported: %s", reflect.TypeOf(amount))
	}

	// Format the number with commas as thousand separators and two decimal places
	formatted := fmt.Sprintf("%.2f", floatAmount)

	// Insert commas into the formatted string for better readability
	intPart, decPart := splitAmount(formatted)
	formattedIntPart := insertCommas(intPart)

	// Combine the integer and decimal parts with the distinguished currency symbol ($) before the number
	finalFormatted := "$" + formattedIntPart + "." + decPart

	// Place the dollar sign after the negative sign if the amount is negative
	if floatAmount < 0 {
		if floatAmount < 0 {
			if floatAmount > -1000 {
				finalFormatted = strings.Replace(finalFormatted, "$-,", "-$", 1)
			} else {
				finalFormatted = strings.Replace(finalFormatted, "$-", "-$", 1)
			}
		}
	}
	return finalFormatted, nil
}

// FormatFinancialPriceUS formats a given float number or string as a financial string in US format with the currency symbol ($)
func FormatFinancialPriceUS(amount interface{}) (string, error) {
	var floatAmount float64
	var err error

	// Politely ascertain the type of the input and convert it to float64 if necessary
	switch v := amount.(type) {
	case float32:
		floatAmount = float64(v)
	case float64:
		floatAmount = v
	case string:
		// Attempt to convert string to float64
		floatAmount, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return "", fmt.Errorf("regrettably, the provided string cannot be converted to a float: %s", v)
		}
	default:
		return "", fmt.Errorf("regrettably, the provided type is unsupported: %s", reflect.TypeOf(amount))
	}

	// Convert floatAmount to string to determine the length of the decimal part
	strAmount := strconv.FormatFloat(floatAmount, 'f', -1, 64)
	decimalIndex := strings.IndexByte(strAmount, '.')

	// Determine the precision based on the length of the decimal part
	precision := len(strAmount) - decimalIndex - 1

	// Format the number with commas as thousand separators and original decimal places
	formatString := fmt.Sprintf("$%%.%df", precision)
	formatted := fmt.Sprintf(formatString, floatAmount)

	// Insert commas into the formatted string for better readability
	intPart, decPart := splitAmount(formatted)
	formattedIntPart := insertCommas(intPart)

	// Combine the integer and decimal parts with the distinguished currency symbol ($) before the number
	finalFormatted := formattedIntPart + "." + decPart

	// Place the dollar sign after the negative sign if the amount is negative
	if floatAmount < 0 {
		if floatAmount < 0 {
			if floatAmount > -1000 {
				finalFormatted = strings.Replace(finalFormatted, "$-,", "-$", 1)
			} else {
				finalFormatted = strings.Replace(finalFormatted, "$-", "-$", 1)
			}
		}
	}
	return finalFormatted, nil
}

// splitAmount splits the formatted amount into integer and decimal parts
func splitAmount(amount string) (string, string) {
	parts := strings.Split(amount, ".")
	if len(parts) != 2 {
		return amount, "00" // Default case, should not happen with "%.2f"
	}
	return parts[0], parts[1]
}

// insertCommas inserts commas into the integer part of the amount
func insertCommas(amount string) string {
	n := len(amount)
	if n <= 3 {
		return amount
	}

	reversed := reverseString(amount)
	var builder strings.Builder
	for i, r := range reversed {
		if i > 0 && i%3 == 0 {
			builder.WriteRune(',')
		}
		builder.WriteRune(r)
	}
	return reverseString(builder.String())
}

// reverseString reverses the input string
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
