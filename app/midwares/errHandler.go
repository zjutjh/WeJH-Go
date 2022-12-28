package midwares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println(c.Errors)
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				var Err *apiException.Error
				if e, ok := err.(*apiException.Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = apiException.OtherError(e.Error())
				} else {
					Err = apiException.ServerError
				}
				// TODO 建立日志系统

				c.JSON(Err.StatusCode, Err)
				return
			}
		}

	}
}

// HandleNotFound
//  404处理
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	c.JSON(err.StatusCode, err)
	return
}
