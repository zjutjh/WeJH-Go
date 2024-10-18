package themeServices

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func AddThemePermission(themeID int, reqStudentIDs []string, themeType string) ([]string, error) {
	if themeType == "all" {
		return nil, nil
	}

	var studentIDs []string
	var invalidStudentIDs []string
	if len(reqStudentIDs) > 0 {
		var existingUsers []models.User
		err := database.DB.Select("student_id").Where("student_id IN ?", reqStudentIDs).Find(&existingUsers).Error
		if err != nil {
			return nil, err
		}

		existingStudentIDMap := make(map[string]bool)
		for _, user := range existingUsers {
			existingStudentIDMap[user.StudentID] = true
		}

		for _, studentID := range reqStudentIDs {
			if existingStudentIDMap[studentID] {
				studentIDs = append(studentIDs, studentID)
			} else {
				invalidStudentIDs = append(invalidStudentIDs, studentID)
			}
		}
	} else {
		return nil, errors.New("reqStudentIDs is invalid")
	}
	if len(studentIDs) == 0 {
		return invalidStudentIDs, nil
	}

	var permissions []models.ThemePermission
	err := database.DB.Where("student_id IN ?", studentIDs).Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	permissionMap := make(map[string]*models.ThemePermission)
	for i, permission := range permissions {
		permissionMap[permission.StudentID] = &permissions[i]
	}

	for _, studentID := range studentIDs {
		permission, exist := permissionMap[studentID]
		if !exist {
			themePermissionData := models.ThemePermissionData{
				ThemeIDs: []int{themeID},
			}
			newPermissionData, err := json.Marshal(themePermissionData)
			if err != nil {
				return nil, err
			}
			newPermission := models.ThemePermission{
				StudentID:       studentID,
				CurrentThemeID:  themeID,
				ThemePermission: string(newPermissionData),
			}
			permission = &newPermission
			permissions = append(permissions, newPermission)
		}

		var themePermissionData models.ThemePermissionData
		err = json.Unmarshal([]byte(permission.ThemePermission), &themePermissionData)
		if err != nil {
			return nil, err
		}
		if !containThemeID(themePermissionData.ThemeIDs, themeID) {
			themePermissionData.ThemeIDs = append(themePermissionData.ThemeIDs, themeID)
			newPermission, err := json.Marshal(themePermissionData)
			if err != nil {
				return nil, err
			}
			permission.ThemePermission = string(newPermission)
		}
	}

	// 使用批量保存
	err = savePermissionsInBatches(permissions)
	if err != nil {
		return nil, err
	}
	return invalidStudentIDs, nil
}

func UpdateCurrentTheme(id int, studentID string) error {
	var permission models.ThemePermission
	err := database.DB.Where("student_id = ?", studentID).First(&permission).Error
	if err != nil {
		return err
	}

	var themePermissionData models.ThemePermissionData
	err = json.Unmarshal([]byte(permission.ThemePermission), &themePermissionData)
	if err != nil {
		return err
	}

	var allThemes []models.Theme
	err = database.DB.Where("type = ?", "all").Find(&allThemes).Error
	if err != nil {
		return err
	}

	for _, theme := range allThemes {
		if !containThemeID(themePermissionData.ThemeIDs, theme.ID) {
			themePermissionData.ThemeIDs = append(themePermissionData.ThemeIDs, theme.ID)
		}
	}

	if !containThemeID(themePermissionData.ThemeIDs, id) {
		return errors.New("the theme ID is not in the user's permission list")
	}

	result := database.DB.Model(models.ThemePermission{}).
		Where("student_id = ?", studentID).
		Update("current_theme_id", id)
	return result.Error
}

func DeleteThemePermission(studentID string) error {
	result := database.DB.Where("student_id = ?", studentID).Delete(&models.ThemePermission{})
	return result.Error
}

func GetThemePermissionByStudentID(studentID string) (models.ThemePermission, error) {
	var record models.ThemePermission
	result := database.DB.Model(models.ThemePermission{}).Where("student_id = ?", studentID).First(&record)
	return record, result.Error
}

func GetThemeNameByID(themePermission models.ThemePermission) ([]string, error) {
	var themePermissionData models.ThemePermissionData
	err := json.Unmarshal([]byte(themePermission.ThemePermission), &themePermissionData)
	if err != nil {
		return nil, err
	}

	var themes []models.Theme
	err = database.DB.Where("id IN ?", themePermissionData.ThemeIDs).Find(&themes).Error
	if err != nil {
		return nil, err
	}

	var allThemes []models.Theme
	err = database.DB.Where("type = ?", "all").Find(&allThemes).Error
	if err != nil {
		return nil, err
	}

	for _, allTheme := range allThemes {
		if !containThemeID(themePermissionData.ThemeIDs, allTheme.ID) {
			themes = append(themes, allTheme)
		}
	}

	var themeNames []string
	for _, theme := range themes {
		themeNames = append(themeNames, theme.Name)
	}
	return themeNames, nil
}

func GetThemesByID(themePermission models.ThemePermission) ([]models.Theme, error) {
	var themePermissionData models.ThemePermissionData
	err := json.Unmarshal([]byte(themePermission.ThemePermission), &themePermissionData)
	if err != nil {
		return nil, err
	}

	var themes []models.Theme
	err = database.DB.Where("id IN ?", themePermissionData.ThemeIDs).Find(&themes).Error
	if err != nil {
		return nil, err
	}

	var allThemes []models.Theme
	err = database.DB.Where("type = ?", "all").Find(&allThemes).Error
	if err != nil {
		return nil, err
	}

	for _, allTheme := range allThemes {
		if !containThemeID(themePermissionData.ThemeIDs, allTheme.ID) {
			themes = append(themes, allTheme)
		}
	}

	return themes, nil
}

func AddDefaultThemePermission(studentID string) error {
	var existingPermission models.ThemePermission
	err := database.DB.Where("student_id = ?", studentID).First(&existingPermission).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			themePermissionData := models.ThemePermissionData{
				ThemeIDs: []int{},
			}
			permission, err := json.Marshal(themePermissionData)
			if err != nil {
				return err
			}

			var defaultThemeID int
			defaultThemeIDStr := config.GetDefaultThemeKey()
			if defaultThemeIDStr != "" {
				defaultThemeID, err = strconv.Atoi(defaultThemeIDStr)
				if err != nil {
					return err
				}
			} else {
				var theme models.Theme
				if err := database.DB.Model(models.Theme{}).Where("type = ?", "all").First(&theme).Error; err != nil {
					return err
				}
				defaultThemeID = theme.ID
			}

			newPermission := models.ThemePermission{
				StudentID:       studentID,
				CurrentThemeID:  defaultThemeID,
				ThemePermission: string(permission),
			}

			result := database.DB.Create(&newPermission)
			return result.Error
		} else {
			return err
		}
	} else {
		return nil
	}
}

func containThemeID(themeIDs []int, id int) bool {
	for _, themeID := range themeIDs {
		if themeID == id {
			return true
		}
	}
	return false
}

const batchSize = 100 // 每次保存 100 条记录

func savePermissionsInBatches(permissions []models.ThemePermission) error {
	totalPermissions := len(permissions)
	for i := 0; i < totalPermissions; i += batchSize {
		end := i + batchSize
		if end > totalPermissions {
			end = totalPermissions
		}

		// 保存当前批次
		err := database.DB.Save(permissions[i:end]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
