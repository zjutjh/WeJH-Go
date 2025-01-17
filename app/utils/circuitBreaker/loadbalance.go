package circuitBreaker

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"wejh-go/config/api/funnelApi"
)

type LoadBalanceType int

const (
	RoundRobin LoadBalanceType = iota
	Random
)

type LoadBalance struct {
	zfLB    *randomLB
	oauthLB *randomLB
}

func (lb *LoadBalance) Pick(zfFlag, oauthFlag bool) (string, funnelApi.LoginType) {
	var loginType funnelApi.LoginType

	if oauthFlag && zfFlag {
		if fastrand.Intn(100) > 50 {
			loginType = funnelApi.Oauth
		} else {
			loginType = funnelApi.ZF
		}
	} else if oauthFlag {
		loginType = funnelApi.Oauth
	} else if zfFlag {
		loginType = funnelApi.ZF
	} else {
		return "", funnelApi.Unknown
	}
	if loginType == funnelApi.Oauth {
		return lb.zfLB.Pick(), loginType
	}
	return lb.oauthLB.Pick(), loginType
}

func (lb *LoadBalance) Remove(api string, loginType funnelApi.LoginType) {
	if loginType == funnelApi.Oauth {
		lb.oauthLB.Remove(api)
	} else {
		lb.zfLB.Remove(api)
	}
}

func (lb *LoadBalance) Add(api string, loginType funnelApi.LoginType) {
	if loginType == funnelApi.Oauth {
		lb.oauthLB.Add(api)
	} else {
		lb.zfLB.Add(api)
	}
}

type loadBalance interface {
	LoadBalance() LoadBalanceType
	Pick() (api string)
	ReBalance(apis []string)
	Remove(api string)
}

type randomLB struct {
	Api  []string
	Size int
}

func newRandomLB(apis []string) loadBalance {
	return &randomLB{Api: apis, Size: len(apis)}
}

func (b *randomLB) LoadBalance() LoadBalanceType {
	return Random
}

func (b *randomLB) Pick() string {
	idx := fastrand.Intn(b.Size)
	return b.Api[idx]
}

func (b *randomLB) ReBalance(apis []string) {
	b.Api, b.Size = apis, len(apis)
}

func (b *randomLB) Add(api ...string) {
	b.Api = append(b.Api, api...)
	b.Size = len(b.Api)
}

func (b *randomLB) Remove(api string) {
	for i, s := range b.Api {
		if s == api {
			b.Api = append(b.Api[:i], b.Api[i+1:]...)
			break
		}
	}
}
