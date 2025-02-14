package circuitBreaker

import (
	"time"
)

type ApiSnapShot struct {
	State      State
	ErrCount   Counter
	TotalCount Counter
	AccessLast time.Time
}

func NewApiSnapShot() *ApiSnapShot {
	return &ApiSnapShot{
		State:      Closed,
		ErrCount:   &atomicCounter{},
		TotalCount: &atomicCounter{},
		AccessLast: time.Time{},
	}
}

func (a *ApiSnapShot) Fail() bool {
	a.ErrCount.Add(1)
	a.TotalCount.Add(1)
	if a.ErrCount.Get() > 50 && a.State == Closed {
		a.State = Open
		return true
	}
	return false
}

func (a *ApiSnapShot) Success() {
	a.ErrCount.Zero()
	a.State = Closed
	a.AccessLast = time.Now()
}

type State int32

func (s State) String() string {
	switch s {
	case Open:
		return "OPEN"
	case HalfOpen:
		return "HALF_OPEN"
	case Closed:
		return "CLOSED"
	}
	return "INVALID"
}

const (
	Open State = iota
	HalfOpen
	Closed
)
