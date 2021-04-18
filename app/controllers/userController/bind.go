package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/config/database"
)

type bindForm struct {
	PassWord string `json:"password"`
}

func BindZFPassword(c *gin.Context) {
	var postForm bindForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	user.ZFPassword = postForm.PassWord

	database.DB.Save(user)
}
func BindLibraryPassword(c *gin.Context) {
	var postForm bindForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	user.LibPassword = postForm.PassWord

	database.DB.Save(user)
}

func BindSchoolCardPassword(c *gin.Context) {
	var postForm bindForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	user.CardPassword = postForm.PassWord

	database.DB.Save(user)
}
