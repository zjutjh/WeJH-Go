package midwares

import (
	"errors"
	"wejh-go/app/apiException"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zjutjh/mygo/nlog"
)

// ErrHandler 中间件用于处理请求错误。
// 如果存在错误，将返回相应的 JSON 响应。
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 向下执行请求
		c.Next()

		// 如果存在错误，则处理错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if err != nil {
				var apiErr *apiException.Error

				// 尝试将错误转换为 apiException
				ok := errors.As(err, &apiErr)

				// 如果转换失败，则使用 ServerError
				if !ok {
					apiErr = apiException.ServerError
					nlog.Pick().WithContext(c).WithError(err).Error("Unknown Error Occurred")
				}
				utils.JsonErrorResponse(c, apiErr.Code, apiErr.Msg)
				return
			}
		}
	}
}

// HandleNotFound 404处理
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	// 记录 404 错误日志
	nlog.Pick().WithContext(c).WithError(err).WithFields(logrus.Fields{
		"path":   c.Request.URL.Path,
		"method": c.Request.Method,
	}).Warn("404 Not Found")
	utils.JsonErrorResponse(c, err.Code, err.Msg)
}
