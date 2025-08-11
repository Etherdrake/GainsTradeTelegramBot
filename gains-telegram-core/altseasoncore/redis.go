package altseasoncore

//// GetGainsTradeContext retrieves GainsTradeContext from Redis
//func GetGainsTradeContext(rdbHoot *hootrdb.HootRedisClient, chain uint32) (GainsTradeContext, error) {
//	ctx := context.Background()
//
//	var chainStr string
//	switch chain {
//	case 0:
//		break
//	case 137:
//		chainStr = "polygon"
//	case 42161:
//		chainStr = "arbitrum"
//	case 421614:
//		chainStr = "sepolia"
//	default:
//		return GainsTradeContext{}, errors.New("unknown chain")
//
//	}
//
//	// Construct the key for Redis
//	key := "fees:" + chainStr
//
//	// Retrieve JSON data from Redis
//	jsonData, err := rdbHoot.JSONGet(ctx, key).Result()
//	if err != nil {
//		return GainsTradeContext{}, fmt.Errorf("failed to retrieve config from Redis: %v", err)
//	}
//
//	// Unmarshal JSON data into GainsTradeContext struct
//	var config GainsTradeContext
//	err = json.Unmarshal([]byte(jsonData), &config)
//	if err != nil {
//		return GainsTradeContext{}, fmt.Errorf("failed to unmarshal JSON data: %v", err)
//	}
//
//	return config, nil
//}
//
//// GetGainsTradeContext retrieves GainsTradeContext from Redis
//func GetBorrowingFeeContext(rdb *redis.Client, chain uint32) (GainsTradeContext, error) {
//	ctx := context.Background()
//
//	var chainStr string
//	switch chain {
//	case 0:
//		break
//	case 137:
//		chainStr = "polygon"
//	case 42161:
//		chainStr = "arbitrum"
//	case 421614:
//		chainStr = "sepolia"
//	default:
//		return GainsTradeContext{}, errors.New("unknown chain")
//
//	}
//
//	// Construct the key for Redis
//	key := "fees:" + chainStr
//
//	// Retrieve JSON data from Redis
//	jsonData, err := rdb.JSONGet(ctx, key).Result()
//	if err != nil {
//		return GainsTradeContext{}, fmt.Errorf("failed to retrieve config from Redis: %v", err)
//	}
//
//	// Unmarshal JSON data into GainsTradeContext struct
//	var config GainsTradeContext
//	err = json.Unmarshal([]byte(jsonData), &config)
//	if err != nil {
//		return GainsTradeContext{}, fmt.Errorf("failed to unmarshal JSON data: %v", err)
//	}
//
//	return config, nil
//}
