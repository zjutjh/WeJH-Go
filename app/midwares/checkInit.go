package midwares

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
)

func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		apiException.AbortWithException(c, apiException.NotInit, nil)
		return
	}
	c.Next()
}
