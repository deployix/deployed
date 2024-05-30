package v1

import "time"

type DateTimeFormatType int

// valid datetime formats
const (
	ANSIC DateTimeFormatType = iota
	UnixDate
	RubyDate
	RFC822
	RFC822Z
	RFC850
	RFC1123
	RFC1123Z
	RFC3339
	RFC3339Nano
)

var (
	// map of time format that stores the layout
	TimeLayout = map[string]string{
		"ansic":       time.ANSIC,
		"unixdate":    time.UnixDate,
		"rubydate":    time.RubyDate,
		"rfc822":      time.RFC822,
		"rfc822z":     time.RFC822Z,
		"rfc850":      time.RFC850,
		"rfc1123":     time.RFC1123,
		"rfc1123z":    time.RFC1123Z,
		"rfc3339":     time.RFC3339,
		"rfc3339nano": time.RFC3339Nano,
	}
)

func (d DateTimeFormatType) String() string {
	return [...]string{time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano}[d]
}

func (d DateTimeFormatType) EnumIndex() int {
	return int(d)
}

func DateTimeLayoutFromTypeName(name string) string {
	return TimeLayout[name]
}
