package wechat

import (
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"log"
)

type driver string

const (
	Memory driver = "memory"
	Redis  driver = "redis"
)

var MiniProgram *miniprogram.MiniProgram

func init() {
	config := getConfigs()

	wc := wechat.NewWechat()
	var wcCache cache.Cache
	switch config.Driver {
	case string(Redis):
		wcCache = setRedis(wcCache)
		break
	case string(Memory):
		wcCache = cache.NewMemory()
		break
	default:
		log.Fatal("ConfigError")
	}

	cfg := &miniConfig.Config{
		AppID:     config.AppId,
		AppSecret: config.AppSecret,
		Cache:     wcCache,
	}

	MiniProgram = wc.GetMiniProgram(cfg)
}
