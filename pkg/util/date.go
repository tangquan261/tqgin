package util

import (
	"time"
)

func DateStringToTime(dateString string) time.Time {

	if len(dateString) <= 0 {
		return time.Now()
	}

	retDate, err := time.Parse("2006-01-02 15:04:05", dateString)

	if err != nil {
		return time.Now()
	}
	return retDate
}

func DateStringToUinx(dateString string) int64 {
	return DateStringToTime(dateString).Unix()
}
