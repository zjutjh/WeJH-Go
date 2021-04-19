package userController

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config/database"
	"wejh-go/config/wechat"
)

type createStudentUserForm struct {
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	StudentID    string `json:"studentID"  binding:"required"`
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	LoginType    string `json:"type"`
}
type createStudentUserWechatForm struct {
	Username     string `json:"username"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	StudentID    string `json:"studentID"  binding:"required"`
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Code         string `json:"code"  binding:"required"`
	LoginType    string `json:"type"`
}

func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)

	if err != nil {
		utils.JsonFailedResponse(c, stateCode.GetOpenIDFail, nil)
		return
	}

	err = userCenterServices.AuthStudent(postForm.StudentID, postForm.IDCardNumber)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UserNotFind, nil)
		return
	}

	user, err := userServices.GetUserByOpenID(session.OpenID)
	if err != nil && user != nil {
		utils.JsonFailedResponse(c, stateCode.UserAlreadyExisted, nil)
		return
	}

	user, err = userServices.GetUserByUsernameAndPassword(postForm.Username, postForm.Password)
	if user != nil {
		user.WechatOpenID = session.OpenID
		database.DB.Save(user)
		utils.JsonSuccessResponse(c, nil)
		return
	}
	h := sha256.New()
	h.Write([]byte(postForm.Password))
	pass := hex.EncodeToString(h.Sum(nil))
	user = &models.User{WechatOpenID: session.OpenID, JHPassword: pass, Username: postForm.Username, Type: models.Undergraduate, StudentID: postForm.StudentID, CreateTime: time.Now()}
	database.DB.Create(&user)
	utils.JsonSuccessResponse(c, nil)
}

func CreateStudentUser(c *gin.Context) {
	var postForm createStudentUserForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}
	user, err := userServices.GetUserByUsername(postForm.Username)
	if user != nil {
		utils.JsonFailedResponse(c, stateCode.UserAlreadyExisted, nil)
		return
	}

	err = userCenterServices.AuthStudent(postForm.StudentID, postForm.IDCardNumber)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.UserNotFind, nil)
		return
	}

	h := sha256.New()
	h.Write([]byte(postForm.Password))
	pass := hex.EncodeToString(h.Sum(nil))
	user = &models.User{JHPassword: pass, Username: postForm.Username, Type: models.Undergraduate, StudentID: postForm.StudentID, CreateTime: time.Now()}
	database.DB.Create(&user)
	sessionServices.SetUserSession(c, user)
	utils.JsonSuccessResponse(c, nil)
}
