package circuitBreaker

import (
	"sync"
	"time"
	"wejh-go/app/models"
	"wejh-go/app/services/funnelServices"
	liveNessConfig "wejh-go/config/LiveNess"
	"wejh-go/config/api/funnelApi"
)

var Probe LiveNessProbe
var user *models.User

func init() {
	Probe = LiveNessProbe{
		Config: liveNessConfig.GetLiveNessConfig(),
		ApiMap: make(map[string]funnelApi.LoginType),
	}
	user.OauthPassword = Probe.Config.OauthPassword
	user.ZFPassword = Probe.Config.ZFPassword
	user.StudentID = Probe.Config.StudentId
	go Probe.Start()
}

type LiveNessProbe struct {
	sync.Mutex
	Config liveNessConfig.LiveNessProbeConfig
	ApiMap map[string]funnelApi.LoginType
}

func (l *LiveNessProbe) Add(api string, loginType funnelApi.LoginType) {
	l.Lock()
	defer l.Unlock()
	l.ApiMap[api] = loginType
}

func (l *LiveNessProbe) Remove(key string) {
	l.Lock()
	defer l.Unlock()
	delete(l.ApiMap, key)
}

func (l *LiveNessProbe) Start() {
	ticker := time.NewTicker(l.Config.Duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for api, loginType := range l.ApiMap {
				_, err := funnelServices.GetClassTable(user, "2023", "上", api, loginType)
				if err == nil {
					CB.Success(api)
					delete(l.ApiMap, api)
				}
			}
		}
	}
}
