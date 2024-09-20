package util

import "time"

const (
	timeFormat = "2006-01-02 15:04:05"
)

func FormatNow() string {
	now := time.Now()
	formattedTime := now.Format(timeFormat)
	return formattedTime
}
