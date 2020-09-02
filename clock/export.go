package clock

import "time"

func Now() time.Time {
	current := gClock.t.Load().(*dial)
	return current.t
}

func DateTime() string {
	current := gClock.t.Load().(*dial)
	return current.DateTime
}

func Date() string {
	current := gClock.t.Load().(*dial)
	return current.Date
}

func Unix() int64 {
	current := gClock.t.Load().(*dial)
	return current.Unix
}