package circuitBreaker

import (
	"sync/atomic"
)

type Counter interface {
	Add(i int64)
	Get() int64
	Zero()
}

type atomicCounter struct {
	x int64
}

func (c *atomicCounter) Add(i int64) {
	atomic.AddInt64(&c.x, i)
}

func (c *atomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.x)
}

func (c *atomicCounter) Zero() {
	atomic.StoreInt64(&c.x, 0)
}
