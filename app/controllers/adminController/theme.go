package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"
)

type CreateThemeData struct {
	ThemeName   string                 `json:"theme_name" binding:"required"`
	ThemeType   string                 `json:"theme_type" binding:"required"`
	IsDarkMode  bool                   `json:"is_dark_mode"`
	ThemeConfig models.ThemeConfigData `json:"theme_config" binding:"required"`
}

// 管理员创建主题色
func CreateTheme(c *gin.Context) {
	var data CreateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = themeServices.CreateTheme(data.ThemeName, data.ThemeType, data.IsDarkMode, data.ThemeConfig)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type UpdateThemeData struct {
	ThemeID     int                    `json:"theme_id" binding:"required"`
	ThemeName   string                 `json:"theme_name" binding:"required"`
	IsDarkMode  bool                   `json:"is_dark_mode"`
	ThemeConfig models.ThemeConfigData `json:"theme_config" binding:"required"`
}

// 管理员更新主题色
func UpdateTheme(c *gin.Context) {
	var data UpdateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = themeServices.UpdateTheme(data.ThemeID, data.ThemeName, data.IsDarkMode, data.ThemeConfig)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

// 管理员获取主题色列表
func GetAllTheme(c *gin.Context) {
	themes, err := themeServices.GetAllTheme()
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_list": themes,
	})
}

type DeleteThemeData struct {
	ID int `form:"id" binding:"required"`
}

// 管理员根据id删除主题色
func DeleteTheme(c *gin.Context) {
	var data DeleteThemeData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = themeServices.CheckThemeExist(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	themeType, isDarkMode, err := themeServices.GetThemeByID(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = themeServices.DeleteTheme(data.ID, themeType, isDarkMode)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type AddThemePermissionData struct {
	ThemeID   int      `json:"theme_id" binding:"required"`
	StudentID []string `json:"student_id" binding:"required"`
}

// 管理员添加用户主题色权限
func AddThemePermission(c *gin.Context) {
	var data AddThemePermissionData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = themeServices.CheckThemeExist(data.ThemeID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	themeType, _, err := themeServices.GetThemeByID(data.ThemeID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	invalidStudentIDs, err := themeServices.AddThemePermission(data.ThemeID, data.StudentID, themeType)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"invalid_student_ids": invalidStudentIDs,
	})
}

type GetThemePermissionData struct {
	StudentID string `form:"student_id" binding:"required"`
}

// 管理员根据学号查询用户主题色权限
func GetThemePermission(c *gin.Context) {
	var data GetThemePermissionData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	themeNames, err := themeServices.GetThemeNameByStudentID(data.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_name": themeNames,
	})
}
