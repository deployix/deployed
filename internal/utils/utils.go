package utils

import "time"

func GetCurrentDateTimeAsString(format string) string {
	return time.Now().Format(format)
}
