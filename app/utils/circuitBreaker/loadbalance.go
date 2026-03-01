package circuitBreaker

import (
	"sync"

	"github.com/bytedance/gopkg/lang/fastrand"
	"wejh-go/app/apiException"
	"wejh-go/config/api/funnelApi"
	"wejh-go/config/logger"
)

type LoadBalanceType int

const (
	Random LoadBalanceType = iota
)

// LoadBalance 维护两套池：ZF / OAuth
type LoadBalance struct {
	zfLB    *randomLB
	oauthLB *randomLB
}

// Pick 原有随机负载均衡逻辑（为了兼容）
func (lb *LoadBalance) Pick(zfFlag, oauthFlag bool) (string, funnelApi.LoginType, error) {
	oauthAvailable := oauthFlag && lb.oauthLB.isAvailable()
	zfAvailable := zfFlag && lb.zfLB.isAvailable()

	if oauthAvailable && zfAvailable {
		// 双池都可用时，随机选择其一
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

// List 返回当前可用后端节点的快照
// - loginType 为 Oauth：返回 OAuth 池
// - loginType 为 ZF：返回 ZF 池
// - Unknown/非法值：返回错误并记录告警日志
func (lb *LoadBalance) List(loginType funnelApi.LoginType) ([]string, error) {
	switch loginType {
	case funnelApi.Oauth:
		if lb.oauthLB == nil {
			return nil, nil
		}
		return lb.oauthLB.list(), nil
	case funnelApi.ZF:
		if lb.zfLB == nil {
			return nil, nil
		}
		return lb.zfLB.list(), nil
	default:
		logger.GetLogFunc(logger.LevelWarn)("unknown login type for load balancer list: " + string(loginType))
		return nil, apiException.ParamError
	}
}

// 在运行时添加节点
func (lb *LoadBalance) Add(api string, loginType funnelApi.LoginType) {
	if loginType == funnelApi.Oauth {
		lb.oauthLB.Add(api)
	} else {
		lb.zfLB.Add(api)
	}
}

// 在运行时移除节点
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
	Api []string
}

func newRandomLB(apis []string) loadBalance {
	return &randomLB{
		Api: apis,
	}
}

func (b *randomLB) LoadBalance() LoadBalanceType {
	return Random
}

// Pick：单次随机选择一个后端
func (b *randomLB) Pick() string {
	b.Lock()
	defer b.Unlock()

	size := len(b.Api)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return b.Api[0]
	}
	return b.Api[fastrand.Intn(size)]
}

// list：返回当前后端列表的拷贝，供并发对冲使用
func (b *randomLB) list() []string {
	b.Lock()
	defer b.Unlock()

	size := len(b.Api)
	if size == 0 {
		return nil
	}
	out := make([]string, size)
	copy(out, b.Api)
	return out
}

func (b *randomLB) ReBalance(apis []string) {
	b.Lock()
	defer b.Unlock()
	b.Api = apis
}

func (b *randomLB) Add(api ...string) {
	b.Lock()
	defer b.Unlock()
	b.Api = append(b.Api, api...)
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
}

func (b *randomLB) isAvailable() bool {
	b.Lock()
	defer b.Unlock()
	return len(b.Api) != 0
}
