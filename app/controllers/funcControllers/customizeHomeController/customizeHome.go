package customizeHomeController

import (
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
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	home, err := customizeHomeServices.GetCustomizeHome(user.Username)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	homes := strings.Split(home.Content, ";")

	utils.JsonSuccessResponse(c, homes)
}

func UpdateCustomizeHome(c *gin.Context) {
	var postForm updateForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	content := strings.Join(postForm.Content, ";")
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	flag := false
	home, err := customizeHomeServices.GetCustomizeHome(user.Username)
	if err == gorm.ErrRecordNotFound {
		flag = true
	} else if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	if flag {
		err := customizeHomeServices.CreateCustomizeHome(models.CustomizeHome{
			Username: user.Username,
			Content:  content,
		})
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	} else {
		err := customizeHomeServices.UpdateCustomizeHome(home.ID, models.CustomizeHome{
			Username: user.Username,
			Content:  content,
		})
		if err != nil {
			_ = c.AbortWithError(200, apiException.ServerError)
			return
		}
	}
	utils.JsonSuccessResponse(c, nil)
}
