package userController

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
	"wejh-go/app/models"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
	"wejh-go/config"
	"wejh-go/service/database"
)

type autoLoginForm struct {
	OpenID    string `json:"openid" binding:"required"`
	LoginType string `json:"type"`
}

func AutoLoginControllers(c *gin.Context) {
	// 读取请求信息
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.ParamError, nil)
		return
	}

	// 从数据中读取绑定信息
	cryptOpenID, _ := goEncrypt.AesCtrEncrypt(
		[]byte(postForm.OpenID),
		[]byte(config.Config.GetString("encryptKey")),
	)
	user := models.User{}
	result := database.DB.Where(
		"open_id = ?",
		base64.StdEncoding.EncodeToString(cryptOpenID),
	).First(&user)

	if result.RowsAffected <= 0 { // 没有找到对应用户
		utils.JsonFailedResponse(c, stateCode.UserNotFind, nil)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"token": postForm.OpenID,
		"user": gin.H{
			"studentID": user.StudentID, // TODO: 添加其他账号的绑定信息
		},
	})

}
