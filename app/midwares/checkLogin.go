package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
)

func CheckLogin(c *gin.Context) {
	isLogin := sessionServices.CheckUserSession(c)
	if !isLogin {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	c.Next()
}
