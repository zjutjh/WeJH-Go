package wechat

import (
	"context"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/zjutjh/mygo/config"
)

func setRedis(wcCache cache.Cache) cache.Cache {

	redisOpts := &cache.RedisOpts{
		Host:        config.Pick().GetString("redis.host") + ":" + config.Pick().GetString("redis.port"),
		Database:    config.Pick().GetInt("redis.db"),
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	}
	wcCache = cache.NewRedis(context.Background(), redisOpts)
	return wcCache
}
