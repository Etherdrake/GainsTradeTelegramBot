package altseasoncore

func GetSpreadWithPriceImpactP(
	pairSpreadP float64,
	buy bool,
	collateral float64,
	leverage float64,
	pairDepth PairDepthConverted,
	oiWindowsSettings OIWindowsSettingsConverted,
	oiWindows OIWindowConverted,
) float64 {
	if pairSpreadP == 0 {
		return 0
	}

	var onePercentDepth uint64
	if buy {
		onePercentDepth = pairDepth.OnePercentDepthAboveUsd
	} else {
		onePercentDepth = pairDepth.OnePercentDepthBelowUsd
	}

	var activeOi float64 = -1

	if oiWindowsSettings.WindowsCount > 0 {
		activeOi = getActiveOi(
			getCurrentOiWindowId(oiWindowsSettings),
			oiWindowsSettings.WindowsCount,
			oiWindows,
			buy,
		)
	}

	if onePercentDepth == 0 || activeOi == -1 || collateral == 0 {
		return pairSpreadP
	}

	onePercentDepthConv := float64(onePercentDepth)

	return pairSpreadP + (activeOi+(collateral*leverage)/2)/onePercentDepthConv/100
}
