package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/app/utils/stateCode"
)

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": stateCode.OK,
		"msg":  "OK",
	})
}

func JsonResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"code": code,
		"msg":  msg,
	})
}
