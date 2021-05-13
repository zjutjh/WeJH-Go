package libraryController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiExpection"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetCurrent(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	list, err := funnelServices.GetCurrentBorrow(user)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, list)
}

func GetHistory(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.NotLogin)
		return
	}
	list, err := funnelServices.GetHistoryBorrow(user)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, list)
}
