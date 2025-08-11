package native

type ApprovePayload struct {
	Guid       int64  `json:"guid"`
	Owner      string `json:"owner"`
	Token      string `json:"token"`
	Spender    string `json:"spender"`
	Amount     string `json:"amount"`
	Collateral uint8  `json:"collateral"`
	Chain      uint64 `json:"chain"`
}

type BalanceOfPayload struct {
	PublicKey  string `json:"public_key"`
	Collateral uint8  `json:"collateral"`
	Chain      uint64 `json:"chain"`
}

type GetGasBalancePayload struct {
	PublicKey  string `json:"public_key"`
	Collateral uint8  `json:"collateral"`
	Chain      uint64 `json:"chain"`
}

type GetAllowancePayload struct {
	Owner      string `json:"owner"`
	Token      string `json:"token"`
	Spender    string `json:"spender"`
	Collateral uint8  `json:"collateral"`
	Chain      uint64 `json:"chain"`
}

type GetBlockTimePayload struct {
	BlockNumber uint64 `json:"block_number"`
	Chain       uint64 `json:"chain"`
}

type HootCoreResponseOK struct {
	Ok string
}

type HootCoreResponseErr struct {
	Err string
}
