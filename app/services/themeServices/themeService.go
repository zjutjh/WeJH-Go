package themeServices

import (
	"encoding/json"
	"errors"
	"wejh-go/app/models"
	"wejh-go/config/database"

	"gorm.io/gorm"
)

func CheckThemeExist(ids ...int) error {
	var themes []models.Theme
	result := database.DB.Model(&models.Theme{}).
		Where("id IN ?", ids).
		Find(&themes)
	if result.Error != nil {
		return result.Error
	}
	if len(themes) != len(ids) {
		return errors.New("some theme not exist")
	}
	return nil
}

func CreateTheme(themeName, themeType string, isDarkMode bool, themeConfig models.ThemeConfig) error {
	config, err := json.Marshal(themeConfig)
	if err != nil {
		return err
	}
	result := database.DB.Create(&models.Theme{
		Name:        themeName,
		Type:        themeType,
		IsDarkMode:  isDarkMode,
		ThemeConfig: string(config),
	})
	return result.Error
}

func UpdateTheme(themeID int, themeName string, isDarkMode bool, themeConfig models.ThemeConfig) error {
	config, err := json.Marshal(themeConfig)
	if err != nil {
		return err
	}
	result := database.DB.Model(&models.Theme{}).
		Where("id = ?", themeID).
		Select("name", "is_dark_mode", "theme_config").
		Updates(&models.Theme{
			Name:        themeName,
			IsDarkMode:  isDarkMode,
			ThemeConfig: string(config),
		})
	return result.Error
}

func GetThemeByID(id int) (models.Theme, error) {
	var theme models.Theme
	if err := database.DB.Model(models.Theme{}).
		Where(&models.Theme{ID: id}).
		First(&theme).Error; err != nil {
		return theme, err
	}
	return theme, nil
}

func GetAllTheme() ([]models.FormatTheme, error) {
	var themes []models.Theme
	result := database.DB.Model(models.Theme{}).Find(&themes)
	if result.Error != nil {
		return nil, result.Error
	}

	var formatThemes []models.FormatTheme
	for _, theme := range themes {
		var themeConfig models.ThemeConfig
		if err := json.Unmarshal([]byte(theme.ThemeConfig), &themeConfig); err != nil {
			return nil, err
		}
		formatTheme := models.FormatTheme{
			Name:        theme.Name,
			ThemeID:     theme.ID,
			ThemeConfig: themeConfig,
			IsDarkMode:  theme.IsDarkMode,
		}
		formatThemes = append(formatThemes, formatTheme)
	}

	return formatThemes, nil
}

func DeleteTheme(id int, themeType string, isDarkMode bool) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Theme{}, id).Error; err != nil {
			return err
		}

		var themeID int
		if err := tx.Model(&models.Theme{}).
			Select("id").
			Where("type = ? AND is_dark_mode = ?", "all", isDarkMode).
			First(&themeID).Error; err != nil {
			return err
		}

		updateField := map[bool]string{false: "current_theme_id", true: "current_theme_dark_id"}[isDarkMode]
		if err := tx.Model(&models.ThemePermission{}).
			Where(updateField+" = ?", id).
			Update(updateField, themeID).Error; err != nil {
			return err
		}

		if themeType == "all" {
			return nil
		}

		var permissions []models.ThemePermission
		if err := tx.Model(models.ThemePermission{}).Find(&permissions).Error; err != nil {
			return err
		}

		updatedPermissionMap := make(map[string]models.ThemePermissionData)
		for _, p := range permissions {
			var data models.ThemePermissionData
			err := json.Unmarshal([]byte(p.ThemePermission), &data)
			if err != nil {
				return err
			}
			updatedThemeIDs := removeThemeID(data.ThemeIDs, id)
			if len(updatedThemeIDs) != len(data.ThemeIDs) {
				data.ThemeIDs = updatedThemeIDs
				if len(updatedThemeIDs) == 0 {
					data.ThemeIDs = []int{}
				}
				updatedPermissionMap[p.StudentID] = data
			}
		}

		for studentID, data := range updatedPermissionMap {
			newPermission, err := json.Marshal(data)
			if err != nil {
				return err
			}
			if err := tx.Model(&models.ThemePermission{}).
				Where("student_id = ?", studentID).
				Update("theme_permission", string(newPermission)).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func removeThemeID(themeIDs []int, id int) []int {
	var updatedThemeIDs []int
	for _, themeID := range themeIDs {
		if themeID != id {
			updatedThemeIDs = append(updatedThemeIDs, themeID)
		}
	}
	return updatedThemeIDs
}
