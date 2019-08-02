package util

import (
	"time"
)

//字符串转换为time.Time时间
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

//字符串转换为uninx时间
func DateStringToUinx(dateString string) int64 {
	return DateStringToTime(dateString).Unix()
}

//时间转换为字符串
func DateTimeToString(date time.Time) string {

	return date.Format("2006-01-02 15:04:05")
}

//unixTime 秒 转换为 time时间
func DateUnixTimeToTime(unixTime int64) time.Time {

	return time.Unix(unixTime, 0)
}

//unixTime秒转换为字符串时间
func DateUnixTimeToString(unixTime int64) string {

	return DateTimeToString(DateUnixTimeToTime(unixTime))
}
