package date_utils

import "time"

func GetIsWeekendOrWeekeDay(date *time.Time) bool {
	dayOfWeek := int(date.Weekday())
	if dayOfWeek == 0 || dayOfWeek == 6 {
		return true 
	}
	return false
}