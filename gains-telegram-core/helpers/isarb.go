package helpers

import "github.com/pkg/errors"

func GetIsArb(chain string) (bool, error) {
	switch chain {
	case "polygon":
		return false, nil
	case "arbitrum":
		return true, nil
	case "":
		return false, errors.New("")
	}
	return false, errors.New("invalid chain string")
}
