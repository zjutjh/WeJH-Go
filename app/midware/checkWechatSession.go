package midware

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

func CheckWechatSession(c *gin.Context) {
	_, err := sessionServices.GetWechatSession(c)
	if err != nil {
		utils.JsonFailedResponse(c, stateCode.GetOpenIDFail, nil)
		c.Abort()
		return
	}
	c.Next()
}
