package altseasoncore

import (
	"math"
	"strconv"
	"time"
)

func getCurrentOiWindowId(oiWindowSettings OIWindowsSettingsConverted) int {
	currentTs := time.Now().Unix()
	return int(math.Floor(float64(int(currentTs)-oiWindowSettings.StartTs) / float64(oiWindowSettings.WindowsDuration)))
}

func getActiveOi(currentOiWindowId int, windowsCount int, oiWindows OIWindowConverted, buy bool) float64 {
	if windowsCount == 0 {
		return 0
	}

	activeOi := 0.0
	for id := currentOiWindowId - (windowsCount - 1); id <= currentOiWindowId; id++ {
		window, exists := (oiWindows.Window)[strconv.Itoa(id)]
		if !exists {
			continue
		}
		if buy {
			activeOi += window.OILongUsd
		} else {
			activeOi += window.OIShortUsd
		}
	}

	return activeOi
}
