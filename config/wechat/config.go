package wechat

import (
	"strings"

	"github.com/zjutjh/mygo/config"
	"go.uber.org/zap"
)

type wechatConfig struct {
	Driver    string
	AppId     string
	AppSecret string
}

func getConfigs() wechatConfig {

	wc := wechatConfig{}
	if !config.Pick().IsSet("wechat.appid") {
		zap.L().Fatal("Config Error")
	}
	if !config.Pick().IsSet("wechat.appsecret") {
		zap.L().Fatal("Config Error")
	}
	wc.AppId = config.Pick().GetString("wechat.appid")
	wc.AppSecret = config.Pick().GetString("wechat.appsecret")

	wc.Driver = string(Memory)
	if config.Pick().IsSet("wechat.driver") {
		wc.Driver = strings.ToLower(config.Pick().GetString("wechat.driver"))
	}
	return wc
}
