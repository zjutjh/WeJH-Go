package circuitBreaker

import (
	"time"
	"wejh-go/config/api/funnelApi"
)

type apiSnapShot struct {
	LoginType  funnelApi.LoginType
	State      State
	ErrCount   Counter
	TotalCount Counter
	AccessLast time.Time
}

func (a *apiSnapShot) Fail() bool {
	a.ErrCount.Add(1)
	a.TotalCount.Add(1)
	a.AccessLast = time.Now()
	if a.ErrCount.Get() > 10 {
		a.State = Open
		return true
	}
	return false
}

func (a *apiSnapShot) Success() {
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
		return "HALFOPEN"
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
