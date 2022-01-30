package test

import "time"

// MustParseTimeFromFormat used for tests only, panics if it fails to parse the time.
func MustParseTimeFromFormat(format string, timeValue string) time.Time {
	t, err := time.Parse(format, timeValue)
	if err != nil {
		panic(err)
	}

	return t
}
