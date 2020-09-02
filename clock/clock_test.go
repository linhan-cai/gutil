package clock

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	timer := time.NewTicker(time.Millisecond)
	var i = 0
	for _ = range timer.C {
		assert.EqualValues(t, time.Now().Unix(), Now().Unix())
		fmt.Println("aaa")
		i++
		if i > 100 {
			timer.Stop()
			break
		}
	}
}

// time    	2000000000	         1.32 ns/op	       0 B/op	       0 allocs/op
// unix		2000000000	         1.40 ns/op	       0 B/op	       0 allocs/op
// datetime 2000000000	         1.34 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNow1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DateTime()
	}
}


// time    	20000000	        66.8 ns/op	       0 B/op	       0 allocs/op
// unix		20000000	        67.3 ns/op	       0 B/op	       0 allocs/op
// datetime 5000000	       		270 ns/op	      32 B/op	       1 allocs/op
func BenchmarkNow2(b *testing.B) {
	for i:=0; i<b.N; i++ {
		time.Now().Format("2006-01-02 15:04:05")
	}
}
