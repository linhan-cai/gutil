package gutil

// 判断是否是同一天
func IsOneDay(timestamps ...int64) bool {
	if len(timestamps) == 0 {
		return true
	}
	day := getDay(timestamps[0])
	for i := 1; i < len(timestamps); i++ {
		if getDay(timestamps[i]) != day {
			return false
		}
	}
	return true
}

// 获取时间戳的天数
func getDay(timestamp int64) int64 {
	return timestamp / 86400
}

// 获取0点的时间戳
func MidNight(timestamp int64) int64 {
	return timestamp - timestamp%86400
}
