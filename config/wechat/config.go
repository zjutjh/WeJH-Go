package wechat

import (
	"go.uber.org/zap"
	"strings"
	"wejh-go/config/config"
)

type wechatConfig struct {
	Driver    string
	AppId     string
	AppSecret string
}

func getConfigs() wechatConfig {

	wc := wechatConfig{}
	if !config.Config.IsSet("wechat.appid") {
		zap.L().Fatal("Config Error")
	}
	if !config.Config.IsSet("wechat.appsecret") {
		zap.L().Fatal("Config Error")
	}
	wc.AppId = config.Config.GetString("wechat.appid")
	wc.AppSecret = config.Config.GetString("wechat.appsecret")

	wc.Driver = string(Memory)
	if config.Config.IsSet("wechat.driver") {
		wc.Driver = strings.ToLower(config.Config.GetString("wechat.driver"))
	}
	return wc
}
