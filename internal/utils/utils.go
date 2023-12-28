package utils

import "time"

func GetCurrentDateTimeAsString(format string) string {
	current_time := time.Now()
	return current_time.Format(format)
}
