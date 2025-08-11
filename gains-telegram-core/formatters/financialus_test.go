package formatters

import "testing"

func TestFormatFinancialUS(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{0.4394033, "$0.4394033"},
		{0.2431, "$0.2431"},
		{0.0245, "$0.0245"},
		{-1234.5678, "-$1,234.5678"},
	}

	for _, tc := range testCases {
		result, err := FormatFinancialUS(tc.input)
		if err != nil {
			t.Errorf("Error formatting %v: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v for input %v", tc.expected, result, tc.input)
		}
	}
}
