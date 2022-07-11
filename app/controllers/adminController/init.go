package adminController

import (
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/apiExpection"
	"wejh-go/app/config"
	"wejh-go/app/utils"
)

type termInfoForm struct {
	yearValue          string    `json:"yearValue"`
	termValue          string    `json:"termValue"`
	termStartDateValue time.Time `json:"termStartDateValue"`
}

type encryptForm struct {
	encryptKey string `json:"encryptKey"`
}

func SetInit(c *gin.Context) {
	err := config.SetInit()
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	utils.JsonSuccessResponse(c, nil)
}

func ResetInit(c *gin.Context) {
	err := config.ResetInit()
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}
	if config.IsSetEncryptKey() {
		err = config.DelEncryptKey()
		if err != nil {
			_ = c.AbortWithError(200, apiExpection.ServerError)
		}
	}
	if config.IsSetTermInfo() {
		errors := config.DelTermInfo()
		for _, err := range errors {
			if err != nil {
				_ = c.AbortWithError(200, apiExpection.ServerError)
			}
		}
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetTermInfo(c *gin.Context) {
	var postForm termInfoForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	err = config.SetTermInfo(postForm.yearValue, postForm.termValue, postForm.termStartDateValue)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetEncryptKey(c *gin.Context) {
	var postForm encryptForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	err = config.SetEncryptKey(postForm.encryptKey)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	utils.JsonSuccessResponse(c, nil)
}
