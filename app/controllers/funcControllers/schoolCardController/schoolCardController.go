package schoolCardController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	balance, err := funnelServices.GetCardBalance(user)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, "Card")
		}
		_ = c.AbortWithError(200, err)
		return
	}
	if balance == nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"balance": balance})
}

func GetToday(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	balance, err := funnelServices.GetCardToday(user)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, "Card")
		}
		_ = c.AbortWithError(200, err)
		return
	}
	if balance == nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

type historyForm struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

func GetHistory(c *gin.Context) {
	var postForm historyForm
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
	list, err := funnelServices.GetCardHistory(user, postForm.Year, postForm.Month)
	if err != nil {
		if err == apiException.NoThatPasswordOrWrong {
			userServices.DelPassword(user, "Card")
		}
		_ = c.AbortWithError(200, err)
		return
	}
	if list == nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, list)
}
