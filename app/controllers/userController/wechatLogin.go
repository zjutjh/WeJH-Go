package userController

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"wejh-go/config"
)

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
	appID := config.Config.GetString("weapp.appid")
	secret := config.Config.GetString("weapp.secret")
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
