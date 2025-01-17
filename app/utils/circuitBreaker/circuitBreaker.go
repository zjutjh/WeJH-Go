package circuitBreaker

import "wejh-go/config/api/funnelApi"

var CB CircuitBreaker

func init() {
	//Todo: add config
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
