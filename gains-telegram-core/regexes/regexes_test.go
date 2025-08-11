package regexes

import "testing"

func TestLeveragePattern(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"10x32", "1032"},
		{"abc123def456", "123456"},
		{"noNumbers", ""},
	}

	for _, tc := range testCases {
		result := LeveragePattern.FindAllString(tc.input, -1)
		concatenatedNumber := ""
		if len(result) > 0 {
			concatenatedNumber = ""
			for _, num := range result {
				concatenatedNumber += num
			}
		}

		if concatenatedNumber != tc.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, concatenatedNumber)
		}
	}
}
