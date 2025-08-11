package formatters

import (
	"fmt"
	"reflect"
)

// FormatFinancialUK graciously formats a given float number as a financial string in UK format with the currency symbol (£)
func FormatFinancialUK(amount interface{}) (string, error) {
	var floatAmount float64

	// Politely ascertain the type of the input and convert it to float64 if necessary
	switch v := amount.(type) {
	case float32:
		floatAmount = float64(v)
	case float64:
		floatAmount = v
	default:
		return "", fmt.Errorf("regrettably, the provided type is unsupported: %s", reflect.TypeOf(amount))
	}

	// Elegantly format the number with commas as thousand separators and two decimal places
	formatted := fmt.Sprintf("%.2f", floatAmount)

	// Affix the distinguished currency symbol (£) before the number
	return "£" + formatted, nil
}
