package midwares

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"

	"github.com/gin-gonic/gin"
)

func CheckLogin(c *gin.Context) {
	isLogin := sessionServices.CheckUserSession(c)
	if !isLogin {
		apiException.AbortWithException(c, apiException.NotLogin, nil)
		return
	}
	c.Next()
}
