package liveNessConfig

import (
	"github.com/spf13/viper"
	"log"
	"time"
	"wejh-go/config/api/funnelApi"
	"wejh-go/config/config"
)

type LiveNessProbeConfig struct {
	StudentId     string
	OauthPassword string
	ZFPassword    string
	Duration      time.Duration
}

func GetLiveNessConfig() LiveNessProbeConfig {
	cfg := LiveNessProbeConfig{}
	if config.Config.IsSet("liveNess.studentId") {
		cfg.StudentId = config.Config.GetString("liveNess.studentId")
	}
	if config.Config.IsSet("liveNess.oauthPassword") {
		cfg.OauthPassword = config.Config.GetString("liveNess.oauthPassword")
	}
	if config.Config.IsSet("liveNess.zfPassword") {
		cfg.ZFPassword = config.Config.GetString("liveNess.zfPassword")
	}
	if config.Config.IsSet("liveNess.duration") {
		cfg.Duration = config.Config.GetDuration("liveNess.duration")
	}
	return cfg
}

type LoadBalanceConfig struct {
	Name string              `json:"name"`
	Url  string              `json:"url"`
	Type funnelApi.LoginType `json:"type"`
}

func GetLoadBalanceConfig() []LoadBalanceConfig {
	var configs []LoadBalanceConfig
	err := viper.UnmarshalKey("loadBalance", &configs)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configs
}
