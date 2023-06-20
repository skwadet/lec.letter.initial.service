package utils

import "time"

func ParseTime(timeToParse string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.999999", timeToParse)
}

func InnerFormatTime(timeToFormat time.Time) string {
	return timeToFormat.Format("2006-01-02T15:04:05.999999")
}

func OuterFormatTime(timeToFormat time.Time) string {
	return timeToFormat.Format("2006-02-01 15:04:05")
}

func OuterFormatDate(timeToFormat time.Time) string {
	return timeToFormat.Format("02.01.2006")
}
