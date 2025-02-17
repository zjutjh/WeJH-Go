package themeServices

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
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
			newPermission, err := AddDefaultThemePermission(studentID)
			if err != nil {
				return nil, err
			}
			permissions = append(permissions, newPermission)
			permission = &permissions[len(permissions)-1]
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

	err = database.DB.Save(&permissions).Error
	if err != nil {
		return nil, err
	}
	return invalidStudentIDs, nil
}

func UpdateCurrentTheme(id int, darkID int, studentID string) error {
	_, themeIDs, err := getPersonalAllThemePermission(studentID)
	if err != nil {
		return err
	}

	if !containThemeID(themeIDs, id) {
		return errors.New("the light theme ID is not in the user's permission list")
	}
	if !containThemeID(themeIDs, darkID) {
		return errors.New("the dark theme ID is not in the user's permission list")
	}

	result := database.DB.Model(&models.ThemePermission{}).
		Where("student_id = ?", studentID).
		Updates(map[string]interface{}{
			"current_theme_id":      id,
			"current_theme_dark_id": darkID,
		})
	return result.Error
}

func GetThemePermissionByStudentID(studentID string) (models.ThemePermission, error) {
	var record models.ThemePermission
	result := database.DB.Model(&models.ThemePermission{}).Where("student_id = ?", studentID).First(&record)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			newPermission, err := AddDefaultThemePermission(studentID)
			if err != nil {
				return models.ThemePermission{}, err
			}
			return newPermission, nil
		} else {
			return models.ThemePermission{}, result.Error
		}
	}
	return record, nil
}

func GetThemeNameByStudentID(studentID string) ([]string, error) {
	themes, _, err := getPersonalAllThemePermission(studentID)
	if err != nil {
		return nil, err
	}

	var themeNames []string
	for _, theme := range themes {
		themeNames = append(themeNames, theme.Name)
	}
	return themeNames, nil
}

func GetThemesByStudentID(studentID string) ([]models.FormattedTheme, error) {
	themes, _, err := getPersonalAllThemePermission(studentID)
	if err != nil {
		return nil, err
	}

	var parsedThemes []models.FormattedTheme
	for _, theme := range themes {
		var themeConfig models.ThemeConfigData
		if err := json.Unmarshal([]byte(theme.ThemeConfig), &themeConfig); err != nil {
			return nil, err
		}

		parsedTheme := models.FormattedTheme{
			Name:        theme.Name,
			ThemeID:     theme.ID,
			ThemeConfig: themeConfig,
			IsDarkMode:  theme.IsDarkMode,
		}
		parsedThemes = append(parsedThemes, parsedTheme)
	}

	return parsedThemes, nil
}

func AddDefaultThemePermission(studentID string) (models.ThemePermission, error) {
	var existingPermission models.ThemePermission
	err := database.DB.Where("student_id = ?", studentID).First(&existingPermission).Error
	if err == nil {
		return existingPermission, nil
	}
	if err != gorm.ErrRecordNotFound {
		return models.ThemePermission{}, err
	}

	themePermissionData := models.ThemePermissionData{
		ThemeIDs: []int{},
	}
	permission, err := json.Marshal(themePermissionData)
	if err != nil {
		return models.ThemePermission{}, err
	}

	var defaultThemeLightID, defaultThemeDarkID int
	if err := database.DB.Model(models.Theme{}).Where("type= ? AND is_dark_mode = ?", "all", false).Select("id").First(&defaultThemeLightID).Error; err != nil {
		return models.ThemePermission{}, err
	}
	if err := database.DB.Model(models.Theme{}).Where("type= ? AND is_dark_mode = ?", "all", true).Select("id").First(&defaultThemeDarkID).Error; err != nil {
		return models.ThemePermission{}, err
	}

	newPermission := models.ThemePermission{
		StudentID:          studentID,
		CurrentThemeID:     defaultThemeLightID,
		CurrentThemeDarkID: defaultThemeDarkID,
		ThemePermission:    string(permission),
	}

	result := database.DB.Create(&newPermission)
	return newPermission, result.Error
}

func containThemeID(themeIDs []int, id int) bool {
	for _, themeID := range themeIDs {
		if themeID == id {
			return true
		}
	}
	return false
}

func getPersonalAllThemePermission(studentID string) ([]models.Theme, []int, error) {
	var themePermission models.ThemePermission
	var themePermissionData models.ThemePermissionData

	result := database.DB.Model(models.ThemePermission{}).Where("student_id = ?", studentID).First(&themePermission)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			themePermissionData = models.ThemePermissionData{ThemeIDs: []int{}}
			_, err := AddDefaultThemePermission(studentID)
			if err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, result.Error
		}
	} else {
		err := json.Unmarshal([]byte(themePermission.ThemePermission), &themePermissionData)
		if err != nil {
			return nil, nil, err
		}
	}

	var themes []models.Theme
	err := database.DB.Where("id IN ?", themePermissionData.ThemeIDs).Find(&themes).Error
	if err != nil {
		return nil, nil, err
	}

	var allThemes []models.Theme
	err = database.DB.Where("type = ?", "all").Find(&allThemes).Error
	if err != nil {
		return nil, nil, err
	}

	themeIDs := make(map[int]bool)
	for _, id := range themePermissionData.ThemeIDs {
		themeIDs[id] = true
	}

	for _, allTheme := range allThemes {
		if !themeIDs[allTheme.ID] {
			themes = append(themes, allTheme)
			themePermissionData.ThemeIDs = append(themePermissionData.ThemeIDs, allTheme.ID)
		}
	}

	return themes, themePermissionData.ThemeIDs, nil
}
