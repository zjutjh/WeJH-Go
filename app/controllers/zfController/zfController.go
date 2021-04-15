package zfController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetClassTable(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		return
	}

	balance, err := funnelServices.GetClassTable(user)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, balance)
}
