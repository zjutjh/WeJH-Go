package circuitBreaker

import (
	"time"
	"wejh-go/config/config"
)

type LiveNessProbeConfig struct {
	StudentId     string
	OauthPassword string
	ZFPassword    string
	Duration      time.Duration
}

func GetLiveNessConfig() LiveNessProbeConfig {
	return LiveNessProbeConfig{
		StudentId:     config.Config.GetString("zfCircuit.studentId"),
		OauthPassword: config.Config.GetString("zfCircuit.oauthPassword"),
		ZFPassword:    config.Config.GetString("zfCircuit.zfPassword"),
		Duration:      config.Config.GetDuration("zfCircuit.duration"),
	}
}

type LoadBalanceConfig struct {
	Apis []string
}

func GetLoadBalanceConfig() LoadBalanceConfig {
	return LoadBalanceConfig{
		Apis: config.Config.GetStringSlice("zfCircuit.apis"),
	}
}
