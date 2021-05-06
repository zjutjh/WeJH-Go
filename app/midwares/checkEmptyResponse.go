package midwares

import (
	"github.com/gin-gonic/gin"
)

func CheckEmptyResponse(c *gin.Context) {
	c.Next()
	if c.Writer.Size() == 0 {

	}
}
