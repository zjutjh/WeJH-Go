package userController

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
	"net/http"
	"wejh-go/app/models"
	"wejh-go/app/utils"
	"wejh-go/config"
	"wejh-go/service/database"
)

// TODO: 未来加入对精弘通行证的去除重复功能

type bindJHForm struct {
	OpenID    string `json:"openid"`
	PassWord  string `json:"password"`
	UserName  string `json:"username"`
	LoginType string `json:"type"`
}

func BindJHControllers(c *gin.Context) {
	var postForm bindJHForm
	err := c.ShouldBindJSON(&postForm) // 获取 POST 请求的 JSON 数据 要用指针类型作为参数
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"绑定错信息错误": err.Error()})
		return
	}

	// TODO: 未来在这里加上判断用户的输入的账号是否在用户中心注册过的功能

	// 对密码和 openID 等关键数据进行加密
	cryptPass, _ := goEncrypt.AesCtrEncrypt(
		[]byte(postForm.PassWord),
		[]byte(config.Config.GetString("encryptKey")),
	)
	cryptOpenID, _ := goEncrypt.AesCtrEncrypt(
		[]byte(postForm.OpenID),
		[]byte(config.Config.GetString("encryptKey")),
	)

	user := models.User{
		StudentID: postForm.UserName,
		Password:  base64.StdEncoding.EncodeToString(cryptPass),
		OpenID:    base64.StdEncoding.EncodeToString(cryptOpenID),
	} // 用获取到的数据生成数据库模型
	database.DB.Create(&user)

	utils.JsonSuccessResponse(c, gin.H{
		"token": postForm.OpenID,
		"user": gin.H{
			"uno": postForm.UserName, // TODO: 添加其他账号的绑定信息
		}})
}
