package midwares

import (
	"wejh-go/app/apiException"
	"wejh-go/app/config"

	"github.com/gin-gonic/gin"
)

func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		apiException.AbortWithException(c, apiException.NotInit, nil)
		return
	}
	c.Next()
}
