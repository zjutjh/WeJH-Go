package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
	"wejh-go/app/config"
)

func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		_ = c.AbortWithError(200, apiExpection.NotInit)
		return
	}
	c.Next()
}
