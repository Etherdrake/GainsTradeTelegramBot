package swaps

import (
	"HootCore/rdbhoot"
	"context"
	"encoding/json"
	"fmt"
)

// SwapInit function to insert SwapInfos struct into Redis with hardcoded values for Ethereum, Arbitrum, and Polygon chains
func SwapInit(ctx context.Context, rdb *rdbhoot.HootRedisClient, guid string) error {
	// Create SwapInfos struct with hardcoded values for Ethereum, Arbitrum, and Polygon chains
	swapInfos := SwapInfos{
		Chains: map[uint64]SwapInfo{
			3: {
				SrcSymbol:        "ETH",
				SrcToken:         "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
				SrcTokenDecimals: "18",
				SrcTokenAmt:      "0",
				DstToken:         "0x6b175474e89094c44da98b954eedeac495271d0f",
				DstSymbol:        "DAI",
				DstTokenDecimals: "18",
			},
			1: {
				SrcSymbol:        "ETH",
				SrcToken:         "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
				SrcTokenDecimals: "18",
				SrcTokenAmt:      "0",
				DstToken:         "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1",
				DstSymbol:        "DAI",
				DstTokenDecimals: "18",
			},
			0: {
				SrcSymbol:        "MATIC",
				SrcToken:         "0x0000000000000000000000000000000000001010",
				SrcTokenDecimals: "18",
				SrcTokenAmt:      "0",
				DstToken:         "0x82e64f49ed5ec1bc6e43dad4fc8af9bb3a2312ee",
				DstSymbol:        "DAI",
				DstTokenDecimals: "18",
			},
		},
	}

	// Marshal SwapInfos struct to JSON
	swapInfosJSON, err := json.Marshal(swapInfos)
	if err != nil {
		return fmt.Errorf("error marshaling SwapInfos struct to JSON: %v", err)
	}

	// Write SwapInfos JSON to Redis using RedisJSON
	if redisErr := rdb.JSONSet(ctx, guid, ":swap", swapInfosJSON).Err(); redisErr != nil {
		return fmt.Errorf("error writing SwapInfos JSON to Redis: %v", redisErr)
	}
	return nil
}
