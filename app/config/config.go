package config

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
	"wejh-go/app/models"
	"wejh-go/config/database"
	"wejh-go/config/redis"
)

var ctx = context.Background()

func getConfig(key string) string {
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err == nil {
		return val
	}
	print(err)
	var config = &models.Config{}
	database.DB.Model(models.Config{}).Where(
		&models.Config{
			Key: key,
		}).First(&config)

	redis.RedisClient.Set(ctx, key, config.Value, 0)
	return config.Value
}

func setConfig(key, value string) error {
	redis.RedisClient.Set(ctx, key, value, 0)
	var config models.Config
	result := database.DB.Where("`key` = ?", key).First(&config)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		config = models.Config{
			Key:        key,
			Value:      value,
			UpdateTime: time.Now(),
		}
		result = database.DB.Create(&config)
	} else {
		config.Value = value
		config.UpdateTime = time.Now()
		result = database.DB.Updates(&config)
	}
	return result.Error
}

func checkConfig(key string) bool {
	intCmd := redis.RedisClient.Exists(ctx, key)
	if intCmd.Val() == 1 {
		return true
	} else {
		return false
	}
}

func delConfig(key string) error {
	redis.RedisClient.Del(ctx, key)
	res := database.DB.Where(&models.Config{
		Key: key,
	}).Delete(models.Config{})
	return res.Error
}
