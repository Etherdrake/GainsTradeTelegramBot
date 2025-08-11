package balances

type BalanceGasTokenPayload struct {
	Chain         string `json:"chain"`
	TraderAddress string `json:"trader_address"`
}
