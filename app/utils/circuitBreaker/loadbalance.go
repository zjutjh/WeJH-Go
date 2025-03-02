package circuitBreaker

import (
	"github.com/bytedance/gopkg/lang/fastrand"
	"sync"
	"wejh-go/app/apiException"
	"wejh-go/config/api/funnelApi"
)

type LoadBalanceType int

const (
	Random LoadBalanceType = iota
)

type LoadBalance struct {
	zfLB    *randomLB
	oauthLB *randomLB
}

func (lb *LoadBalance) Pick(zfFlag, oauthFlag bool) (string, funnelApi.LoginType, error) {
	oauthAvailable := oauthFlag && lb.oauthLB.isAvailable()
	zfAvailable := zfFlag && lb.zfLB.isAvailable()

	if oauthAvailable && zfAvailable {
		if fastrand.Intn(100) > 50 {
			return lb.oauthLB.Pick(), funnelApi.Oauth, nil
		}
		return lb.zfLB.Pick(), funnelApi.ZF, nil
	}

	if oauthAvailable {
		return lb.oauthLB.Pick(), funnelApi.Oauth, nil
	}
	if zfAvailable {
		return lb.zfLB.Pick(), funnelApi.ZF, nil
	}

	return "", funnelApi.Unknown, apiException.NoApiAvailable
}

func (lb *LoadBalance) Add(api string, loginType funnelApi.LoginType) {
	if loginType == funnelApi.Oauth {
		lb.oauthLB.Add(api)
	} else {
		lb.zfLB.Add(api)
	}
}

func (lb *LoadBalance) Remove(api string, loginType funnelApi.LoginType) {
	if loginType == funnelApi.Oauth {
		lb.oauthLB.Remove(api)
	} else {
		lb.zfLB.Remove(api)
	}
}

type loadBalance interface {
	LoadBalance() LoadBalanceType
	Pick() (api string)
	ReBalance(apis []string)
	Add(api ...string)
	Remove(api string)
	isAvailable() bool
}

type randomLB struct {
	sync.Mutex
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
	b.Lock()
	defer b.Unlock()
	if b.Size == 0 {
		return ""
	}
	idx := fastrand.Intn(b.Size)
	return b.Api[idx]
}

func (b *randomLB) ReBalance(apis []string) {
	b.Lock()
	defer b.Unlock()
	b.Api, b.Size = apis, len(apis)
}

func (b *randomLB) Add(api ...string) {
	b.Lock()
	defer b.Unlock()
	b.Api = append(b.Api, api...)
	b.Size = len(b.Api)
}

func (b *randomLB) Remove(api string) {
	b.Lock()
	defer b.Unlock()
	for i, s := range b.Api {
		if s == api {
			b.Api = append(b.Api[:i], b.Api[i+1:]...)
			break
		}
	}
	b.Size = len(b.Api)
}

func (b *randomLB) isAvailable() bool {
	b.Lock()
	defer b.Unlock()
	return b.Size != 0
}
