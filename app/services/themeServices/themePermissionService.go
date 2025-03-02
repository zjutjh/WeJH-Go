package themeServices

import (
	"encoding/json"
	"errors"
	"wejh-go/app/models"
	"wejh-go/config/database"

	"slices"

	"gorm.io/gorm"
)

func AddThemePermission(themeID int, reqStudentIDs []string, themeType string) ([]string, error) {
	if themeType == "all" {
		return nil, nil
	}

	var studentIDs []string
	var invalidStudentIDs []string
	if len(reqStudentIDs) > 0 {
		var existingUsers []models.User
		if err := database.DB.Select("student_id").
			Where("student_id IN ?", reqStudentIDs).
			Find(&existingUsers).Error; err != nil {
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
	if err := database.DB.Where("student_id IN ?", studentIDs).Find(&permissions).Error; err != nil {
		return nil, err
	}
	permissionMap := make(map[string]*models.ThemePermission)
	for i, p := range permissions {
		permissionMap[p.StudentID] = &permissions[i]
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
		if err := json.Unmarshal([]byte(permission.ThemePermission), &themePermissionData); err != nil {
			return nil, err
		}
		if !slices.Contains(themePermissionData.ThemeIDs, themeID) {
			themePermissionData.ThemeIDs = append(themePermissionData.ThemeIDs, themeID)
			newPermission, err := json.Marshal(themePermissionData)
			if err != nil {
				return nil, err
			}
			permission.ThemePermission = string(newPermission)
		}
	}

	if err := database.DB.Save(&permissions).Error; err != nil {
		return nil, err
	}
	return invalidStudentIDs, nil
}

func UpdateCurrentTheme(id int, darkID int, studentID string) error {
	themeIDs, err := GetPermittedThemeIDs(studentID)
	if err != nil {
		return err
	}

	if !slices.Contains(themeIDs, id) {
		return errors.New("the light theme ID is not in the user's permission list")
	}
	if !slices.Contains(themeIDs, darkID) {
		return errors.New("the dark theme ID is not in the user's permission list")
	}

	err = database.DB.Model(&models.ThemePermission{}).
		Where("student_id = ?", studentID).
		Updates(models.ThemePermission{
			CurrentThemeID:     id,
			CurrentThemeDarkID: darkID,
		}).Error
	return err
}

func GetThemePermissionByStudentID(studentID string) (models.ThemePermission, error) {
	var permission models.ThemePermission
	result := database.DB.Model(&models.ThemePermission{}).Where("student_id = ?", studentID).First(&permission)
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
	return permission, nil
}

func GetPermittedThemesFormat(studentID string) ([]models.FormatTheme, error) {
	themes, err := getPermittedThemes(studentID)
	if err != nil {
		return nil, err
	}

	var formatThemes []models.FormatTheme
	for _, theme := range themes {
		var themeConfig models.ThemeConfig
		if err := json.Unmarshal([]byte(theme.ThemeConfig), &themeConfig); err != nil {
			return nil, err
		}
		formatThemes = append(formatThemes, models.FormatTheme{
			Name:        theme.Name,
			ThemeID:     theme.ID,
			ThemeConfig: themeConfig,
			IsDarkMode:  theme.IsDarkMode,
		})
	}

	return formatThemes, nil
}

func GetPermittedThemeIDs(studentID string) ([]int, error) {
	themes, err := getPermittedThemes(studentID)
	if err != nil {
		return nil, err
	}

	var themeIDs []int
	for _, theme := range themes {
		themeIDs = append(themeIDs, theme.ID)
	}
	return themeIDs, nil
}

func GetPermittedThemeNames(studentID string) ([]string, error) {
	themes, err := getPermittedThemes(studentID)
	if err != nil {
		return nil, err
	}

	var themeNames []string
	for _, theme := range themes {
		themeNames = append(themeNames, theme.Name)
	}
	return themeNames, nil
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
	if err := database.DB.Model(models.Theme{}).
		Select("id").
		Where("type = all AND is_dark_mode = false").
		First(&defaultThemeLightID).Error; err != nil {
		return models.ThemePermission{}, err
	}
	if err := database.DB.Model(models.Theme{}).
		Select("id").
		Where("type = all AND is_dark_mode = true").
		First(&defaultThemeDarkID).Error; err != nil {
		return models.ThemePermission{}, err
	}

	newPermission := models.ThemePermission{
		StudentID:          studentID,
		CurrentThemeID:     defaultThemeLightID,
		CurrentThemeDarkID: defaultThemeDarkID,
		ThemePermission:    string(permission),
	}

	err = database.DB.Create(&newPermission).Error
	return newPermission, err
}

func getPermittedThemes(studentID string) ([]models.Theme, error) {
	var themePermission models.ThemePermission
	var themePermissionData models.ThemePermissionData

	if err := database.DB.Where("student_id = ?", studentID).First(&themePermission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			_, err := AddDefaultThemePermission(studentID)
			if err != nil {
				return nil, err
			}
			themePermissionData = models.ThemePermissionData{
				ThemeIDs: []int{},
			}
		} else {
			return nil, err
		}
	} else {
		if err := json.Unmarshal([]byte(themePermission.ThemePermission), &themePermissionData); err != nil {
			return nil, err
		}
	}

	var themes []models.Theme
	query := database.DB.Where("type = all")
	if len(themePermissionData.ThemeIDs) > 0 {
		query = query.Or("id IN ?", themePermissionData.ThemeIDs)
	}
	if err := query.Find(&themes).Error; err != nil {
		return nil, err
	}

	return themes, nil
}
