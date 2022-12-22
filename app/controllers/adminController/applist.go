package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

type createAppListForm struct {
	Title           string `json:"title" binding:"required"`
	Route           string `json:"route" binding:"required"`
	BackgroundColor string `json:"backgroundColor" binding:"required"`
	Icon            string `json:"icon" binding:"required"`
	Require         string `json:"require" binding:"required"`
}

type deleteAppListForm struct {
	ID int64 `json:"id" binding:"required"`
}

func CreateApplist(c *gin.Context) {
	var postForm createAppListForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = applistServices.CreateApplist(models.AppList{
		Title:           postForm.Title,
		Route:           postForm.Route,
		BackgroundColor: postForm.BackgroundColor,
		Icon:            postForm.Icon,
		Require:         postForm.Require,
	})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateApplist(c *gin.Context) {
	var postForm models.AppList
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = applistServices.UpdateApplist(postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
func DeleteApplist(c *gin.Context) {
	var postForm deleteAppListForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = applistServices.DeleteApplist(postForm.ID)
	if err != nil {

		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
