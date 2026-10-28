package userServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"

	"go.uber.org/zap"
)

func UpdateUserThemePreference(userID uint, theme string) error {
	// 直接更新数据库
	err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("theme_preference", theme).Error
	if err != nil {
		zap.L().Error("更新用户主题偏好数据库失败", zap.Uint("userID", userID), zap.Error(err))
		return err
	}

	zap.L().Info("更新用户主题偏好成功", zap.Uint("userID", userID), zap.String("theme", theme))
	return nil
}
