package userController

import (
	"net/http"

	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

// 定义请求结构体
type ThemeRequest struct {
	Theme string `json:"theme" binding:"required,oneof=light dark auto"`
}

func SetThemePreference(c *gin.Context) {
	//获取用户
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	// 限制请求体大小，防止恶意请求
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1024)

	//参数绑定与校验
	var req ThemeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	//调用 Service 层
	if err := userServices.UpdateUserThemePreference(uint(user.ID), req.Theme); err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	//统一成功返回
	utils.JsonSuccessResponse(c, gin.H{
		"themePreference": req.Theme,
	})
}
