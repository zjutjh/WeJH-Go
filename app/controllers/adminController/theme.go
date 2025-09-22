package adminController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateThemeData struct {
	ThemeName   string             `json:"theme_name" binding:"required"`
	ThemeType   string             `json:"theme_type" binding:"required"`
	IsDarkMode  bool               `json:"is_dark_mode"`
	ThemeConfig models.ThemeConfig `json:"theme_config" binding:"required"`
}

// 管理员创建主题色
func CreateTheme(c *gin.Context) {
	var data CreateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = themeServices.CreateTheme(data.ThemeName, data.ThemeType, data.IsDarkMode, data.ThemeConfig)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type UpdateThemeData struct {
	ThemeID     int                `json:"theme_id" binding:"required"`
	ThemeName   string             `json:"theme_name" binding:"required"`
	IsDarkMode  bool               `json:"is_dark_mode"`
	ThemeConfig models.ThemeConfig `json:"theme_config" binding:"required"`
}

// 管理员更新主题色
func UpdateTheme(c *gin.Context) {
	var data UpdateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = themeServices.UpdateTheme(data.ThemeID, data.ThemeName, data.IsDarkMode, data.ThemeConfig)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

// 管理员获取主题色列表
func GetAllTheme(c *gin.Context) {
	themes, err := themeServices.GetAllTheme()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
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
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = themeServices.CheckThemeExist(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	theme, err := themeServices.GetThemeByID(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	err = themeServices.DeleteTheme(data.ID, theme.Type, theme.IsDarkMode)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
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
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = themeServices.CheckThemeExist(data.ThemeID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	theme, err := themeServices.GetThemeByID(data.ThemeID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	invalidStudentIDs, err := themeServices.AddThemePermission(data.ThemeID, data.StudentID, theme.Type)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
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
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	themeNames, err := themeServices.GetPermittedThemeNames(data.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_name": themeNames,
	})
}
