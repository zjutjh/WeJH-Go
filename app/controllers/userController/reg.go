package userController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/themeServices"
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
	Type         uint   `json:"type"  ` // 用户类型 0-本科生 1-研究生
}
type createStudentUserWechatForm struct {
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	StudentID    string `json:"studentID"  binding:"required"`
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
	Code         string `json:"code"  binding:"required"`
	Type         uint   `json:"type" ` // 用户类型 0-本科生 1-研究生
}

func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)

	if err != nil {
		apiException.AbortWithException(c, apiException.OpenIDError, err)
		return
	}
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.Username = strings.ToUpper(postForm.Username)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userServices.CreateStudentUserWechat(
		postForm.Username,
		postForm.Password,
		postForm.StudentID,
		postForm.IDCardNumber,
		postForm.Email,
		session.OpenID,
		postForm.Type)
	if err != nil && !errors.Is(err, apiException.ReactiveError) {
		apiException.AbortWithError(c, err)
		return
	}

	_, err = themeServices.AddDefaultThemePermission(postForm.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func CreateStudentUser(c *gin.Context) {
	var postForm createStudentUserForm
	errBind := c.ShouldBindJSON(&postForm)
	if errBind != nil {
		apiException.AbortWithException(c, apiException.ParamError, errBind)
		return
	}
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.Username = strings.ToUpper(postForm.Username)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userServices.CreateStudentUser(
		postForm.Username,
		postForm.Password,
		postForm.StudentID,
		postForm.IDCardNumber,
		postForm.Email,
		postForm.Type)
	if err != nil && !errors.Is(err, apiException.ReactiveError) {
		apiException.AbortWithError(c, err)
		return
	}

	_, err = themeServices.AddDefaultThemePermission(postForm.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	err = sessionServices.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
