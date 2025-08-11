package tradfin_utils

import "time"

// We use CheckForexOpen also for the check on commodities
func CheckForexOpen(currentTime time.Time) bool {
	// Convert current time to New York time zone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	now := currentTime.In(loc)

	// Get day of the week (1 = Monday, 7 = Sunday)
	weekday := now.Weekday()

	// Get hour of the day
	hour := now.Hour()

	// Check if it's a holiday or weekend
	isHolidayOrWeekend :=
		// Christmas 2024
		(now.Month() == 12 && now.Day() >= 25 && now.Day() <= 27) ||
			// New Year's Eve 2024
			(now.Month() == 1 && now.Day() >= 1 && now.Day() <= 2) ||
			// Add checks for other holidays...

			// Saturday or Sunday
			weekday == time.Saturday || weekday == time.Sunday

	// Check if it's after market closing time on Friday (DST or non-DST)
	isAfterFridayMarketClose := weekday == time.Friday && ((hour >= 16 && now.IsDST()) || hour >= 17)

	// Check if it's before market opening time on Monday (DST or non-DST)
	isBeforeMondayMarketOpen := weekday == time.Monday && ((hour < 9 && now.IsDST()) || hour < 10)

	// Combine all conditions to check if the market is closed
	isClosed := isHolidayOrWeekend || isAfterFridayMarketClose || isBeforeMondayMarketOpen

	return !isClosed
}
