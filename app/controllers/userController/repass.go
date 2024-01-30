package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

type RePassForm struct {
	IDCard    string `json:"iid" binding:"required"`
	StudentID string `json:"stuid" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func ResetPass(c *gin.Context) {
	var postForm RePassForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	if user.Username != postForm.StudentID {
		_ = c.AbortWithError(200, apiException.StudentIdError)
		return
	}

	if err = userServices.ResetPass(user, postForm.IDCard, postForm.Password); err != nil {
		if err == apiException.StudentNumAndIidError {
			_ = c.AbortWithError(200, apiException.StudentNumAndIidError)
			return
		} else if err == apiException.PwdError {
			_ = c.AbortWithError(200, apiException.PwdError)
			return
		}
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
