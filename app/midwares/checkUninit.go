package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/config"
	"wejh-go/app/utils"
	"wejh-go/app/utils/stateCode"
)

func CheckUninit(c *gin.Context) {
	inited := config.GetInit()
	if inited {
		utils.JsonFailedResponse(c, stateCode.NotInit, nil)
		c.Abort()
		return
	}
	c.Next()
}
