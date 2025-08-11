package listenerutils

import "errors"

func ExtractPubkey(inputString string) (string, error) {
	if len(inputString) <= 42 {
		return "", errors.New("input too short: " + inputString)
	}
	return inputString[:42], nil
}
