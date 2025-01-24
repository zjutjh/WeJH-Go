package circuitBreaker

import (
	liveNessConfig "wejh-go/config/LiveNess"
	"wejh-go/config/api/funnelApi"
)

var CB CircuitBreaker

func init() {
	lb := LoadBalance{
		zfLB:    &randomLB{},
		oauthLB: &randomLB{},
	}
	for _, config := range liveNessConfig.GetLoadBalanceConfig() {
		if config.Type == funnelApi.Oauth {
			lb.oauthLB.Add(config.Url)
		} else if config.Type == funnelApi.ZF {
			lb.zfLB.Add(config.Url)
		}
	}
	CB = CircuitBreaker{
		lb:       lb,
		SnapShot: make(map[string]*apiSnapShot),
	}
}

type CircuitBreaker struct {
	lb       LoadBalance
	SnapShot map[string]*apiSnapShot
}

func (c *CircuitBreaker) GetApi(zfFlag, oauthFlag bool) (string, funnelApi.LoginType) {
	return c.lb.Pick(zfFlag, oauthFlag)
}

func (c *CircuitBreaker) Fail(api string) {
	if c.SnapShot[api].Fail() {
		c.lb.Remove(api, c.SnapShot[api].LoginType)
		Probe.Add(api, c.SnapShot[api].LoginType)
	}
}

func (c *CircuitBreaker) Success(api string) {
	c.lb.Add(api, c.SnapShot[api].LoginType)
	c.SnapShot[api].Success()
}
