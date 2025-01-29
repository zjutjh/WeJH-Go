package circuitBreaker

import (
	"wejh-go/config/api/funnelApi"
	cbConfig "wejh-go/config/circuitBreaker"
)

var CB CircuitBreaker

type CircuitBreaker struct {
	LB       LoadBalance
	SnapShot map[string]*ApiSnapShot
}

func init() {
	lb := LoadBalance{
		zfLB:    &randomLB{},
		oauthLB: &randomLB{},
	}
	snapShot := make(map[string]*ApiSnapShot)

	for _, api := range cbConfig.GetLoadBalanceConfig().Apis {
		lb.Add(api, funnelApi.Oauth)
		lb.Add(api, funnelApi.ZF)
		snapShot[api+string(funnelApi.Oauth)] = NewApiSnapShot()
		snapShot[api+string(funnelApi.ZF)] = NewApiSnapShot()
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
	if c.SnapShot[api+string(loginType)].Fail() {
		c.LB.Remove(api, loginType)
		Probe.Add(api, loginType)
	}
}

func (c *CircuitBreaker) Success(api string, loginType funnelApi.LoginType) {
	c.SnapShot[api+string(loginType)].Success()
}
