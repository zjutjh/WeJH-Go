package customizeHomeController

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/customizeHomeServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type updateForm struct {
	Content []string `json:"content" binging:"required"`
}

func GetCustomizeHome(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	home, err := customizeHomeServices.GetCustomizeHome(user.Username)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	homes := strings.Split(home.Content, ";")

	utils.JsonSuccessResponse(c, homes)
}

func UpdateCustomizeHome(c *gin.Context) {
	var postForm updateForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	content := strings.Join(postForm.Content, ";")
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	flag := false
	home, err := customizeHomeServices.GetCustomizeHome(user.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		flag = true
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if flag {
		err := customizeHomeServices.CreateCustomizeHome(models.CustomizeHome{
			Username: user.Username,
			Content:  content,
		})
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	} else {
		err := customizeHomeServices.UpdateCustomizeHome(home.ID, models.CustomizeHome{
			Username: user.Username,
			Content:  content,
		})
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}
	utils.JsonSuccessResponse(c, nil)
}
