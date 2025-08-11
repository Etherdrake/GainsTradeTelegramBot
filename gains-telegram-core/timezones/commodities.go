package timezones

import "time"

var commoditiesMarketsTimeZone = "America/New_York"

func IsCommoditiesOpen(dateToCheck time.Time) bool {
	loc, err := time.LoadLocation(commoditiesMarketsTimeZone)
	if err != nil {
		// handle error if timezone loading fails
		return false
	}
	now := dateToCheck.In(loc)

	weekday := now.Weekday()
	hour := now.Hour()
	dayOfMonth := now.Day()
	month := int(now.Month())
	//minute := now.Minute()

	isClosed :=
		// Christmas 2023
		(month == 12 && dayOfMonth >= 25 && dayOfMonth <= 27) ||
			// New Year's Eve 2023
			(month == 1 && dayOfMonth >= 1 && dayOfMonth <= 2) ||
			// Friday Closing
			(weekday == time.Friday && hour >= 17) ||
			// Saturday Closed
			(weekday == time.Saturday) ||
			// Saturday Opening
			(weekday == time.Sunday && hour <= 18) ||
			// Daily Closing
			(hour == 17)

	return !isClosed
}
