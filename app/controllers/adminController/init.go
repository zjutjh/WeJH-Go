package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/utils"
)

type SystemInfoForm struct {
	YearValue          string `json:"yearValue"`
	TermValue          string `json:"termValue"`
	TermStartDateValue string `json:"termStartDateValue"`
	ScoreYearValue     string `json:"scoreYearValue"`
	ScoreTermValue     string `json:"scoreTermValue"`
	SchoolBusUrlValue  string `json:"schoolBusUrlValue"`
	JpgUrlValue        string `json:"jpgUrlValue"`
	FileUrlValue       string `json:"fileUrlValue"`
	RegisterTips       string `json:"registerTips"`
}

type encryptForm struct {
	EncryptKey string `json:"encryptKey"`
}

func SetInit(c *gin.Context) {
	err := config.SetInit()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
	}

	utils.JsonSuccessResponse(c, nil)
}

func ResetInit(c *gin.Context) {

	if config.IsSetEncryptKey() {
		err := config.DelEncryptKey()
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}
	if config.IsSetTermInfo() {
		errors := config.DelTermInfo()
		for _, err := range errors {
			if err != nil {
				apiException.AbortWithException(c, apiException.ServerError, err)
				return
			}
		}
	}
	err := config.ResetInit()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetSystemInfo(c *gin.Context) {
	var postForm SystemInfoForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = config.SetSystemInfo(postForm.YearValue, postForm.TermValue, postForm.TermStartDateValue, postForm.ScoreYearValue, postForm.ScoreTermValue, postForm.JpgUrlValue, postForm.FileUrlValue, postForm.RegisterTips, postForm.SchoolBusUrlValue)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func SetEncryptKey(c *gin.Context) {
	var postForm encryptForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = config.SetEncryptKey(postForm.EncryptKey)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
