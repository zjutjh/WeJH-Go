package userController

import (
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/services/funnelServices"
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

	_, err = funnelServices.GetExam(user, string(rune(time.Now().Year())), "3")
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	database.DB.Save(user)
	utils.JsonSuccessResponse(c, nil)

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

	_, err = funnelServices.GetCurrentBorrow(user)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	database.DB.Save(user)
	utils.JsonSuccessResponse(c, nil)
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
	_, err = funnelServices.GetCardBalance(user)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	database.DB.Save(user)
	utils.JsonSuccessResponse(c, nil)
}
