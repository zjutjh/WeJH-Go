package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wejh-go/app/utils/stateCode"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data any) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// JsonSuccessResponse 返回成功json格式数据
func JsonSuccessResponse(c *gin.Context, data any) {
	JsonResponse(c, http.StatusOK, stateCode.OK, "OK", data)
}

// JsonErrorResponse 返回错误json格式数据
func JsonErrorResponse(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusOK, code, msg, nil)
}
