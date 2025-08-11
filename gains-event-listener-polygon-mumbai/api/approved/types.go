package approved

type CheckAllowanceGNS struct {
	TraderAddress string `json:"trader_address"`
}

type ApproveDAI struct {
	Guid          int64  `json:"guid"`
	TraderAddress string `json:"trader_address"`
	Amount        string `json:"amount"`
}
