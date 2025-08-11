package stringbuilders

import (
	"HootCore/pairmaps"
	"HootCore/rdbhoot"
	"HootCore/utils"
	"errors"
	"fmt"
	"log"
	"strings"
)

func BuildCryptoPairString(rdbHoot *rdbhoot.HootRedisClient, page int) string {
	header := "📌 **All Crypto Pairs**\n\nSelect Asset To Trade (24hr)\n\n"
	var pairsList string

	pageArray := pairmaps.FilteredIndexToCryptoPage(page, 10)

	for _, pair := range pageArray {
		pairReturn := strings.Replace(pair, "/", "x", -1)
		name := pairmaps.SymbolToName[pairmaps.CryptoToIndex[pair]]
		idx := pairmaps.CryptoToIndex[pair]
		prctChange, rdbErr := rdbHoot.GetPrctChange(idx)
		if rdbErr != nil {
			log.Println(rdbErr)
		}

		price, rdbPriceErr := rdbHoot.GetPrice(idx)
		if rdbPriceErr != nil {
			log.Println(rdbPriceErr)
		}

		decLength, err := utils.GetDecimalLengthBasedOnIntegerLength(price)
		if err != nil {
			log.Println(err)
		}

		upDownDot, err := GetGreenRedDot(prctChange)
		if err != nil {
			log.Println("Error GetGreenRedDot: ", err)
		}

		formatString := fmt.Sprintf("/%%s: %%s *%%.%df* %%s *%%.2f%%%%*\n", decLength)
		result := fmt.Sprintf(formatString, pairReturn, name, price, upDownDot, prctChange)

		pairsList += result
	}
	return header + pairsList
}

func BuildFxPairString(rdbHoot *rdbhoot.HootRedisClient, page int) string {
	header := "📌 **All Forex Pairs**\n\nSelect Asset To Trade (24hr)\n\n"
	var pairsList string

	pageArray := pairmaps.FilteredIndexToFxPage(page, 10)

	for _, pair := range pageArray {
		pairReturn := strings.Replace(pair, "/", "x", -1)
		name := pairmaps.FxToFlag[pair]
		idx := pairmaps.FxToIndex[pair]
		prctChange, rdbErr := rdbHoot.GetPrctChange(idx)
		if rdbErr != nil {
			log.Println(rdbErr)
		}

		price, rdbPriceErr := rdbHoot.GetPrice(idx)
		if rdbPriceErr != nil {
			log.Println(rdbPriceErr)
		}

		upDownDot, err := GetGreenRedDot(prctChange)
		if err != nil {
			log.Println("Error GetGreenRedDot: ", err)
		}

		decLength, err := utils.GetDecimalLengthBasedOnIntegerLength(price)
		if err != nil {
			log.Println("Error decimalLength: ", err)
		}

		formatString := fmt.Sprintf("/%%s: %%s *%%.%df* %%s *%%.2f%%%%\n*", decLength)
		result := fmt.Sprintf(formatString, pairReturn, name, price, upDownDot, prctChange)

		pairsList += result
	}
	return header + pairsList
}

func BuildCommoditiesPairString(rdbHoot *rdbhoot.HootRedisClient) string {
	header := "📌 **All Commodity Pairs**\n\nSelect Asset To Trade (24hr)\n\n"
	var pairsList string

	// Get the Commodities array
	commoditiesArray := pairmaps.IndexToCommodity

	// Construct the list
	for _, pair := range commoditiesArray {
		pairReturn := strings.Replace(pair, "/", "x", -1)
		name := pairmaps.SymbolToName[pairmaps.CryptoToIndex[pair]]
		idx := pairmaps.CryptoToIndex[pair]
		prctChange, rdbErr := rdbHoot.GetPrctChange(idx)
		if rdbErr != nil {
			log.Println(rdbErr)
		}

		price, rdbPriceErr := rdbHoot.GetPrice(idx)
		if rdbPriceErr != nil {
			log.Println(rdbPriceErr)
		}

		decLength, err := utils.GetDecimalLengthBasedOnIntegerLength(price)
		if err != nil {
			log.Println(err)
		}

		upDownDot, err := GetGreenRedDot(prctChange)
		if err != nil {
			log.Println("Error GetGreenRedDot: ", err)
		}

		formatString := fmt.Sprintf("/%%s: %%s *%%.%df* %%s *%%.2f%%%%*\n", decLength)
		result := fmt.Sprintf(formatString, pairReturn, name, price, upDownDot, prctChange)

		pairsList += result
	}
	return header + pairsList
}

func GetGreenRedDot(price interface{}) (string, error) {
	greenDot := "🟢"
	redDot := "🔴"

	switch v := price.(type) {
	case float64:
		switch v < 0 {
		case true:
			return redDot, nil
		case false:
			return greenDot, nil
		}
	case int64:
		switch v < 0 {
		case true:
			return redDot, nil
		case false:
			return greenDot, nil
		}
	case uint:
		return greenDot, nil
	default:
		return "", errors.New("input is not a numeric type")
	}
	return "", errors.New("input is not a numeric type")
}
