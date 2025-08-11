package benchmark

import (
	"HootCore/tradecache"
	"log"
	"sync"
	"testing"
)

func BenchmarkMapOperations(b *testing.B) {
	// Create a TradeCache instance with a mutex
	tc := tradecache.New()
	var mu sync.Mutex

	// Perform initial setup if needed

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Perform map operations here
			// For example, insert and retrieve a trade from the map

			// Insert operation
			var guid int64
			guid = 947102709

			// Retrieve operation
			mu.Lock()
			_, exists := tc.Cache[guid]
			log.Println(exists)
			mu.Unlock()

			// Additional map operations
			mu.Lock()
			_ = tc.GenerateOrderID(guid)
			_ = tc.SetOrderTypeToMarket(guid)
			_ = tc.SetOrderTypeToLimit(guid)
			_ = tc.SetOrderTypeToStop(guid)
			_ = tc.IncrementPage(guid, "someClass")
			_ = tc.ResetPage(guid)
			_ = tc.DecrementPage(guid)
			_ = tc.SetPairIndex(123, guid)
			_ = tc.DecrementPairIndex(guid)
			_ = tc.IncrementPairIndex(guid)
			_ = tc.ToggleLongShort(guid)
			_ = tc.ToggleRealPaper(guid)
			mu.Unlock()

			// Add more map operations as needed
		}
	})
}
