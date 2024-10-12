package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"
)

type CreateThemeData struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	ThemeConfig string `json:"theme_config"`
}

// 管理员创建主题色
func CreateTheme(c *gin.Context) {
	var data CreateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	record := models.Theme{
		Name:        data.Name,
		Type:        data.Type,
		ThemeConfig: data.ThemeConfig,
	}
	themeID, err := themeServices.CreateTheme(record)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if data.Type == "all" {
		studentIDs, err := themeServices.GetAllStudentIDs()
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}

		_, err = themeServices.AddThemePermission(themeID, studentIDs)
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
}

type UpdateThemeData struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ThemeConfig string `json:"theme_config" binding:"required"`
}

// 管理员更新主题色
func UpdateTheme(c *gin.Context) {
	var data UpdateThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	record := models.Theme{
		Name:        data.Name,
		ThemeConfig: data.ThemeConfig,
	}
	err = themeServices.UpdateTheme(data.ID, record)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

// 管理员获取主题色列表
func GetThemes(c *gin.Context) {
	themes, err := themeServices.GetThemes()
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

	err = themeServices.DeleteTheme(data.ID)
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

	invalidStudentIDs, err := themeServices.AddThemePermission(data.ThemeID, data.StudentID)
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

	themePermission, err := themeServices.GetThemePermissionByStudentID(data.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	themeNames, err := themeServices.GetThemeNameByID(themePermission)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_name": themeNames,
	})
}
