package adminController

import (
	"github.com/gin-gonic/gin"
	"log"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/utils"
)

type termInfoForm struct {
	YearValue          string `json:"yearValue"`
	TermValue          string `json:"termValue"`
	TermStartDateValue string `json:"termStartDateValue"`
	ScoreYearValue     string `json:"scoreYearValue"`
	ScoreTermValue     string `json:"scoreTermValue"`
}

type encryptForm struct {
	EncryptKey string `json:"encryptKey"`
}

func SetInit(c *gin.Context) {
	err := config.SetInit()
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
	}

	utils.JsonSuccessResponse(c, nil)
}

func ResetInit(c *gin.Context) {

	if config.IsSetEncryptKey() {
		err := config.DelEncryptKey()
		if err != nil {
			log.Println(err.Error())
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	if config.IsSetTermInfo() {
		errors := config.DelTermInfo()
		for _, err := range errors {
			if err != nil {
				log.Println(err.Error())
				_ = c.AbortWithError(200, apiException.ServerError)
				return
			}
		}
	}
	err := config.ResetInit()
	if err != nil {
		log.Println(err.Error())
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetTermInfo(c *gin.Context) {
	var postForm termInfoForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = config.SetTermInfo(postForm.YearValue, postForm.TermValue, postForm.TermStartDateValue, postForm.ScoreYearValue, postForm.ScoreTermValue)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetEncryptKey(c *gin.Context) {
	var postForm encryptForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	err = config.SetEncryptKey(postForm.EncryptKey)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
