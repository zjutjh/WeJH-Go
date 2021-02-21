package user_controllers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
	"net/http"
	"wejh-go/conf"
	"wejh-go/database"
)

type autoLoginForm struct {
	OpenID    string `json:"openid"`
	LoginType string `json:"type"`
}

func AutoLoginControllers(c *gin.Context) {
	// 读取请求信息
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"错误":       err.Error(),
			"data":     nil,
			"err_code": -403,
			"err_msg":  nil,
		})
		return
	}
	if postForm.OpenID == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"data":     nil,
			"err_code": -401,
			"err_msg":  "缺少用户标识",
		})
	}

	// 从数据中读取绑定信息
	cryptOpenID, _ := goEncrypt.AesCtrEncrypt(
		[]byte(postForm.OpenID),
		[]byte(conf.Config.GetString("encryptKey")),
	)
	user := database.User{}
	result := database.DB.Where(
		"open_id = ?",
		base64.StdEncoding.EncodeToString(cryptOpenID),
	).First(&user)

	if result.RowsAffected <= 0 { // 没有找到对应用户
		c.JSON(http.StatusForbidden, gin.H{
			"data":     nil,
			"err_code": -403,
			"err_msg":  "自动登陆失败",
		})
	} else { // 返回用户数据
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token": postForm.OpenID,
				"user": gin.H{
					"uno": user.Uno, // TODO: 添加其他账号的绑定信息
				},
			},
			"errcode":  200,
			"errmsg":   "登陆成功",
			"redirect": nil,
		})
	}
}
