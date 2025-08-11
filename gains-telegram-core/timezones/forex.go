package timezones

import "time"

var forexMarketsTimeZone = "America/New_York"

func IsForexOpen(dateToCheck time.Time) bool {
	loc, err := time.LoadLocation(forexMarketsTimeZone)
	if err != nil {
		// Handle error
		return false
	}

	now := dateToCheck.In(loc)
	weekday := now.Weekday()
	hour := now.Hour()
	dayOfMonth := now.Day()
	month := now.Month()
	isInDST := IsInDst(now)

	isClosed :=
		// Christmas 2023
		(month == 12 && dayOfMonth >= 25 && dayOfMonth <= 27) ||
			// New Year's Eve 2023
			(month == 1 && dayOfMonth >= 1 && dayOfMonth <= 2) ||
			// Friday after 4PM (DST) and 5PM (non-DST)
			(weekday == time.Friday && ((isInDST && hour >= 16) || hour >= 17)) ||
			// Saturday
			weekday == time.Saturday ||
			// Sunday before 4PM (DST) and 5PM (non-DST)
			(weekday == time.Sunday && ((isInDST && hour < 16) || hour < 17))

	return !isClosed
}

func IsInDst(t time.Time) bool {
	// Determine the year of the given time
	year, _, _ := t.Date()

	// Second Sunday in March
	start := time.Date(year, time.March, 1, 0, 0, 0, 0, t.Location())
	start = start.AddDate(0, 0, 14-int(start.Weekday()))
	if start.Day() > 7 {
		start = start.AddDate(0, 0, 7)
	}

	// First Sunday in November
	end := time.Date(year, time.November, 1, 0, 0, 0, 0, t.Location())
	end = end.AddDate(0, 0, 7-int(end.Weekday()))

	// Check if the given time falls within the DST period
	return t.After(start) && t.Before(end)
}
