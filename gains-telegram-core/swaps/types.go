package swaps

var GasTokenAddress = map[uint64]string{
	3: "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
	0: "0x0000000000000000000000000000000000001010",
	1: "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
}

type SwapInfos struct {
	Chains map[uint64]SwapInfo `json:"chains"`
}

// SwapInfo represents information about a swap
type SwapInfo struct {
	SrcSymbol        string `json:"srcTicker"`
	SrcToken         string `json:"srcToken"`
	SrcTokenDecimals string `json:"srcTokenDecimals"`
	SrcTokenAmt      string `json:"srcTokenAmt"`
	DstSymbol        string `json:"dstTicker"`
	DstToken         string `json:"dstToken"`
	DstTokenDecimals string `json:"dstTokenDecimals"`
}
