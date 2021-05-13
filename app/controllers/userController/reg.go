package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/config/wechat"
)

type createStudentUserForm struct {
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	StudentID    string `json:"studentID"  binding:"required"`
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	LoginType    string `json:"type"`
}
type createStudentUserWechatForm struct {
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	StudentID    string `json:"studentID"  binding:"required"`
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	Code         string `json:"code"  binding:"required"`
	LoginType    string `json:"type"`
}

func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)

	if err != nil {
		_ = c.AbortWithError(200, apiExpection.OpenIDError)
		return
	}

	user, err := userServices.CreateStudentUserWechat(
		postForm.Username,
		postForm.Password,
		postForm.StudentID,
		postForm.IDCardNumber,
		postForm.Email,
		session.OpenID)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.UserAlreadyExisted)
		return
	}
	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func CreateStudentUser(c *gin.Context) {
	var postForm createStudentUserForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	user, err := userServices.CreateStudentUser(
		postForm.Username,
		postForm.Password,
		postForm.StudentID,
		postForm.IDCardNumber,
		postForm.Email)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.UserAlreadyExisted)
		return
	}
	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
