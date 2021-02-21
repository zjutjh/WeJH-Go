package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/database"
)

type bindJHForm struct {
	OpenID    string `json:"openid"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	LoginType string `json:"type"`
}

func BindJHControllers(c *gin.Context) {
	var postForm bindJHForm
	err := c.ShouldBindJSON(&postForm) // 获取 POST 请求的 JSON 数据 要用指针类型作为参数
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"错误": err.Error()})
		return
	}
	user := database.User{
		Uno:    postForm.Username,
		Pass:   postForm.Password, // 暂时明文存储
		OpenID: postForm.OpenID,
	} // 用获取到的数据生成数据库模型
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"errcode":  200,
		"errmsg":   "绑定成功",
		"redirect": nil,
		"data": gin.H{
			"token": postForm.OpenID,
		},
	})
}
