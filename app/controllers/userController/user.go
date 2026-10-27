package userController

import (
	"errors"
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
	// 1. 获取当前登录用户
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	// 2. 参数绑定与校验
	var req ThemeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 3. 调用 Service 层更新数据库
	if err := userServices.UpdateUserThemePreference(uint(user.ID), req.Theme); err != nil {
		if errors.Is(err, userServices.ErrInvalidTheme) {
			apiException.AbortWithException(c, apiException.ParamError, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	// 4. 统一成功返回
	utils.JsonSuccessResponse(c, gin.H{
		"themePreference": req.Theme,
	})
}
