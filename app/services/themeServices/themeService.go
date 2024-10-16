package themeServices

import (
	"encoding/json"
	"strconv"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func CheckThemeExist(id int) error {
	var theme models.Theme
	result := database.DB.Model(models.Theme{}).Where("id = ?", id).First(&theme)
	return result.Error
}

func CreateTheme(record models.Theme) error {
	result := database.DB.Create(&record)
	return result.Error
}

func UpdateTheme(id int, record models.Theme) error {
	result := database.DB.Model(models.Theme{}).Where(&models.Theme{ID: id}).Updates(&record)
	return result.Error
}

func GetThemeByID(id int) (models.Theme, error) {
	var record models.Theme
	result := database.DB.Model(models.Theme{}).Where(&models.Theme{ID: id}).First(&record)
	return record, result.Error
}

func GetThemes() ([]models.Theme, error) {
	var themes []models.Theme
	result := database.DB.Model(models.Theme{}).Find(&themes)
	return themes, result.Error
}

func DeleteTheme(id int, themeType string) error {
	tx := database.DB.Begin()
	if err := tx.Delete(&models.Theme{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	var theme models.Theme
	if err := tx.Where("type = ?", "all").First(&theme).Error; err != nil {
		tx.Rollback()
		return err
	}

	var defaultThemeID int
	defaultThemeIDStr := config.GetDefaultThemeKey()
	if defaultThemeIDStr != "" {
		defaultThemeID, _ = strconv.Atoi(defaultThemeIDStr)
		if id == defaultThemeID {
			defaultThemeID = theme.ID
			err := config.SetDefaultThemeKey(strconv.Itoa(defaultThemeID))
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		defaultThemeID = theme.ID
	}

	if err := tx.Model(&models.ThemePermission{}).
		Where("current_theme_id = ?", id).
		Update("current_theme_id", defaultThemeID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if themeType == "all" {
		tx.Commit()
		return nil
	}

	var permissions []models.ThemePermission
	result := tx.Model(models.ThemePermission{}).Find(&permissions)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	updatedPermissionMap := make(map[string]models.ThemePermissionData)
	for _, permission := range permissions {
		var themePermissionData models.ThemePermissionData
		err := json.Unmarshal([]byte(permission.ThemePermission), &themePermissionData)
		if err != nil {
			tx.Rollback()
			return err
		}
		updatedThemeIDs := removeThemeID(themePermissionData.ThemeIDs, id)
		if len(updatedThemeIDs) != len(themePermissionData.ThemeIDs) {
			themePermissionData.ThemeIDs = updatedThemeIDs
			if len(updatedThemeIDs) == 0 {
				themePermissionData.ThemeIDs = []int{}
			}
			updatedPermissionMap[permission.StudentID] = themePermissionData
		}
	}
	for studentID, data := range updatedPermissionMap {
		newPermission, err := json.Marshal(data)
		if err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Model(&models.ThemePermission{}).
			Where("student_id = ?", studentID).
			Update("theme_permission", string(newPermission)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
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
