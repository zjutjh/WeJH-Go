package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
	"wejh-go/app/services/sessionServices"
)

func CheckLogin(c *gin.Context) {
	isLogin := sessionServices.CheckUserSession(c)
	if !isLogin {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	c.Next()
}
