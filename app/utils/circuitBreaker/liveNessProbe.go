package circuitBreaker

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/utils/fetch"

	"wejh-go/config/api/funnelApi"
	cbConfig "wejh-go/config/circuitBreaker"
)

var Probe *LiveNessProbe

func init() {
	Probe = NewLiveNessProbe(cbConfig.GetLiveNessConfig())
}

type LiveNessProbe struct {
	sync.Mutex
	ApiMap   map[string]funnelApi.LoginType
	Duration time.Duration
	User     *models.User
}

func NewLiveNessProbe(config cbConfig.LiveNessProbeConfig) *LiveNessProbe {
	user := &models.User{
		StudentID:     config.StudentId,
		OauthPassword: config.OauthPassword,
		ZFPassword:    config.ZFPassword,
	}
	return &LiveNessProbe{
		ApiMap:   make(map[string]funnelApi.LoginType),
		Duration: config.Duration,
		User:     user,
	}
}

func (l *LiveNessProbe) Add(api string, loginType funnelApi.LoginType) {
	l.Lock()
	defer l.Unlock()
	l.ApiMap[api+string(loginType)] = loginType
}

func (l *LiveNessProbe) Remove(key string) {
	l.Lock()
	defer l.Unlock()
	delete(l.ApiMap, key)
}

func (l *LiveNessProbe) Start(ctx context.Context) {
	ticker := time.NewTicker(l.Duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for apiKey, loginType := range l.ApiMap {
				api := strings.TrimSuffix(apiKey, string(loginType))
				if err := liveNess(l.User, api, loginType); err == nil {
					CB.LB.Add(api, loginType)
					CB.Success(api, loginType)
					l.Remove(apiKey)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func liveNess(u *models.User, api string, loginType funnelApi.LoginType) error {
	var password string
	if loginType == funnelApi.Oauth {
		password = u.OauthPassword
	} else {
		password = u.ZFPassword
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", string(loginType))
	form.Add("year", strconv.Itoa(time.Now().Year()-1))
	form.Add("term", "上")

	f := fetch.Fetch{}
	f.Init()

	rc := struct {
		Code int `json:"code" binding:"required"`
	}{}
	for i := 0; i < 5; i++ {
		res, err := f.PostForm(api+string(funnelApi.ZFClassTable), form)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(res, &rc); err != nil {
			return err
		}
		if rc.Code != 413 {
			break
		}
	}
	if rc.Code == 200 || rc.Code == 412 || rc.Code == 416 {
		return nil
	}
	return apiException.ServerError
}
