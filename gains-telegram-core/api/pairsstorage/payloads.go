package pairsstorage

type GetPairMaxLeveragePayload struct {
	Index uint64 `json:"index"`
	Chain uint64 `json:"chain"`
}

type GetPairMinLeveragePayload struct {
	Index uint64 `json:"index"`
	Chain uint64 `json:"chain"`
}

type GetPairMinPosSizeUsdPayload struct {
	Index uint64 `json:"index"`
	Chain uint64 `json:"chain"`
}
