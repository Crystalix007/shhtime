package formats

import (
	"errors"
	"fmt"
	"time"
)

// ErrInvalidTimeFormat is returned when a time format is invalid.
var ErrInvalidTimeFormat = errors.New("invalid time format")

// PredefinedTimeFormat is a type that represents a predefined time format.
type PredefinedTimeFormat string

const (
	FormatANSIC       PredefinedTimeFormat = time.ANSIC
	FormatUnixDate    PredefinedTimeFormat = time.UnixDate
	FormatRubyDate    PredefinedTimeFormat = time.RubyDate
	FormatRFC822      PredefinedTimeFormat = time.RFC822
	FormatRFC822Z     PredefinedTimeFormat = time.RFC822Z
	FormatRFC850      PredefinedTimeFormat = time.RFC850
	FormatRFC1123     PredefinedTimeFormat = time.RFC1123
	FormatRFC1123Z    PredefinedTimeFormat = time.RFC1123Z
	FormatRFC3339     PredefinedTimeFormat = time.RFC3339
	FormatRFC3339Nano PredefinedTimeFormat = time.RFC3339Nano
	FormatKitchen     PredefinedTimeFormat = time.Kitchen
	FormatStamp       PredefinedTimeFormat = time.Stamp
	FormatStampMilli  PredefinedTimeFormat = time.StampMilli
	FormatStampMicro  PredefinedTimeFormat = time.StampMicro
	FormatStampNano   PredefinedTimeFormat = time.StampNano
	FormatDateTime    PredefinedTimeFormat = time.DateTime
	FormatDateOnly    PredefinedTimeFormat = time.DateOnly
	FormatTimeOnly    PredefinedTimeFormat = time.TimeOnly
)

// FormatFromSnakeCase converts a snake_case string to a PredefinedTimeFormat.
func FormatFromSnakeCase(snake string) (PredefinedTimeFormat, error) {
	switch snake {
	case "ansic":
		return FormatANSIC, nil
	case "unix_date":
		return FormatUnixDate, nil
	case "ruby_date":
		return FormatRubyDate, nil
	case "rfc822":
		return FormatRFC822, nil
	case "rfc822z":
		return FormatRFC822Z, nil
	case "rfc850":
		return FormatRFC850, nil
	case "rfc1123":
		return FormatRFC1123, nil
	case "rfc1123z":
		return FormatRFC1123Z, nil
	case "rfc3339":
		return FormatRFC3339, nil
	case "rfc3339_nano":
		return FormatRFC3339Nano, nil
	case "kitchen":
		return FormatKitchen, nil
	case "stamp":
		return FormatStamp, nil
	case "stamp_milli":
		return FormatStampMilli, nil
	case "stamp_micro":
		return FormatStampMicro, nil
	case "stamp_nano":
		return FormatStampNano, nil
	case "datetime":
		return FormatDateTime, nil
	case "date_only":
		return FormatDateOnly, nil
	case "time_only":
		return FormatTimeOnly, nil
	}

	return "", fmt.Errorf("%w: %s", ErrInvalidTimeFormat, snake)
}
