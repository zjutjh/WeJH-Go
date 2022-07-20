package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
)

func CheckUninit(c *gin.Context) {
	inited := config.GetInit()
	if inited {
		_ = c.AbortWithError(200, apiException.NotInit)
		return
	}
	c.Next()
}
