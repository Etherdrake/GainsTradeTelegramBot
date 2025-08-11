package pairmaps

type MultiChainTokenInfo struct {
	EthereumAddress string
	PolygonAddress  string
	ArbitrumAddress string
}

type TokenInfo struct {
	Ticker  string
	Address string
}

var GasToken = map[string]string{
	"polygon":  "0x0000000000000000000000000000000000001010",
	"arbitrum": "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
}

var MultiChainTokenInfoUSDC = MultiChainTokenInfo{
	EthereumAddress: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	PolygonAddress:  "0x3c499c542cef5e3811e1192ce70d8cc03d5c3359",
	ArbitrumAddress: "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
}

var MultiChainTokenInfoUSDT = MultiChainTokenInfo{
	EthereumAddress: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	PolygonAddress:  "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
	ArbitrumAddress: "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
}

var MultiChainTokenInfoWETH = MultiChainTokenInfo{
	EthereumAddress: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	PolygonAddress:  "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
	ArbitrumAddress: "0x2f5e87c9312fa29aed5c179e456625d79015299c",
}

var MultiChainTokenInfoWBTC = MultiChainTokenInfo{
	EthereumAddress: "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
	PolygonAddress:  "0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6",
	ArbitrumAddress: "0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f",
}

var StarterSwapMap = map[string]MultiChainTokenInfo{
	"USDC": MultiChainTokenInfoUSDC,
	"USDT": MultiChainTokenInfoUSDT,
	"WETH": MultiChainTokenInfoWETH,
}

var SwapStarterPolygon = []TokenInfo{
	{
		Ticker:  "MATIC",
		Address: "0x0000000000000000000000000000000000001010",
	},
	{
		Ticker:  "DAI",
		Address: "0x8f3cf7ad23cd3cadbd9735aff958023239c6a063",
	},
	{
		Ticker:  "USDC",
		Address: "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
	},
	{
		Ticker:  "USDT",
		Address: "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
	},
	{
		Ticker:  "WETH",
		Address: "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
	},
}

var SwapStarterArbitrum = []TokenInfo{
	{
		Ticker:  "WETH",
		Address: "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
	},
	{
		Ticker:  "DAI",
		Address: "0xda10009cbd5d07dd0cecc66161fc93d7c9000da1",
	},
	{
		Ticker:  "USDC",
		Address: "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
	},
	{
		Ticker:  "USDT",
		Address: "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
	},
}
