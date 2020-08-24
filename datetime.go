package gutil

import "time"

// 判断是否是同一天
func IsOneDay(timestamps ...int64) bool {
	if len(timestamps) == 0 {
		return true
	}
	day := MiddleNight(timestamps[0])
	for i := 1; i < len(timestamps); i++ {
		if MiddleNight(timestamps[i]) != day {
			return false
		}
	}
	return true
}

// 获取0点的时间戳
func MiddleNight(timestamp int64) int64 {
	today := time.Unix(timestamp, 0)
	return time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location()).Unix()
}
