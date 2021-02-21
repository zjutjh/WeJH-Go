package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"wejh-go/conf"
	"wejh-go/database"
)

type bindJHForm struct {
	OpenID    string `json:"openid"`
	PassWord  string `json:"password"`
	UserName  string `json:"username"`
	LoginType string `json:"type"`
}

type autoLoginForm struct {
	OpenID    string `json:"openid"`
	LoginType string `json:"type"`
}

type WeAppForm struct {
	Code string `json:"code"`
}

type WeAppResultBody struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	ErrMSG     string `json:"errmsg"`
}

func BindJHControllers(c *gin.Context) {
	var postForm bindJHForm
	err := c.ShouldBindJSON(&postForm) // 获取 POST 请求的 JSON 数据 要用指针类型作为参数
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"绑定错信息错误": err.Error()})
		return
	}

	// TODO: 未来在这里加上判断用户的输入的账号是否在用户中心注册过的功能

	user := database.User{
		Uno:    postForm.UserName,
		Pass:   postForm.PassWord, // TODO: 添加加密存储
		OpenID: postForm.OpenID,
	} // 用获取到的数据生成数据库模型
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{ // 服务器返回用户信息
		"errcode":  200,
		"errmsg":   "绑定成功",
		"redirect": nil,
		"data": gin.H{
			"token": postForm.OpenID,
			"user": gin.H{
				"uno": postForm.UserName, // TODO: 添加其他账号的绑定信息
			},
		},
	})
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

	user := database.User{}
	result := database.DB.Where("open_id = ?", postForm.OpenID).First(&user)
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

func WeAppController(c *gin.Context) {
	// 从用户输入中获取小程序 code
	var postForm WeAppForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"参数错误":     err.Error(),
			"data":     nil,
			"err_code": -403,
			"err_msg":  nil,
		})
		return
	}

	// 获取用户的 openID
	code := postForm.Code
	appID := conf.Config.GetString("weapp.appid")
	secret := conf.Config.GetString("weapp.secret")
	response, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code",
		appID,
		secret,
		code,
	)) // 发起请求，从腾讯服务器那里接收 openID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg": "小程序 openID 获取失败",
		})
		return
	}
	body, _ := ioutil.ReadAll(response.Body) // 读取响应内容
	var bodyJSON WeAppResultBody
	err = json.Unmarshal(body, &bodyJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_code": -2,
			"err_msg":  "腾讯服务器返回错误，请检查小程序相关配置是否正确",
			"redirect": nil,
		})
	} else { // 返回用户 openID
		c.JSON(http.StatusOK, gin.H{
			"err_code": 200,
			"err_msg":  "获取 openID 成功",
			"data": gin.H{
				"openid": bodyJSON.OpenID,
			},
		})
	}
	_ = response.Body.Close()
}
