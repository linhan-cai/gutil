package gpool

import (
	"fmt"
	"testing"
)

func TestNewGoroutinePool(t *testing.T) {
	pool := NewGoroutinePool(10, 20)

	for i := 0; i < 100; i++ {
		go func(i int) {
			pool.Add(func() {
				fmt.Println("i", i)
			})
		}(i)
	}
}
