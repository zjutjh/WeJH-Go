package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/applistServices"
	"wejh-go/app/utils"
)

type createAppListForm struct {
	Title string `json:"title" binding:"required"`
	Route string `json:"route"`
}
type updateAppListForm struct {
	ID    int64  `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Route string `json:"route"`
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

	err = applistServices.CreateApplist(models.AppList{Title: postForm.Title, Route: postForm.Route})
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateApplist(c *gin.Context) {
	var postForm updateAppListForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = applistServices.UpdateApplist(postForm.ID, models.AppList{
		Title: postForm.Title,
		Route: postForm.Route,
	},
	)
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
