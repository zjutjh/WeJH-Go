package schoolCardController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetBalance(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		return
	}

	balance, err := funnelServices.GetCardBalance(user)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"balance": balance})
}

func GetToday(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		return
	}

	balance, err := funnelServices.GetCardToday(user)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}

func GetHistory(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		return
	}

	list, err := funnelServices.GetCardHistory(user, 2021, 3)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, list)
}
