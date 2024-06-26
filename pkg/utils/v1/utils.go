package v1

import (
	"strconv"
	"time"
)

func GetCurrentDateTimeAsString(format string) string {
	return time.Now().Format(format)
}

// ConvertStringToType converts string value into the proper datatype
func ConvertStringToType(in string) interface{} {
	// check if string type is bool
	boolType, err := strconv.ParseBool(in)
	if err == nil {
		return boolType
	}

	// check if string type is int
	intType, err := strconv.Atoi(in)
	if err == nil {
		return intType
	}

	// default to string
	return in
}
