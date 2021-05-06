package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

func CheckLogin(c *gin.Context) {
	isLogin := sessionServices.CheckUserSession(c)
	if !isLogin {
		utils.JsonFailedResponse(c, stateCode.NotLogin, nil)
		c.Abort()
		return
	}
	c.Next()
}
