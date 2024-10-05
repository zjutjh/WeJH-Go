package themeController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"
)

func GetThemeList(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	themePermission, err := themeServices.GetThemePermissionByStudentID(user.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	themes, err := themeServices.GetThemesByID(themePermission)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_list":       themes,
		"current_theme_id": themePermission.CurrentThemeID,
	})
}

type ChooseCurrentThemeData struct {
	ID int `json:"id" binding:"required"`
}

func ChooseCurrentTheme(c *gin.Context) {
	var data ChooseCurrentThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	err = themeServices.CheckThemeExist(data.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = themeServices.UpdateCurrentTheme(data.ID, user.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
