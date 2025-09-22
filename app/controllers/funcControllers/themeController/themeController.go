package themeController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/themeServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetThemeList(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	themePermission, err := themeServices.GetThemePermissionByStudentID(user.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	themes, err := themeServices.GetPermittedThemesFormat(user.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"theme_list":            themes,
		"current_theme_id":      themePermission.CurrentThemeID,
		"current_theme_dark_id": themePermission.CurrentThemeDarkID,
	})
}

type ChooseCurrentThemeData struct {
	ID     int `json:"id"`
	DarkID int `json:"dark_id"`
}

func ChooseCurrentTheme(c *gin.Context) {
	var data ChooseCurrentThemeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	err = themeServices.CheckThemeExist(data.ID, data.DarkID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	err = themeServices.UpdateCurrentTheme(data.ID, data.DarkID, user.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
