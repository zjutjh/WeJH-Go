package yxyServices

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	r "wejh-go/config/redis"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func GetElecRoomStrConcat(yxyUid, campus string) (*string, error) {
	cacheKey := "elec:room_str_concat:" + campus + ":" + yxyUid
	cachedRoomStrConcat, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if errors.Is(err, redis.Nil) {
		balance, err := ElectricityBalance(yxyUid, campus)
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

func GetElecConsumptionRecords(yxyUid, campus, roomStrConcat string) (*EleConsumptionRecords, error) {
	cacheKey := "elec:consumption_records:" + roomStrConcat
	cachedRecords, err := r.RedisClient.Get(ctx, cacheKey).Result()
	if errors.Is(err, redis.Nil) {
		records, err := ElectricityConsumptionRecords(yxyUid, campus, roomStrConcat)
		if err != nil {
			return nil, err
		}
		recordsJSON, err := json.Marshal(records)
		if err != nil {
			return nil, err
		}
		now := time.Now()
		midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		ttl := time.Until(midnight)
		err = r.RedisClient.Set(ctx, cacheKey, recordsJSON, ttl).Err()
		if err != nil {
			return nil, err
		}
		return records, nil
	} else if err != nil {
		return nil, err
	}
	var records EleConsumptionRecords
	err = json.Unmarshal([]byte(cachedRecords), &records)
	if err != nil {
		return nil, err
	}
	return &records, nil
}
