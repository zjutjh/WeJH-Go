package userController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

type DelForm struct {
	IDCard    string `json:"iid" binding:"required"`
	StudentID string `json:"stuid" binding:"required"`
}

func DelAccount(c *gin.Context) {
	var postForm DelForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	user, err := sessionServices.GetUserSession(c)

	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	if user.Username != postForm.StudentID {
		apiException.AbortWithException(c, apiException.StudentIdError, nil)
		return
	}

	if err = userServices.DelAccount(user, postForm.IDCard); err != nil {
		apiException.AbortWithError(c, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
