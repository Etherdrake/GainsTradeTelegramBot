package utils

import (
	"time"
)

func BlockNumberToTime(blockNumber int64, averageBlockTime time.Duration) time.Time {
	// Calculate the time of the block based on the block number and average block time
	blockTime := time.Now().Add(time.Duration(-blockNumber) * averageBlockTime)

	return blockTime
}
