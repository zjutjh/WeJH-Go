package circuitBreaker

import (
	"sync"
	"wejh-go/config/api/funnelApi"
	cbConfig "wejh-go/config/circuitBreaker"
)

var CB CircuitBreaker

type CircuitBreaker struct {
	LB       LoadBalance
	SnapShot sync.Map
}

func init() {
	lb := LoadBalance{
		zfLB:    &randomLB{},
		oauthLB: &randomLB{},
	}
	snapShot := sync.Map{}

	for _, api := range cbConfig.GetLoadBalanceConfig().Apis {
		lb.Add(api, funnelApi.Oauth)
		lb.Add(api, funnelApi.ZF)
		snapShot.Store(api+string(funnelApi.Oauth), NewApiSnapShot())
		snapShot.Store(api+string(funnelApi.ZF), NewApiSnapShot())
	}

	CB = CircuitBreaker{
		LB:       lb,
		SnapShot: snapShot,
	}
}

func (c *CircuitBreaker) GetApi(zfFlag, oauthFlag bool) (string, funnelApi.LoginType, error) {
	return c.LB.Pick(zfFlag, oauthFlag)
}

func (c *CircuitBreaker) Fail(api string, loginType funnelApi.LoginType) {
	value, _ := c.SnapShot.Load(api + string(loginType))
	if value.(*ApiSnapShot).Fail() {
		c.LB.Remove(api, loginType)
		Probe.Add(api, loginType)
	}
}

func (c *CircuitBreaker) Success(api string, loginType funnelApi.LoginType) {
	value, _ := c.SnapShot.Load(api + string(loginType))
	value.(*ApiSnapShot).Success()
}
