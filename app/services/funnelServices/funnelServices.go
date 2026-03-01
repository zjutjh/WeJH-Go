package funnelServices

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"sync"

	"wejh-go/app/apiException"
	"wejh-go/app/utils/circuitBreaker"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/funnelApi"
)

// Funnel 协议中的业务返回码（与 funnel 端枚举保持一致）
const (
	funnelCodeSuccess        = 200 // 请求成功
	funnelCodeInvalidArgs    = 410 // 参数错误
	funnelCodeWrongPassword  = 412 // 密码错误
	funnelCodeCaptchaFailed  = 413 // 验证码错误
	funnelCodeSessionExpired = 414 // 缓存过期
	funnelCodeOauthError     = 415 // Oauth / 统一登录缓存错误
	funnelCodeOAuthNotUpdate = 416 // 统一密码未更新
)

// FunnelResponse 后端统一响应格式
type FunnelResponse struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}

// 对单个后端节点做一次调用：
// - 默认只请求一次
// - 如果返回 413（验证码错误），最多额外重试 4 次（总共 5 次）
func singleHostRequest(ctx context.Context, host string, api funnelApi.FunnelApi, form url.Values) (FunnelResponse, error) {
	f := fetch.Fetch{}
	f.Init()

	var rc FunnelResponse
	var res []byte
	var err error

	// 最多 5 次重试，遇到 413（验证码错误）才继续重试
	for i := 0; i < 5; i++ {
		// 已经被上层取消，则直接退出
		select {
		case <-ctx.Done():
			return FunnelResponse{}, ctx.Err()
		default:
		}

		res, err = f.PostForm(host+string(api), form)
		if err != nil {
			return FunnelResponse{}, apiException.RequestError
		}
		if err = json.Unmarshal(res, &rc); err != nil {
			return FunnelResponse{}, apiException.RequestError
		}

		// 只有 413（验证码错误）时才继续下一次重试，其它 code 直接退出循环
		if rc.Code != funnelCodeCaptchaFailed {
			break
		}
	}

	return rc, nil
}

// FetchHandleOfPost：
// - 非 ZF 接口：单节点调用（保留原逻辑）
// - ZF 接口：并发对冲到所有当前可用节点 + 简单熔断 / 恢复
func FetchHandleOfPost(form url.Values, host string, api funnelApi.FunnelApi) (interface{}, error) {
	loginType := funnelApi.LoginType(form.Get("type"))
	// 「是否 ZF 接口」用原来的约定：URL 中包含 "zf"
	zfFlag := strings.Contains(string(api), "zf")

	// 非 ZF 接口：保持原来的串行逻辑
	if !zfFlag {
		// 非对冲场景用 Background 的 ctx，行为与旧实现一致
		rc, err := singleHostRequest(context.Background(), host, api, form)
		if err != nil {
			// 对调用异常统一视为 ServerError
			return nil, apiException.ServerError
		}

		switch rc.Code {
		case funnelCodeSuccess:
			return rc.Data, nil

		case funnelCodeWrongPassword:
			// funnel 返回「密码错误」
			return nil, apiException.NoThatPasswordOrWrong

		case funnelCodeOAuthNotUpdate:
			// funnel 返回「统一密码未更新」
			return nil, apiException.OAuthNotUpdate

		// 410 / 413 / 414 / 415 以及其它未知 code：
		// 统一视为服务侧异常，向上抛 ServerError
		default:
			return nil, apiException.ServerError
		}
	}

	// 拿出当前健康的节点集合
	hosts, err := circuitBreaker.CB.LB.List(loginType)
	if err != nil {
		return nil, err
	}

	// 调用方通过 GetApi 传进来的 host 优先级最高，把它挪到列表最前面
	if host != "" {
		idx := -1
		for i, h := range hosts {
			if h == host {
				idx = i
				break
			}
		}

		if idx == -1 {
			// 原列表中没有这个 host，头插一份
			hosts = append([]string{host}, hosts...)
		} else if idx > 0 {
			// 已存在：和第一个元素交换，避免额外分配
			hosts[0], hosts[idx] = hosts[idx], hosts[0]
		}
	}

	if len(hosts) == 0 {
		// 所有节点都已熔断
		return nil, apiException.NoApiAvailable
	}

	type result struct {
		host string
		rc   FunnelResponse
		err  error
	}

	// 对冲用的 ctx：一旦某个节点拿到最终结果，cancel() 终止其它 goroutine 的后续工作
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultCh := make(chan result, len(hosts))
	var wg sync.WaitGroup

	// 并发对冲
	for _, h := range hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			// 独立 goroutine 需要自己的 recover，避免撞穿 gin 的 Recovery
			defer func() {
				if r := recover(); r != nil {
					// 保守起见，这里记一次 Fail，避免隐藏掉持续异常节点
					circuitBreaker.CB.Fail(h, loginType)
				}
			}()

			// 如果已经有其他节点成功了，可以尽量避免无意义请求
			select {
			case <-ctx.Done():
				return
			default:
			}

			rc, err := singleHostRequest(ctx, h, api, form)

			// 如果上层已经 cancel，不再阻塞在写 channel 上
			select {
			case resultCh <- result{host: h, rc: rc, err: err}:
			case <-ctx.Done():
				// 上层已经有结果了，丢弃即可
			}
		}(h)
	}

	// 等所有协程结束后关闭通道
	go func() {
		defer close(resultCh)
		wg.Wait()
	}()

	var firstErr error

	// 竞争结果：
	// - 第一个 200 直接返回数据，并标记该节点 Success
	// - 第一个 412 / 416 属于用户态错误，也直接返回，不再等其它节点
	// - 其它（410 / 413 / 414 / 415 / 未知）视作节点异常，Fail 一次，继续等待其它节点
	for r := range resultCh {
		if r.err != nil {
			// context.Canceled 表示其他节点已经成功，不认为该节点异常
			if errors.Is(r.err, context.Canceled) {
				continue
			}

			// 网络 / 解析错误等：认为节点异常
			circuitBreaker.CB.Fail(r.host, loginType)
			if firstErr == nil {
				// 统一向上映射成 ServerError
				firstErr = apiException.ServerError
			}
			continue
		}

		switch r.rc.Code {
		case funnelCodeSuccess:
			// 节点健康
			circuitBreaker.CB.Success(r.host, loginType)
			cancel()
			return r.rc.Data, nil

		case funnelCodeWrongPassword:
			// 密码错误：业务错误，节点本身是健康的
			circuitBreaker.CB.Success(r.host, loginType)
			cancel()
			return nil, apiException.NoThatPasswordOrWrong

		case funnelCodeOAuthNotUpdate:
			// 统一密码未更新：业务错误，节点健康
			circuitBreaker.CB.Success(r.host, loginType)
			cancel()
			return nil, apiException.OAuthNotUpdate

		// 410 / 413 / 414 / 415 以及其它未知 code
		default:
			// 视作节点异常（或至少是该线路上的不可用状态），触发熔断统计，继续等其它节点
			circuitBreaker.CB.Fail(r.host, loginType)
			if firstErr == nil {
				firstErr = apiException.ServerError
			}
		}
	}

	if firstErr == nil {
		firstErr = apiException.ServerError
	}
	return nil, firstErr
}
