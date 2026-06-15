package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	if parsed, err := time.Parse("1/2/2006 15:04:05", date); err == nil {
		return parsed
	}
	if parsed, err := time.Parse("January 2, 2006 15:04:05", date); err == nil {
		return parsed
	}
	if parsed, err := time.Parse("Monday, January 2, 2006 15:04:05", date); err == nil {
		return parsed
	}
	panic("can't parse: " + date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	schedule := Schedule(date)
	return schedule.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	schedule := Schedule(date)
	hour, _, _ := schedule.Clock()
	return 12 <= hour && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	schedule := Schedule(date)
	return schedule.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, now.Location())
}
