package gpool

import "sync"

type GoroutinePool struct {
	taskChan  chan func()
	cap       int
	wg        *sync.WaitGroup
	closed    bool
	closeOnce *sync.Once
}

func NewGoroutinePool(cap int, chanSize int32) *GoroutinePool {
	p := &GoroutinePool{
		taskChan: make(chan func(), chanSize),
		cap:      cap,
	}

	for i := 0; i < cap; i++ {
		go func() {
			for handle := range p.taskChan {
				handle()
			}
		}()
	}
	return p
}

func (g *GoroutinePool) Add(handle func()) {
	g.taskChan <- handle
}
