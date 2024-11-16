package wechat

import (
	"context"
	"wejh-go/config/redis"

	"github.com/silenceper/wechat/v2/cache"
)

func setRedis(wcCache cache.Cache) cache.Cache {

	redisOpts := &cache.RedisOpts{
		Host:        redis.RedisInfo.Host + ":" + redis.RedisInfo.Port,
		Database:    redis.RedisInfo.DB,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	}
	wcCache = cache.NewRedis(context.Background(), redisOpts)
	return wcCache
}
