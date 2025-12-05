package config

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
	"github.com/zjutjh/mygo/nedis"
	"gorm.io/gorm"
)

var ctx = context.Background()

func getConfig(key string) string {
	val, err := nedis.Pick().Get(ctx, key).Result()
	if err == nil {
		return val
	}
	print(err)
	var config = &models.Config{}
	ndb.Pick().Model(models.Config{}).Where(
		&models.Config{
			Key: key,
		}).First(&config)

	nedis.Pick().Set(ctx, key, config.Value, 0)
	return config.Value
}

func setConfig(key, value string) error {
	nedis.Pick().Set(ctx, key, value, 0)
	var config models.Config
	result := ndb.Pick().Where("`key` = ?", key).First(&config)
	fmt.Print(config)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		config = models.Config{
			Key:        key,
			Value:      value,
			UpdateTime: time.Now(),
		}
		result = ndb.Pick().Create(&config)
	} else {
		config.Value = value
		config.UpdateTime = time.Now()
		result = ndb.Pick().Updates(&config)
	}
	return result.Error
}

func checkConfig(key string) bool {
	intCmd := nedis.Pick().Exists(ctx, key)
	if intCmd.Val() == 1 {
		return true
	} else {
		return false
	}
}

func delConfig(key string) error {
	nedis.Pick().Del(ctx, key)
	res := ndb.Pick().Where(&models.Config{
		Key: key,
	}).Delete(models.Config{})
	return res.Error
}
