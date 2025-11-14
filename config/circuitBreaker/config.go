package circuitBreaker

import (
	"time"

	"github.com/zjutjh/mygo/config"
)

type LiveNessProbeConfig struct {
	StudentId     string
	OauthPassword string
	ZFPassword    string
	Duration      time.Duration
}

func GetLiveNessConfig() LiveNessProbeConfig {
	return LiveNessProbeConfig{
		StudentId:     config.Pick().GetString("zfCircuit.studentId"),
		OauthPassword: config.Pick().GetString("zfCircuit.oauthPassword"),
		ZFPassword:    config.Pick().GetString("zfCircuit.zfPassword"),
		Duration:      config.Pick().GetDuration("zfCircuit.duration"),
	}
}

type LoadBalanceConfig struct {
	Apis []string
}

func GetLoadBalanceConfig() LoadBalanceConfig {
	return LoadBalanceConfig{
		Apis: config.Pick().GetStringSlice("zfCircuit.apis"),
	}
}
