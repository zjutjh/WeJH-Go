package yxyServices

import (
	"context"
	"time"
	r "wejh-go/config/redis"

	"golang.org/x/sync/singleflight"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	g   singleflight.Group
)

func GetElecRoomStrConcat(token, campus, yxyUid string) (*string, error) {
	cacheKey := "elec:room_str_concat:" + campus + ":" + yxyUid
	cachedRoomStrConcat, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		balance, err := ElectricityBalance(token, campus)
		if err != nil {
			return nil, err
		}
		err = r.RedisClient.Set(ctx, cacheKey, balance.RoomStrConcat, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}
		return &balance.RoomStrConcat, nil
	} else if err != nil {
		return nil, err
	}
	return &cachedRoomStrConcat, nil
}

func GetElecAuthToken(yxyUid string) (*string, error) {
	cacheKey := "elec:auth_token:" + yxyUid
	cachedToken, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// 使用 singleflight 防止缓存击穿
		token, err, _ := g.Do(cacheKey, func() (interface{}, error) {
			t, e := Auth(yxyUid)
			if e != nil {
				return nil, e
			}
			e = r.RedisClient.Set(ctx, cacheKey, *t, 7*24*time.Hour).Err()
			if e != nil {
				return nil, e
			}
			return t, nil
		})
		if err != nil {
			return nil, err
		}
		return token.(*string), nil
	} else if err != nil {
		return nil, err
	}
	return &cachedToken, nil
}
