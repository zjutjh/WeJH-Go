package midwares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(c.Errors)
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				var Err *apiExpection.Error
				if e, ok := err.(*apiExpection.Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = apiExpection.OtherError(e.Error())
				} else {
					Err = apiExpection.ServerError
				}
				// 记录一个错误的日志

				c.JSON(Err.StatusCode, Err)
				return
			}
		}

	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := apiExpection.NotFound
	c.JSON(err.StatusCode, err)
	return
}
