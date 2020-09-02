package clock

import (
	"sync/atomic"
	"time"
)

var gClock *Clock

const step = time.Millisecond * 20

func init() {
	gClock = &Clock{}
	gClock.fresh()
	go gClock.ticking()
}

type dial struct {
	t        time.Time
	DateTime string
	Date     string
	Unix 	 int64
}

type Clock struct {
	t atomic.Value
}

func (c *Clock) ticking() {
	// 首次时间差校准
	<- time.NewTimer(time.Duration(Now().UnixNano()) % step).C
	c.fresh()

	// ticker 校准
	timer := time.NewTicker(step)
	for _ = range timer.C {
		c.fresh()
	}
}

func (c *Clock) fresh() {
	now := time.Now()
	current := &dial{
		t: now,
		DateTime: now.Format("2006-01-02 15:04:05"),
		Date: now.Format("2006-01-02"),
		Unix: now.Unix(),
	}
	c.t.Store(current)
}