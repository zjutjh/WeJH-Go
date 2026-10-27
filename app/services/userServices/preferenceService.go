package userServices

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wejh-go/app/models"
	"wejh-go/config/database"
	"wejh-go/config/redis"

	goredis "github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

var ErrInvalidTheme = errors.New("无效的主题模式")

func UpdateUserThemePreference(userID uint, theme string) error {
	// 1. 兜底校验
	if theme != "light" && theme != "dark" && theme != "auto" {
		return ErrInvalidTheme
	}

	// 2. 更新数据库
	err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("theme_preference", theme).Error
	if err != nil {

		zap.L().Error("更新用户主题偏好失败",
			zap.Uint("userID", userID),
			zap.String("theme", theme),
			zap.Error(err),
		)
		return err
	}
	// 3. 删除 Redis 缓存（下次读取时会重新从数据库加载最新值）
	ctx := context.Background()
	cacheKey := fmt.Sprintf(userThemeCachePrefix, userID)
	if err := redis.RedisClient.Del(ctx, cacheKey).Err(); err != nil {
		//缓存删除失败不阻断主流程，只记录日志
		zap.L().Error("删除用户主题缓存失败", zap.String("key", cacheKey), zap.Error(err))
	}

	zap.L().Info("更新用户主题偏好成功", zap.Uint("userID", userID), zap.String("theme", theme))
	return nil
}

// 获取用户主题偏好（供 GetUserInfo 或其他地方使用）
const userThemeCachePrefix = "wejh:user:theme:%d"

func GetUserThemePreference(userID uint) (string, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf(userThemeCachePrefix, userID)

	// 1. 优先从 Redis 读取
	theme, err := redis.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		return theme, nil // 缓存命中，直接返回
	}
	if !errors.Is(err, goredis.Nil) {
		// 如果不是 key 不存在，而是真正的系统错误（比如网络断了），记录日志
		zap.L().Error("读取用户主题缓存失败", zap.String("key", cacheKey), zap.Error(err))
	}
	// 2. 缓存未命中，从 MySQL 查询
	var user models.User
	if err := database.DB.Select("theme_preference").First(&user, userID).Error; err != nil {
		zap.L().Error("查询用户主题数据库失败", zap.Uint("userID", userID), zap.Error(err))
		return "", err
	}

	// 如果数据库里是空的，给个默认值
	theme = user.ThemePreference
	if theme == "" {
		theme = "auto"
	}

	// 3. 写入 Redis，设置 7 天过期
	if err := redis.RedisClient.Set(ctx, cacheKey, theme, 7*24*time.Hour).Err(); err != nil {
		zap.L().Error("写入用户主题缓存失败", zap.String("key", cacheKey), zap.Error(err))
	}

	return theme, nil
}
