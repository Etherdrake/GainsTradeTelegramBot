// Open Trade Human Representation

user: 0xbc9b019aa5885ba50f3770b6fd7e5da981007068: address            // User address          | user: 0xbc9b019aa5885ba50f3770b6fd7e5da981007068: address            // User address
index: 0: uint32                                                     // Index                 | index: 2: uint32                                                     // Index
pair_index: 0: uint16                                                // Pair index            | pair_index: 0: uint16                                                // Pair index
leverage: 38.5: uint24                                               // Leverage              | leverage: 38500: uint24                                              // Leverage (as 38.5 * 1000 = 38500)
long: true: bool                                                     // Long position         | long: true: bool                                                     // Long position
is_open: false: bool                                                 // Is open               | is_open: true: bool                                                  // Is open
collateral_index: 1: uint8                                           // Collateral index -> DAI | collateral_index: 1: uint8                                           // Collateral index
trade_type: 0: uint8                                                 // Trade type -> LIMIT   | trade_type: 1: uint8                                                 // Trade type
collateral_amount: 500: uint64                                       // Collateral amount     | collateral_amount: 500000000000000000000: uint120                   // Collateral amount (as 500 * 10^18)
open_price: 128000: uint64                                           // Open price            | open_price: 509605900000000: uint64                                  // Open price
tp: 62873.5: uint64                                                  // Take profit           | tp: 628734551948051: uint64                                           // Take profit
sl: 50828.2: uint64                                                  // Stop loss             | sl: 508282248311688: uint120                                        // Stop loss
placeholder: 0: uint192                                              // Placeholder -> What even is this? | placeholder: 0: uint192                                              // Placeholder
max_slippage_p: 1040: uint16                                         // Max slippage percentage -> 2% | max_slippage_p: 2000: uint16                                         // Max slippage percentage
referrer: 0x0000000000000000000000000000000000000000: address        // Referrer address      | referrer: 0x0000000000000000000000000000000000000000: address        // Referrer address


// Golang Struct

// TradePayload represents the trade details
type TradePayload struct {
User             string `json:"user"`
Index            uint32 `json:"index"`
PairIndex        uint16 `json:"pair_index"`
Leverage         uint32 `json:"leverage"`
Long             bool   `json:"long"`
IsOpen           bool   `json:"is_open"`
CollateralIndex  uint8  `json:"collateral_index"`
TradeType        uint8  `json:"trade_type"`
CollateralAmount uint64 `json:"collateral_amount"`
OpenPrice        uint64 `json:"open_price"`
Tp               uint64 `json:"tp"`
Sl               uint64 `json:"sl"`
Placeholder      uint64 `json:"__placeholder"`
}

// OpenTradePayload represents the payload for opening a trade
type OpenTradePayload struct {
Guid            int64        `json:"guid"`
Trade           TradePayload `json:"trade"`
Chain           uint64       `json:"chain"`
CollateralToken string       `json:"collateral_token"`
}