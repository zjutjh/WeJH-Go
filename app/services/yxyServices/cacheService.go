package yxyServices

import (
	"context"
	"time"
	r "wejh-go/config/redis"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

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
		token, err := Auth(yxyUid)
		if err != nil {
			return nil, err
		}
		err = r.RedisClient.Set(ctx, cacheKey, *token, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}
		return token, nil
	} else if err != nil {
		return nil, err
	}
	return &cachedToken, nil
}
