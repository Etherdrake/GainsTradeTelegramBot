package swaps

type UnizenTokenInfoResponse struct {
	TokenInfo TokenInfo `json:"tokenInfo"`
}

type TokenInfo struct {
	CirculatingSupply int64          `json:"circulatingSupply"`
	Description       string         `json:"description"`
	MarketCap         float64        `json:"marketCap"`
	MaxSupply         interface{}    `json:"maxSupply"`
	PriceInUsd        float64        `json:"priceInUsd"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
	TotalSupply       int64          `json:"totalSupply"`
	Volume24H         float64        `json:"volume24h"`
	PercentChange24H  float64        `json:"percentChange24h"`
	IsStableCoin      bool           `json:"isStableCoin"`
	Contracts         []ContractInfo `json:"contracts"`
}

type ContractInfo struct {
	Symbol          string `json:"symbol"`
	Blockchain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
	Decimals        int    `json:"decimals"`
	ChainID         int    `json:"chain_id"`
	Label           string `json:"label,omitempty"`
	PossibleSpam    bool   `json:"possible_spam,omitempty"`
}
