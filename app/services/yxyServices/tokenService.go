package yxyServices

import (
	"context"
	"time"
	"wejh-go/config/redis"
)

var ctx = context.Background()

func GetToken(key string) (*string, error) {
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func SetToken(key, value string) {
	redis.RedisClient.Set(ctx, key, value, 15*time.Minute)
}

func CheckToken(key string) bool {
	intCmd := redis.RedisClient.Exists(ctx, key)
	if intCmd.Val() == 1 {
		return true
	} else {
		return false
	}
}

func DelToken(key string) {
	redis.RedisClient.Del(ctx, key)
}
